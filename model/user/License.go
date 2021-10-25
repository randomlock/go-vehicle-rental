package user

import "time"

type LicenseType string

const (
    BLOCKED LicenseType = "BLOCKED"
    ACTIVE LicenseType = "ACTIVE"
)

type License struct {
    Id int
    ExpiryTime time.Time
    Status LicenseType
}
