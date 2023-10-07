package data

import "time"

// Create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := `
	INSERT INTO "sessions" (uuid, email, user_id, created_at)
	VALUES ($1, $2, $3, $4)
	RETURNING "id", "uuid", "email", "user_id", "created_at"
	`

	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(
		createUuid(),
		user.Email,
		user.Id,
		time.Now(),
	).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

func GetUserBySessionUuid(session_uuid string) (valid bool, user User) {
	err := Db.QueryRow(`SELECT "user_id" FROM "sessions" WHERE "uuid" = $1`, session_uuid).
		Scan(&user.Id)
	if err != nil {
		valid = false
		return
	}
	if user.Id != 0 {
		valid = true
		Db.QueryRow(`
		SELECT "id", "uuid", "first_name", "last_name", "email", "created_at"
		FROM "users"
		WHERE "id" = $1
		`, user.Id).Scan(&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt)
	}
	return
}

// Delete session from database
func DeleteSessionByUuid(uuid string) (err error) {
	_, err = Db.Exec(`DELETE FROM "sessions" WHERE "uuid" = $1`, uuid)
	return
}

// Delete all sessions from database
func SessionDeleteAll() (err error) {
	_, err = Db.Exec(`DELETE FROM sessions`)
	return
}

// Get the session for an existing user
func (user *User) GetSession() (session Session, err error) {
	row := Db.QueryRow(`SELECT "id", "uuid", "email", "user_id", "created_at" FROM "sessions" WHERE "user_id" = $1`, user.Id)
	err = row.Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}
