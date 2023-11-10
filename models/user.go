package models

import (
	"api-estadia-express/init/db"
	"strconv"
	"strings"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"password"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) FindUserWithFilter() []User {
	instancePG := db.ConnectionToPostgres()
	var queryBuilder strings.Builder
	var conditions []string

	sql := `SELECT * FROM users`
	queryBuilder.WriteString(sql)

	if u.Id > 0 {
		conditions = append(conditions, "id = "+strconv.Itoa(u.Id))
	}

	if u.Email != "" {
		conditions = append(conditions, "email LIKE UPPER('%"+u.Email+"%')")
	}

	if u.Name != "" {
		conditions = append(conditions, "name LIKE UPPER('%"+u.Name+"%')")
	}

	if len(conditions) > 0 {
		queryBuilder.WriteString(" WHERE ")
		queryBuilder.WriteString(strings.Join(conditions, " AND "))
	}

	rows, err := instancePG.Query(queryBuilder.String())

	if err != nil {

	}

}
