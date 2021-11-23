package database

// これは内側に依存しているので依存関係に関してはOK
import "Clean-Architecture/domain"

type UserRepository struct {
	SqlHandler
}

func (repo *UserRepository) Store(u domain.User) (id int, err error) {
	// 実行したクエリの結果を渡している。
	result, err := repo.Execute(
		"INSERT INTO users (first_name, last_name) VALUES (?, ?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return
	}
	
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}