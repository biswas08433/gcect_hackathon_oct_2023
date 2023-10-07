package data

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Db *sql.DB
	// ctx context.Context
)

type User struct {
	Id          int
	Uuid        string
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AvgRating   float64
	RatingCount int
	CreatedAt   time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

type Address struct {
	Line1     string
	Line2     string
	Landmark  string
	City      string
	State     string
	Pincode   int
	Country   int
	Latitude  string
	Longitude string
}

var (
	repopulate_database = true
)

func init() {
	var err error
	Db, err = sql.Open("sqlite3", "./teachwise_data.db")
	if err != nil {
		log.Fatalln("db_init_error", err.Error())
	}
	if Db.Ping() != nil {
		log.Fatalln("db_ping_error", err.Error())
	}

	if repopulate_database {
		fake_data, _ := readData("fake_user_data.csv")

		subjects := []string{"literature", "history"}
		for i, record := range fake_data {
			created, err := time.Parse("2006-01-02 15:04:05", record[8])

			if err != nil {
				log.Fatalln("Parse Failed")
			}

			user := User{
				FirstName: record[2],
				LastName:  record[3],
				Email:     record[4],
				Password:  record[5],
				CreatedAt: created,
			}
			log.Println(i, user)
			user.Create()
			user.AddSubject(subjects[0])
		}

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

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
