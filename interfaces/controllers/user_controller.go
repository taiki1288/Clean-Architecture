package controllers

import "Clean-Architecture/usecase"

type UserController struct {
	Interactor usecase.UserInteractor
}