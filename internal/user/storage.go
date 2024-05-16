package user

import (
	"github.com/D1Y0RBEKORIFJONOV/rest-api-project/internal/postgres"
)

func CreateUser(users *Users, user *User) (err error) {
	db := postgres.DB{}

	user.ID, err = db.UserInsertInto(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	users.Users = append(users.Users, *user)
	return nil
}

func ReadUser() (users Users, err error) {
	db := postgres.DB{}
	err = db.ConnectDB()
	if err != nil {
		return users, err
	}
	defer db.DB.Close()
	query := `
	SELECT * FROM users ;
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		users.Users = append(users.Users, user)
	}
	return users, nil
}
