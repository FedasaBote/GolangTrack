package models

type User struct {
	ID 			string 		`bson:"_id,omitempty" json:"id,omitempty"`
	Email 		string 		`bson:"email" json:"email"`
	Password 	string 		`bson:"password" json:"password"`
	Role 		string 		`bson:"role,omitempty" json:"role"`
}