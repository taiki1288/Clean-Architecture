package domain
// この層はどこからでも呼び出せる。

// User型を宣言している。
type User struct {
	ID        int
	FirstName string
	LastName  string
	FullName  string
}

// User型のスライスを宣言している
type Users []User