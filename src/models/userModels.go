package models

type UserModel struct {
	Name   string `form:"name" json:"name"`     // 名称
	Age    int    `form:"int" json:"age"`       // 年龄
	MyType string `form:"myType" json:"myType"` // 类型
}
