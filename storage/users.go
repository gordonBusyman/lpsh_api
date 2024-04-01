package storage

import (
	"database/sql"
	"errors"
)

// Users is the storage struct.
type Users struct {
	db *sql.DB
}

// NewUsers returns a new Users struct.
func NewUsers(db *sql.DB) Users {
	return Users{
		db: db,
	}
}

// RetrieveSettings retrieves the settings of a User.
func (u Users) RetrieveSettings(code int) (string, error) {
	var settings string
	err := u.db.QueryRow("SELECT settings FROM users WHERE code = ?", code).
		Scan(&settings)
	if errors.Is(err, sql.ErrNoRows) {
		return "", errors.New("settings not found")
	}
	if err != nil {
		return "", errors.New("error retrieving settings")
	}

	return settings, nil
}

//********************************************************************************************************************
//
// The code below is commented out because it is not used in the application.
// Uncomment it if you want to use it.
//
//********************************************************************************************************************

//// Resource is the resource struct that describes User.
//type Resource struct {
//	ID       int    `json:"id"`
//	TgName   string `json:"tg_name"`
//	FullName string `json:"full_name"`
//
//	Code     int    `json:"code"`
//	Settings string `json:"name"`
//
//	CreatedAt string `json:"created_at"`
//}

//
//// Retrieve retrieves a User.
//func (u Users) Retrieve(id int) (Resource, error) {
//	var r Resource
//	err := u.db.QueryRow("SELECT * FROM users WHERE id = ?", id).
//		Scan(&r.ID, &r.TgName, &r.FullName, &r.Code, &r.Settings, &r.CreatedAt)
//	if err != nil {
//		return Resource{}, err
//	}
//
//	return r, nil
//}
//
//// Create creates a User.
//func (u Users) Create(r Resource) error {
//	_, err := u.db.Exec("INSERT INTO users (tg_name, full_name, code, settings) VALUES (?, ?, ?, ?)",
//		r.TgName, r.FullName, r.Code, r.Settings)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// UpdateCode updates the code of a User.
//func UpdateCode(u Users, id, code int) error {
//	_, err := u.db.Exec("UPDATE users SET code = ? WHERE id = ?", code, id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// Delete deletes a User.
//func (u Users) Delete(id int) error {
//	_, err := u.db.Exec("DELETE FROM users WHERE id = ?", id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
