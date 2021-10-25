package vehicle

import (
    "math/rand"
    "reflect"
    "time"

    "../../model/vehicle"
    "../../utils"
    "github.com/pkg/errors"
)

type VehicleRepository struct {
    vehicles map[vehicle.Type][]vehicle.Vehicle
    vehicleIDMapping map[string]vehicle.Vehicle
}

func (repository *VehicleRepository) AddMockVehicle(facility string , count int) {
    vehicleType := []vehicle.Type{vehicle.TYPE_SUV, vehicle.TYPE_SEDAN, vehicle.TYPE_HATCHBACK, vehicle.TYPE_BIKE}

    for i := 0; i < count; i++ {
        vType := vehicleType[rand.Intn(len(vehicleType))]
        veh := vehicle.Vehicle{
            PlateNumber: utils.RandStringRunes(10),
            Color:       "green",
            Make:        utils.RandStringRunes(4),
            Model:       utils.RandStringRunes(5),
            Year:        "1994",
            Cost:        float64(utils.RandIntRunes(1000)),
            Status:      vehicle.STATUS_AVAILABLE,
            Type:        vType,
            FacilityID:  facility,
        }
        repository.vehicles[vType] = append(repository.vehicles[vType], veh)
        repository.vehicleIDMapping[veh.PlateNumber] = veh
    }
}

type QueryOption func(vehicles []vehicle.Vehicle) ([]vehicle.Vehicle, error)

func Where(key string, value interface{}) QueryOption {
    return func(vehicles []vehicle.Vehicle) (res []vehicle.Vehicle, err error) {
        for _, v := range vehicles {
            r := reflect.ValueOf(v)
            f := reflect.Indirect(r).FieldByName(key)
            if f.Kind() == reflect.String && f.String() == value {
                res = append(res, v)
            }
            if f.Kind() == reflect.Int && f.Int() == value {
                res = append(res, v)
            }
        }

        return
    }
}

func (repository VehicleRepository) First(vehType vehicle.Type, opts ...QueryOption) (*vehicle.Vehicle, error) {

    if _, exists := repository.vehicles[vehType]; !exists {
        return nil, errors.New("given invalid vehicle type is not available")
    }
    vehicles := repository.vehicles[vehType]
    var res []vehicle.Vehicle
    var err error
    for _, opt := range opts {
        res, err = opt(vehicles)
        if err != nil {
            return nil, err
        }
    }

    if len(res) < 1 {
        return nil, errors.New("no vehicle found with applied filter")
    }
    return &res[0], nil
}

func (repository VehicleRepository) getLatestVehicle(vehType vehicle.Type) (*vehicle.Vehicle, error) {
    return repository.First(
        vehType,
        Where("year", time.Now().Year()),
        Where("model", "audi"),
    )
}