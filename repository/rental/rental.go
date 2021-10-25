package facility

import (
    "../../model/facility"
    "../../model/user"
    "../../utils"
)

type FacilityRepository struct {
    facilities map[string]facility.Facility
}

func (repository *FacilityRepository) RentVehicle()  {
    address := user.Address{
        Street:     "baker street",
        PinCode:    "110025",
        City:       "DELHI",
        Country:    "INDIA",
        Coordinate: user.Coordinate{ X: 100, Y: 200 },
    }
    fac := facility.NewFacility(utils.RandStringRunes(10), utils.RandStringRunes(10), address)
    repository.facilities[fac.FacilityNumber] = *fac
}
