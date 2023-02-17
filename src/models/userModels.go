package models

type UserModel struct {
	Name   string `json:"name"  form:"name"`      // 名称
	Age    int    `json:"age" form:"age"`         // 年龄
	MyType string `json:"my_type" form:"my_type"` // 类型
}
