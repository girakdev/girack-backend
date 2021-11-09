package usecase

import (
	"girack/app/domain"
)

type UserRepository interface {
  Store(domain.User) (int, error)
  FindById(int) (domain.User, error)
  Users() (domain.Users, error)
}

type UserInteractor struct {
  UseRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
  _, err = interactor.UseRepository.Store(u)
  return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
  user, err = interactor.UseRepository.FindById(identifier)
  return
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
  user, err = interactor.UserRepository.FindAll()
  return
}
