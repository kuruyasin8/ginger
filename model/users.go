package model

type User struct {
	ID        uint   `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Age       uint   `json:"age" bson:"age"`
	Birthday  int64  `json:"birthday" bson:"birthday"`
}
