package data

import (
	"time"

	_ "github.com/lib/pq"
)

// ------------------------------   USERS   ---------------------------------------------

// Create a new user, save user info into the database
func (user *User) Create() (err error) {
	// Postgres does not automatically return the last insert id, because it would be wrong to assume
	// you're always using a sequence. You need to use the RETURNING keyword in your insert to get this
	// information from postgres.
	statement := "INSERT INTO users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the User struct
	err = stmt.QueryRow(createUuid(), user.Name, user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	_, err = Db.Exec("UPDATE users SET name = $2, email = $3 where id = $1", user.Id, user.Name, user.Email)
	return
}

// Delete user from database
func (user *User) Delete() (err error) {
	_, err = Db.Exec("DELETE * FROM users WHERE id = $1", user.Id)
	return
}

// Get a single user given the email
func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	row := Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email)
	err = row.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func GetUserBySessionUuid(session_uuid string) (valid bool, user User) {
	err := Db.QueryRow("SELECT user_id  FROM sessions WHERE uuid = $1", session_uuid).
		Scan(&user.Id)
	if err != nil {
		valid = false
		return
	}
	if user.Id != 0 {
		valid = true
		Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", user.Id).
			Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	}
	return
}

// Get a single user given the uuid
func GetUserByUuid(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}

func GetAllUsers() (users []User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at"

	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(createUuid(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// Delete all users from database
func UserDeleteAll() (err error) {
	statement := "DELETE FROM users"
	_, err = Db.Exec(statement)
	return
}

// Get the session for an existing user
func (user *User) GetSession() (session Session, err error) {
	row := Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.Id)
	err = row.Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// ------------------------------   SESSIONS   ---------------------------------------------

// Check if session is valid in the database
func CheckSessionValidity(uuid string) (valid bool, err error) {
	var sess Session
	err = Db.QueryRow("SELECT id FROM sessions WHERE uuid = $1", uuid).
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
	err = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", session.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

// Delete session from database
func DeleteSessionByUuid(uuid string) (err error) {
	_, err = Db.Exec("DELETE FROM sessions WHERE uuid = $1", uuid)
	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	_, err = Db.Exec("DELETE FROM sessions")
	return
}
