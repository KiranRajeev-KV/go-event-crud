package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id`
	err := m.DB.QueryRowContext(ctx, stmt, user.Email, user.Password, user.Name).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}
