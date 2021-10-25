package user

import (
    "time"

    "../../model/user"
    "../../utils"
    "github.com/pkg/errors"
)

type UserRepository struct {
    customers map[string]*user.User
    admins map[string]*user.User
}

func (repository *UserRepository) RegisterCustomers(name string) (err error)  {
    userinfo := user.UserInfo{
        Name:   name,
        Gender: user.MALE,
    }
    license := user.License{
        Id:         utils.RandIntRunes(10),
        ExpiryTime: time.Time{},
        Status:     user.ACTIVE,
    }
    
    
    u := user.NewUser(license, userinfo, user.CUSTOMER)
    if u.License.ExpiryTime.Sub(time.Now()) > 0 {
        return errors.New("License is expired. cannot register the customer")
    }
    repository.customers[u.Id] = u
    return
}

func (repository *UserRepository) RegisterAdmin(name string) (err error)  {
    userinfo := user.UserInfo{
        Name:   name,
        Gender: user.MALE,
    }
    license := user.License{
        Id:         utils.RandIntRunes(10),
        ExpiryTime: time.Time{},
        Status:     user.ACTIVE,
    }


    u := user.NewUser(license, userinfo, user.ADMIN)
    if u.License.ExpiryTime.Sub(time.Now()) > 0 {
        return errors.New("License is expired. cannot register the customer")
    }
    repository.admins[u.Id] = u
    return
}