package user


type UserType string

const (
    ADMIN UserType = "ADMIN"
    CUSTOMER UserType = "CUSTOMER"
)



type User struct {
    Id  string
    License License
    Info UserInfo
    UserType UserType
}

func NewUser(license License, info UserInfo, userType UserType) *User {
    return &User{License: license, Info: info, UserType: userType}
}




