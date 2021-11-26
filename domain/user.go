package domain

import "fmt"

// この層はどこからでも呼び出せる。

// User型を宣言している。
type User struct {
	ID        int
	FirstName string
	LastName  string
	FullName  string
}

// ユーザー情報を再構築してそれを返すメソッド
func (u *User) Build() *User {
	u.FullName = fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	return u
}

// User型のスライスを宣言している
type Users []User