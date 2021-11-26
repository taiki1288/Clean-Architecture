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
	// LastInsertIdメソッドで最後に保存されたidを取得している？
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func (repo *UserRepository) FindById(identifier int) (user domain.User, err error) {
	// クエリを実行して行を返すようにしている。
	row, err := repo.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	// deferを使って最後に実行されるようにしている。
	defer row.Close()
	if err != nil {
		return
	}
	var id int
	var firstName string
	var lastName string
	// Scanメソッドで読み取れるようにしている
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		// Scanはポインタ変数を渡してインターフェースを実装している。
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	user.Build()
	return
}

func (repo *UserRepository) FindAll() (users domain.Users, err error) {
	rows, err := repo.Query("SELECT id, first_name, last_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, *user.Build())
	}
	return
}