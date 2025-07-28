package service

import (
	"app/internal"
	"errors"
)

type ServiceVehicleDefaultStub struct {
}

func NewServiceVehicleDefaultStub() *ServiceVehicleDefaultStub {
	return &ServiceVehicleDefaultStub{}
}

func (s *ServiceVehicleDefaultStub) FindByColorAndYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {
	if color == "blue" {
		return v, errors.New("error")
	}
	v = make(map[int]internal.Vehicle)
	v[0] = internal.Vehicle{VehicleAttributes: internal.VehicleAttributes{Color: color, FabricationYear: fabricationYear}}

	return v, nil
}

func (s *ServiceVehicleDefaultStub) FindByBrandAndYearRange(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	if brand == "err" {
		return v, errors.New("error")
	}
	return
}

func (s *ServiceVehicleDefaultStub) AverageMaxSpeedByBrand(brand string) (a float64, err error) {
	switch brand {
	case "err":
		return a, errors.New("error")
	case "not found":
		return a, internal.ErrServiceNoVehicles
	}
	return
}

func (s *ServiceVehicleDefaultStub) AverageCapacityByBrand(brand string) (a int, err error) {
	switch brand {
	case "err":
		return a, errors.New("error")
	case "not found":
		return a, internal.ErrServiceNoVehicles
	}
	return
}

func (s *ServiceVehicleDefaultStub) SearchByWeightRange(query internal.SearchQuery, ok bool) (v map[int]internal.Vehicle, err error) {
	if query.FromWeight == 0.0 {
		return v, errors.New("error")
	}
	return
}
