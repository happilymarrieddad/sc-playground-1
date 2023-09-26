package types

import "time"

type Customer struct {
	ID        int64      `json:"id" xorm:"'id' pk autoincr"`
	Name      string     `validate:"required" json:"name" xorm:"name"`
	CreatedAt time.Time  `json:"created_at" xorm:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" xorm:"updated_at"`
}

func (*Customer) TableName() string {
	return "customers"
}
