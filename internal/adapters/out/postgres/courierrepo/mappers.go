package courierrepo

import (
	"delivery/internal/core/domain/kernel"
	"delivery/internal/core/domain/model/courier"
)

func DomainToDTO(aggregate *courier.Courier) CourierDTO {
	var courierDTO CourierDTO

	courierDTO.ID = aggregate.ID()
	courierDTO.Name = aggregate.Name()
	courierDTO.Speed = aggregate.Speed()
	courierDTO.Location = LocationDTO{
		X: aggregate.Location().X(),
		Y: aggregate.Location().Y(),
	}
	courierDTO.StoragePlaces = make([]*StoragePlaceDTO, 0)
	for _, stPlace := range aggregate.StoragePlaces() {
		storagePlaceDTO := &StoragePlaceDTO{
			ID:          stPlace.ID(),
			OrderID:     stPlace.OrderID(),
			Name:        stPlace.Name(),
			TotalVolume: stPlace.TotalVolume(),
			CourierID:   aggregate.ID(),
		}
		courierDTO.StoragePlaces = append(courierDTO.StoragePlaces, storagePlaceDTO)
	}

	return courierDTO
}

func DtoToDomain(dto CourierDTO) *courier.Courier {
	var aggregate *courier.Courier
	var storagePlaces []*courier.StoragePlace

	for _, stPlace := range dto.StoragePlaces {
		place := courier.RestoreStoragePlace(stPlace.ID, stPlace.Name,
			stPlace.TotalVolume, stPlace.OrderID)
		storagePlaces = append(storagePlaces, place)

	}
	location, _ := kernel.NewLocation(dto.Location.X, dto.Location.Y)
	aggregate = courier.RestoreCourier(dto.ID, dto.Name, dto.Speed, location, storagePlaces)

	return aggregate
}
