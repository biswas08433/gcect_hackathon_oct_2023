package data

import (
	"time"
)

// ------------------------------   USERS   ---------------------------------------------

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence. You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := `
			INSERT INTO "users" ("uuid", "first_name", "last_name", "email", "password", "created_at") 
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING "id", "uuid", "created_at";
	`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(
		createUuid(),
		user.FirstName,
		user.LastName,
		user.Email,
		Encrypt(user.Password),
		time.Now(),
	).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	_, err = Db.Exec(`
	UPDATE "users" SET "first_name" = $2, "first_name" = $3, "email" = $4
	WHERE "id" = $1`, user.Id, user.FirstName, user.LastName, user.Email)
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	_, err = Db.Exec(`DELETE * FROM "users" WHERE "id" = $1`, user.Id)
	return
}

// Get a single user given the email
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	row := Db.QueryRow(`SELECT "id", "uuid", "first_name", "last_name", "email", "password", "created_at"
	FROM "users"
	WHERE "email" = $1`, email)
	err = row.Scan(&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	return
}

// Get a single user given the uuid
func GetUserByUuid(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow(`
		SELECT "id", "uuid", "first_name", "last_name", "email", "password", "created_at"
		FROM "users" 
		WHERE "uuid" = $1
		`, uuid).Scan(&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func GetAllUsers() (users []User, err error) {
	rows, err := Db.Query(`
	SELECT "id", "uuid", "first_name", "last_name", "email", "password", "created_at" 
	FROM "users"
	`)
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := `DELETE FROM "users"`
	_, err = Db.Exec(statement)
	return
}

// ------------------------------   SESSIONS   ---------------------------------------------

// Check if session is valid in the database
func CheckSessionValidity(uuid string) (valid bool, err error) {
	var sess Session
	err = Db.QueryRow(`SELECT "id" FROM "sessions" WHERE "uuid" = $1`, uuid).
		Scan(&sess.Id)
	if err != nil {
		valid = false
		return
	}
	if sess.Id != 0 {
		valid = true
	}
	return
}

// Get the user from the session
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow(`
	SELECT "id", "uuid", "first_name", "last_name", "email", "created_at"
	FROM users 
	WHERE id = $1
	`, session.UserId).Scan(&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	return
}

func (user *User) AddSubject(subject string) (err error) {
	var subject_id int
	statement := `
			INSERT INTO "subjects" ("title") 
			VALUES ($1)
			RETURNING "id";
	`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// use QueryRow to return a row and scan the returned id into the User struct
	_ = stmt.QueryRow(
		subject,
	).Scan(&subject_id)

	statement = `
		INSERT INTO "teacher_subject" ("teacher_id", "subject_id")
		VALUES ($1, $2);
	`
	stmt, err = Db.Prepare(statement)

	if err != nil {
		return
	}
	// use QueryRow to return a row and scan the returned id into the User struct
	_ = stmt.QueryRow(
		user.Id,
		subject_id,
	).Scan()
	return
}
