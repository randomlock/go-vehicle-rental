package user

type Gender string

const (
    MALE Gender = "MALE"
    FEMALE Gender = "FEMALE"
)

type UserInfo struct {
    Name string
    address Address
    Gender Gender
}