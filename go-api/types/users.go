package types

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userType string

const (
	AdminUserType    userType = "Admin"
	StandardUserType userType = "Standard User"
)

func NewUser(firstName, lastName, email, password string, customerID int64) *User {
	u := &User{
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		CustomerID: customerID,
	}

	u.SetPassword(password)

	return u
}

type User struct {
	ID         int64      `json:"id" xorm:"'id' pk autoincr"`
	FirstName  string     `json:"first_name" xorm:"first_name"`
	LastName   string     `json:"last_name" xorm:"last_name"`
	Email      string     `json:"email" xorm:"email"`
	Password   string     `json:"-" xorm:"password"`
	Addr1      string     `json:"addr1" xorm:"addr1"`
	Addr2      string     `json:"addr2" xorm:"addr2"`
	City       string     `json:"city" xorm:"city"`
	State      string     `json:"state" xorm:"state"`
	County     string     `json:"county" xorm:"county"`
	PostalCode string     `json:"postal_code" xorm:"postal_code"`
	BirthMonth int64      `json:"birth_month" xorm:"birth_month"`
	BirthDay   int64      `json:"birth_day" xorm:"birth_day"`
	BirthYear  int64      `json:"birth_year" xorm:"birth_year"`
	UserType   userType   `json:"type" xorm:"user_type"`
	CustomerID int64      `validate:"required" json:"customer_id" xorm:"customer_id"`
	CreatedAt  time.Time  `json:"createdAt" xorm:"created_at"`
	UpdatedAt  *time.Time `json:"updatedAt" xorm:"updated_at"`
}

func (*User) TableName() string {
	return `users`
}

func (u *User) SetPassword(psw string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) PasswordMatches(psw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(psw)) == nil
}
