package usecase

import "Clean-Architecture/domain"

type UserInteractor struct {
	UserRepository UserRepository
}