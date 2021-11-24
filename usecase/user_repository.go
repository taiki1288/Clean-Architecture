package usecase

import "Clean-Architecture/domain"

// database層は外側だから、ここでもDIPを利用して依存ルールを守る。
type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}