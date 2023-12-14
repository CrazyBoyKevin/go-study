package models

type User struct {
	ID         int64
	OpenId     string
	Nickname   string
	AvatarUrl  string
	Phone      string
	Gender     int32
	CreateTime string
	UpdateTime string
	Deleted    int32
}
