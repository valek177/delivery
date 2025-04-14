# Демо проект к курсу "Domain Driven Design и Clean Architecture на языке Go"
📚 Подробнее о курсе: [microarch.ru/courses/ddd/languages/go](https://microarch.ru/courses/ddd/languages/go?utm_source=gitlab&utm_medium=repository&utm_content=basket)

---

# OpenApi (генерация HTTP сервера)
```
oapi-codegen -config configs/server.cfg.yaml https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/delivery/contracts/openapi.yml
```

# БД
```
https://pressly.github.io/goose/installation/
```

# Запросы к БД
```
-- Выборки
SELECT * FROM public.couriers;
SELECT * FROM public.transports;
SELECT * FROM public.orders;

SELECT * FROM public.outbox;

-- Очистка БД (все кроме справочников)
DELETE FROM public.couriers;
DELETE FROM public.transports;
DELETE FROM public.orders;
DELETE FROM public.outbox;

-- Добавить курьеров
    
-- Пеший
INSERT INTO public.couriers(
    id, name, location_x, location_y, status)
VALUES ('bf79a004-56d7-4e5f-a21c-0a9e5e08d10d', 'Пеший', 1, 3, 'Free');
INSERT INTO public.transports(
    id, name, speed, courier_id)
VALUES ('921e3d64-7c68-45ed-88fb-97ceb8148a7e', 'Пешком', 1, 'bf79a004-56d7-4e5f-a21c-0a9e5e08d10d');


-- Вело
INSERT INTO public.couriers(
    id, name, location_x, location_y, status)
VALUES ('db18375d-59a7-49d1-bd96-a1738adcee93', 'Вело', 4,5, 'Free');
INSERT INTO public.transports(
    id, name, speed, courier_id)
VALUES ('b96a9d83-aefa-4d06-99fb-e630d17c3868', 'Велосипед', 2, 'db18375d-59a7-49d1-bd96-a1738adcee93');

-- Авто
INSERT INTO public.couriers(
    id, name, location_x, location_y, status)
VALUES ('407f68be-5adf-4e72-81bc-b1d8e9574cf8', 'Авто', 7,9, 'Free');
INSERT INTO public.transports(
    id, name, speed,courier_id)
VALUES ('c24d
```

# gRPC Client
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
export PATH="$PATH:$(go env GOPATH)/bin"

curl -o ./api/proto/geo_service.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/geo/contracts/contract.proto
protoc --go_out=./pkg/clients/geo --go-grpc_out=./pkg/clients/geo ./api/proto/geo_service.proto

```

# Kafka
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
export PATH="$PATH:$(go env GOPATH)/bin"

curl -o ./api/proto/basket_confirmed.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/basket/contracts/basket_confirmed.proto
protoc --go_out=./pkg ./api/proto/basket_confirmed.proto

curl -o ./api/proto/order_status_changed.proto https://gitlab.com/microarch-ru/ddd-in-practice/system-design/-/raw/main/services/delivery/contracts/order_status_changed.proto
protoc --go_out=./pkg ./api/proto/order_status_changed.proto
```

# Тестирование
```
mockery
```

# Документация используемых библилиотек
* [Goose] (https://github.com/pressly/goose)
* [Oapi-codegen] (https://github.com/oapi-codegen/oapi-codegen)
* [Protobuf] (https://protobuf.dev/reference/go/go-generated/)
* [gRPC] (https://grpc.io/docs/languages/go/)
* [Mockery] (https://vektra.github.io/mockery/latest/)

# Лицензия

Код распространяется под лицензией [MIT](./LICENSE).  
© 2025 microarch.ru