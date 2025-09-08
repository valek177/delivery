package cmd

import (
	"log"
	"reflect"
	"sync"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"

	kafkain "delivery/internal/adapters/in/kafka"
	grpcout "delivery/internal/adapters/out/grpc/geo"
	kafkaout "delivery/internal/adapters/out/kafka"
	"delivery/internal/adapters/out/postgres"
	"delivery/internal/adapters/out/postgres/outboxrepo"
	"delivery/internal/core/application/eventhandlers"
	"delivery/internal/core/application/usecases/commands"
	"delivery/internal/core/application/usecases/queries"
	"delivery/internal/core/domain/model/order"
	"delivery/internal/core/domain/services"
	"delivery/internal/core/ports"
	"delivery/internal/jobs"
	"delivery/internal/pkg/ddd"
	"delivery/internal/pkg/outbox"
)

type CompositionRoot struct {
	configs Config
	gormDb  *gorm.DB

	onceGeo   sync.Once
	geoClient ports.GeoClient

	closers []Closer
}

func NewCompositionRoot(configs Config, gormDb *gorm.DB) *CompositionRoot {
	return &CompositionRoot{
		configs: configs,
		gormDb:  gormDb,
	}
}

func (cr *CompositionRoot) NewOrderDispatcher() services.OrderDispatcher {
	orderDispatcher := services.NewOrderDispatcher()
	return orderDispatcher
}

func (cr *CompositionRoot) NewUnitOfWork() ports.UnitOfWork {
	unitOfWork, err := postgres.NewUnitOfWork(cr.gormDb, cr.NewMediatrWithSubscriptions())
	if err != nil {
		log.Fatalf("cannot create UnitOfWork: %v", err)
	}
	return unitOfWork
}

func (cr *CompositionRoot) NewUnitOfWorkFactory() ports.UnitOfWorkFactory {
	unitOfWorkFactory, err := postgres.NewUnitOfWorkFactory(cr.gormDb,
		cr.NewMediatrWithSubscriptions())
	if err != nil {
		log.Fatalf("cannot create UnitOfWorkFactory: %v", err)
	}
	return unitOfWorkFactory
}

func (cr *CompositionRoot) NewCreateOrderCommandHandler() commands.CreateOrderCommandHandler {
	commandHandler, err := commands.NewCreateOrderCommandHandler(cr.NewUnitOfWorkFactory(),
		cr.NewGeoClient())
	if err != nil {
		log.Fatalf("cannot create CreateOrderCommandHandler: %v", err)
	}
	return commandHandler
}

func (cr *CompositionRoot) NewCreateCourierCommandHandler() commands.CreateCourierCommandHandler {
	commandHandler, err := commands.NewCreateCourierCommandHandler(cr.NewUnitOfWorkFactory())
	if err != nil {
		log.Fatalf("cannot create CreateCourierCommandHandler: %v", err)
	}
	return commandHandler
}

func (cr *CompositionRoot) NewAssignOrderCommandHandler() commands.AssignOrderCommandHandler {
	commandHandler, err := commands.NewAssignOrderCommandHandler(
		cr.NewUnitOfWorkFactory(), cr.NewOrderDispatcher())
	if err != nil {
		log.Fatalf("cannot create AssignOrderCommandHandler: %v", err)
	}
	return commandHandler
}

func (cr *CompositionRoot) NewMoveCouriersCommandHandler() commands.MoveCouriersCommandHandler {
	commandHandler, err := commands.NewMoveCouriersCommandHandler(
		cr.NewUnitOfWorkFactory())
	if err != nil {
		log.Fatalf("cannot create MoveCouriersCommandHandler: %v", err)
	}
	return commandHandler
}

func (cr *CompositionRoot) NewGetCouriersQueryHandler() queries.GetCouriersQueryHandler {
	queryHandler, err := queries.NewGetCouriersQueryHandler(cr.gormDb)
	if err != nil {
		log.Fatalf("cannot create GetCouriersQueryHandler: %v", err)
	}
	return queryHandler
}

func (cr *CompositionRoot) NewGetNotCompletedOrdersQueryHandler() queries.GetNotCompletedOrdersQueryHandler {
	queryHandler, err := queries.NewGetNotCompletedOrdersQueryHandler(cr.gormDb)
	if err != nil {
		log.Fatalf("cannot create GetNotCompletedOrdersQueryHandler: %v", err)
	}
	return queryHandler
}

func (cr *CompositionRoot) NewAssignOrdersJob() cron.Job {
	job, err := jobs.NewAssignOrdersJob(cr.NewAssignOrderCommandHandler())
	if err != nil {
		log.Fatalf("cannot create AssignOrdersJob: %v", err)
	}
	return job
}

func (cr *CompositionRoot) NewMoveCouriersJob() cron.Job {
	job, err := jobs.NewMoveCouriersJob(cr.NewMoveCouriersCommandHandler())
	if err != nil {
		log.Fatalf("cannot create MoveCouriersJob: %v", err)
	}
	return job
}

func (cr *CompositionRoot) NewGeoClient() ports.GeoClient {
	cr.onceGeo.Do(func() {
		client, err := grpcout.NewClient(cr.configs.GeoServiceGrpcHost)
		if err != nil {
			log.Fatalf("cannot create GeoClient: %v", err)
		}
		cr.RegisterCloser(client)
		cr.geoClient = client
	})
	return cr.geoClient
}

func (cr *CompositionRoot) NewBasketConfirmedConsumer() kafkain.BasketConfirmedConsumer {
	consumer, err := kafkain.NewBasketConfirmedConsumer(
		[]string{cr.configs.KafkaHost},
		cr.configs.KafkaConsumerGroup,
		cr.configs.KafkaBasketConfirmedTopic,
		cr.NewCreateOrderCommandHandler(),
	)
	if err != nil {
		log.Fatalf("cannot create BasketConfirmedConsumer: %v", err)
	}
	cr.RegisterCloser(consumer)
	return consumer
}

func (cr *CompositionRoot) NewOrderProducer() ports.OrderProducer {
	producer, err := kafkaout.NewOrderProducer([]string{cr.configs.KafkaHost},
		cr.configs.KafkaOrderChangedTopic)
	if err != nil {
		log.Fatalf("ERROR: cannot create OrderProducer: %v", err)
	}

	cr.RegisterCloser(producer)
	return producer
}

func (cr *CompositionRoot) NewOrderCompletedDomainEventHandler() ddd.EventHandler {
	producer := cr.NewOrderProducer()
	handler, err := eventhandlers.NewOrderCompletedDomainEventHandler(producer)
	if err != nil {
		log.Fatalf("ERROR: cannot create OrderCompletedDomainEventHandler: %v", err)
	}

	return handler
}

func (cr *CompositionRoot) NewMediatrWithSubscriptions() ddd.Mediatr {
	mediatr := ddd.NewMediatr()
	mediatr.Subscribe(cr.NewOrderCompletedDomainEventHandler(), order.NewEmptyCompletedDomainEvent())
	return mediatr
}

func (cr *CompositionRoot) NewOutboxRepository() outboxrepo.OutboxRepository {
	repository, err := outboxrepo.NewRepository(cr.gormDb)
	if err != nil {
		log.Fatalf("ERROR: cannot create OutboxRepository: %v", err)
	}

	return repository
}

func (cr *CompositionRoot) NewEventRegistry() outbox.EventRegistry {
	registry, err := outbox.NewEventRegistry()
	if err != nil {
		log.Fatalf("ERROR: cannot create EventRegistry: %v", err)
	}

	err = registry.RegisterDomainEvent(reflect.TypeOf(order.CompletedDomainEvent{}))
	if err != nil {
		log.Fatalf("ERROR: cannot register domain event: %v", err)
	}

	return registry
}

func (cr *CompositionRoot) NewOutboxJob() cron.Job {
	job, err := jobs.NewOutboxJob(cr.NewOutboxRepository(), cr.NewEventRegistry(),
		cr.NewMediatrWithSubscriptions())
	if err != nil {
		log.Fatalf("ERROR: cannot create OutboxJob: %v", err)
	}

	return job
}
