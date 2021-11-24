package usecase

import "Clean-Architecture/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err := interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}