package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// User
type User struct {

	// Email
	Email string `json:"email" xml:"email"`

	// Password
	Password string `json:"password" xml:"password"`

	// Phone
	Phone string `json:"phone" xml:"phone"`

	// UserStatus User Status
	UserStatus int32 `json:"userStatus" xml:"userStatus"`

	// ID
	ID int64 `json:"id" xml:"id"`

	// Username
	Username string `json:"username" xml:"username"`

	// FirstName
	FirstName string `json:"firstName" xml:"firstName"`

	// LastName
	LastName string `json:"lastName" xml:"lastName"`
}
