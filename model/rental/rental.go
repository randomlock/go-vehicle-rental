package rental

import "time"

type Status string

const (
    AVAILABLE Status = "BLOCKED"
    BOOKED Status = "ACTIVE"
    INUSE  Status = "IN-USE"
    LOST Status = "LOST"
)

type Rental struct {
    ID string
    CustomerID string
    VehicleID string
    ReceiptID string
    Status Status
    StartDate time.Time
    EndDate time.Time
}
