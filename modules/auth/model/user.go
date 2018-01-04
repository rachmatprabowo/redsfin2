package model

import (
	"github.com/rachmatprabowo/redsfin2/core"
)

// User struct of user
type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	core.BaseModel
}

// Users _
type Users []User

// Save _
func (usr *User) Save() bool {
	_, err := core.MasterDB.Exec("INSERT INTO auth_user (status, username, email, password, role_id, created_by) VALUES (1, $1, $2, $3, $4, $5)", usr.Username, usr.Email, usr.Password, usr.Role.ID, &usr.BaseModel.CreatedBy)
	return core.CheckErr(err, "Unknown errors was occured")
}

// Update _
func (usr *User) Update(oldPassword, newPassword string) bool {
	var password string
	row := core.MasterDB.QueryRow("SELECT password FROM auth_user WHERE id = $1", usr.ID)
	err := row.Scan(&password)
	core.CheckErr(err, "unknown error was occured")

	if password == newPassword {
		_, err := core.MasterDB.Exec("UPDATE auth_user SET username = $1, email = $2, password = $3, role_id = $4 WHERE id = $5", usr.Username, usr.Email, newPassword, usr.Role.ID, usr.ID)
		return core.CheckErr(err, "unknown error was occured")
	}
	return false
}

// Load _
func (usr *User) Load(id int) bool {
	row := core.MasterDB.QueryRow("SELECT id, username, email, role_id, created_date, created_by FROM auth_user WHERE id = $1", id)
	err := row.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.Role.ID, &usr.BaseModel.CreatedDate, &usr.BaseModel.CreatedBy)
	return core.CheckErr(err, " unknown error was occured")
}

// LoadAll _
func (usrs *Users) LoadAll() bool {
	rows, err := core.MasterDB.Query("SELECT id, username, email, role_id, created_date, created_by FROM auth_user")
	if core.CheckErr(err, "unknown error was occured") {
		for rows.Next() {
			var usr User
			rows.Scan(&usr.ID, &usr.Username, &usr.Email, &usr.Role.ID, &usr.BaseModel.CreatedDate, &usr.BaseModel.CreatedBy)
			*usrs = append(*usrs, usr)
		}
		rows.Close()
		return true
	}

	return false
}
