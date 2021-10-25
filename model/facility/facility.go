package facility

import (
    "../user"
)


type Facility struct {
    FacilityName string
    FacilityNumber string
    Address user.Address
}

func NewFacility(facilityName string, facilityNumber string, address user.Address) *Facility {
    return &Facility{FacilityName: facilityName, FacilityNumber: facilityNumber, Address: address}
}

// func (f *Facility) AddVehicle(vehicle vehicle.Vehicle)  {
//     f.Vehicles[vehicle.Type] = append(f.Vehicles[vehicle.Type], vehicle)
// }
//
// func (f *Facility) RemoveVehicle(vehicle vehicle.Vehicle)  {
//     vehicles := f.Vehicles[vehicle.Type]
//
//     for i , veh := range vehicles {
//         if veh.PlateNumber == vehicle.PlateNumber {
//             f.Vehicles[vehicle.Type] = append(f.Vehicles[vehicle.Type][:i], f.Vehicles[vehicle.Type][i+1:]...)
//             break
//         }
//     }
// }