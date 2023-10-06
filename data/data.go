package data

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Db *sql.DB
	// ctx context.Context
)

type User struct {
	Id        int
	Uuid      string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./teachwise_data.db")
	if err != nil {
		log.Fatalln("db_init_error", err.Error())
	}
	if Db.Ping() != nil {
		log.Fatalln("db_ping_error", err.Error())
	}
}

func Encrypt(plaintext string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(plaintext)))
}

func createUuid() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
