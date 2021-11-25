package controllers

// gin.Contextを使用するためそれに利用するためのメソッドのインターフェース
type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
}