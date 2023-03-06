package models

type User struct {
	Id           int    `json:"id" bson:"id" validate:"omitempty"`
	FirstName    string `json:"firstName" bson:"firstName" validate:"required,gte=2"`
	LastName     string `json:"lastName" bson:"lastName"`
	Email        string `json:"email" bson:"email" validate:"required_with=Id,email"`
	BusinessType string `json:"businessType" bson:"businessType"`
	PhoneNo      string `json:"phoneNo" bson:"phoneNo" validate:"number"`
	CompanyName  string `json:"companyName" bson:"companyName" `
	Country      string `json:"country" bson:"country" validate:"required"`
}
