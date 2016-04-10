package user

import (
	"time"
)

// Entity information
type Entity struct {
	Id         string    `db:"id" json:"id"`
	First_name string    `db:"first_name" json:"first_name"` // First name
	Last_name  string    `db:"last_name" json:"last_name"`   // Last name
	Email      string    `db:"email" json:"email"`
	Password   string    `db:"password" json:"password"`
	Status_id  int       `db:"status_id" json:"status_id"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	Updated_at time.Time `db:"updated_at" json:"updated_at"`
	Deleted_at time.Time `db:"deleted_at" json:"deleted_at"`
}
