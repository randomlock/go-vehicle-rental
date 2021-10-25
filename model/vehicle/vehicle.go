package vehicle


type Status string
type Type string

const (
    STATUS_AVAILABLE Status = "BLOCKED"
    STATUS_BOOKED Status = "ACTIVE"
    STATUS_INUSE  Status = "IN-USE"
    STATUS_LOST Status = "LOST"

    TYPE_SUV Type = "SUV"
    TYPE_BIKE Type = "BIKE"
    TYPE_SEDAN Type = "SEDAN"
    TYPE_HATCHBACK Type = "HATCHBACK"
)

type Vehicle struct {
    PlateNumber string
    Color string
    Make string
    Model string
    Year string
    Cost float64
    Status Status
    Type Type
    FacilityID string
}

