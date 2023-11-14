package models

import (
	"api-estadia-express/init/logger"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) CreateUser(db *sql.DB) error {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id")

	if err := db.QueryRow(queryBuilder.String(), &u.Name, &u.Email, &u.Password).Scan(&u.Id); err != nil {
		logger.Error("error to create user", err)
		return errors.New("error to create user")
	}

	return nil
}

func (u *User) FindUsersWithFilter(db *sql.DB) ([]User, error) {
	var listUsers []User
	var queryBuilder strings.Builder
	var conditions []string

	sql := `SELECT * FROM users`
	queryBuilder.WriteString(sql)

	if u.Id > 0 {
		conditions = append(conditions, "id = "+strconv.Itoa(u.Id))
	}

	if u.Email != "" {
		conditions = append(conditions, "email LIKE '%"+u.Email+"%'")
	}

	if u.Name != "" {
		conditions = append(conditions, "name LIKE '%"+u.Name+"%'")
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}

	rows, err := db.Query(queryBuilder.String())

	if err != nil {
		logger.Error("user not find", err)
		return listUsers, errors.New("user not find")
	}

	if rows != nil {
		for rows.Next() {
			single := User{}
			if err := rows.Scan(&single.Id, &single.Name, &single.Password, &single.Email); err != nil {
				logger.Error("error scan user", err)
				return listUsers, errors.New("error scan user")
			}
			listUsers = append(listUsers, single)
		}
	}

	return listUsers, nil
}

func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
