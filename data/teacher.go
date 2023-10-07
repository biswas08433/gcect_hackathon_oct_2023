package data

import "log"

func TrendingTeachers() (trending_teachers []User, err error) {
	rows, err := Db.Query(`
	SELECT "id", "uuid", "first_name", "last_name", "email"	FROM "users"
	ORDER BY "avg_rating" DESC
	LIMIT 6
	`)
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		trending_teachers = append(trending_teachers, user)
	}
	rows.Close()
	return
}

func GetTeachersBySubject(subject string) (teachers []User, err error) {
	rows, err := Db.Query(`
	SELECT "id", "uuid", "first_name", "last_name", "email" FROM "users"
	WHERE "id" = (
		SELECT "teacher_id"
		FROM "teacher_subject"
		WHERE "subject_id" = (
			SELECT "id" FROM "subjects"
			WHERE "title" = $1
		)
	)
	`, subject)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.FirstName, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		teachers = append(teachers, user)
	}
	log.Println(teachers)
	return
}
