package model

type User struct {
	FirstName   string       `json:"first_name" bson:"first_name"`
	LastName    string       `json:"last_name" bson:"last_name"`
	Birthday    int64        `json:"birthday" bson:"birthday"`
	Email       string       `json:"email" bson:"email"`
	Password    string       `json:"password" bson:"-"`
	Credentials *Credentials `json:"credentials" bson:"credentials"`
}

type Credentials struct {
	Hash       string `json:"hash" bson:"hash"`
	Salt       string `json:"salt" bson:"salt"`
	Verified   bool   `json:"verified" bson:"verified"`
	CreatedAt  int64  `json:"created_at" bson:"created_at"`
	ModifiedAt int64  `json:"modified_at" bson:"modified_at"`
}
