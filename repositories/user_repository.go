package repositories

import (
	"database/sql"

	"github.com/rymccue/golang-auth-microservice/app"
	"github.com/rymccue/golang-auth-microservice/utils/crypto"
)

// GetUserByEmail gets a user by email
func GetUserByEmail(db *sql.DB, email string) (*app.User, error) {
	const sqlstr = `
	select
		first_name,
		last_name,
		email,
		password,
		salt
	from users
	where email = $1
	`
	var user app.User
	err := db.QueryRow(sqlstr, email).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Salt)
	return &user, err
}

// AddUserToDatabase creates a new user
func AddUserToDatabase(db *sql.DB, firstName, lastName, email, password string) error {
	const sqlstr = `
	insert into users (
		first_name,
		last_name,
		email,
		password,
		salt
	) values (
		$1,
		$2,
		$3,
		$4,
		$5
	) returning id
	`
	salt := crypto.GenerateSalt()
	hashedPassword := crypto.HashPassword(password, salt)
	var id int
	err := db.QueryRow(sqlstr, firstName, lastName, email, hashedPassword, salt).Scan(&id)
	return err
}

func CheckEmailExists(db *sql.DB, email string) (bool, error) {
	const sqlstr = "select exists(select 1 from users where email = $1)"

	var exists bool
	err := db.QueryRow(sqlstr, email).Scan(&exists)
	return exists, err
}
