package usecase

import "Clean-Architecture/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

// User情報を取得して返却するメソッド
func (interactor *UserInteractor) Add(u domain.User) (user domain.User, err error) {
	// ユーザーを保存する処理
	identifier, err := interactor.UserRepository.Store(u)
	if err != nil {
		return
	}
	// 保存されたユーザー情報を返す処理
	user, err = interactor.UserRepository.FindById(identifier)
	return
}

// ユーザー一覧を返すメソッド
func (interactor *UserInteractor) Users() (user domain.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user domain.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}