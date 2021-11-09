package usecase

import "girack/domain"


type UserRepository interface {
    Store(domain.User) (int, error)
    FindById(int) (domain.User, error)
    FindAll() (domain.Users, error)
}
