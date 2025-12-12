package models

type Person struct {
	No       int    `json:"no"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
}

type PersonInfo struct {
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
}
