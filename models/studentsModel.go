package models

type Student struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
}
