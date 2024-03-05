package models

type PersonalInfo struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName string `bson:"last_name" json:"last_name"`
	DOB string `bson:"dob" json:"dob"`
	Gender string `bson:"gender" json:"gender"`
	PhoneNumber string `bson:"phone_number" json:"phone_number"`
	Email string `bson:"email" json:"email"`
	Profession string `bson:"profession" json:"profession"`
}


type User struct {
	UserID int `bson:"user_id" json:"user_id"`
	Username string `bson:"username" json:"username"` 
	Password string `bson:"password" json:"password"`
	PersonalInfo PersonalInfo `bson:"personal_info" json:"personal_info"`
}