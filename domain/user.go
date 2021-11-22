package domain
// この層はどこからでも呼び出せる。

type User struct {
	ID        int
	FirstName string
	LastName  string
}

type Users []User