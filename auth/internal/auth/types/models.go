package types

type (
	User struct {
		Username string `json:"name" bson:"name"`
		Password string `json:"password" bson:"password"`
	}
)
