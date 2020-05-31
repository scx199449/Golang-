package entity

type User struct {
	Id          int64  `json:"_id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	MobilePhone string `json:"mobilephone" bson:"mobilephone"`
}
