package model

type UserSearch struct {
	Id int `form:"id" binding:"min=1,required"`
}

type User struct {
	Age  int64  `json:"age"`  // Age
	Id   int64  `json:"id"`   // Id
	Name string `json:"name"` // Name
}

func (u *User) TableName() string {
	return "user"
}
