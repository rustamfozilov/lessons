package main

import (
	"fmt"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "humo"
	password := "pass"
	dbname := "humo_db"

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dushanbe",
		host, user, password, dbname, port)

	conn, err := gorm.Open(postgresDriver.Open(connString))
	if err != nil {
		log.Printf("%s GetPostgresConnection -> Open error: ", err.Error())
		return
	}

	log.Println("Postgres Connection success: ", host)

	///------------------------------------------------------------------------------

	newUser := User{
		Name:    "TEstNAme",
		UserAge: 99,
		Phone:   "123456789",
		Active:  false,
	}
	SecondUser := User{
		Name:    "TEstNAme",
		UserAge: 99,
		Phone:   "123456789",
		Active:  false,
	}

	var users []*User

	users = append(users, &newUser, &SecondUser)

	log.Println("users", users)

	//err = AddUser(conn, &newUser)
	//if err != nil {
	//	log.Println(err)
	//}

	//err = AddUsers(conn, users)
	//if err != nil {
	//	log.Println(err)
	//}

	//getUser, err := GetUser(conn, 5)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("GET:", *getUser)

	getUsers, err := GetUsers(conn, 99)
	if err != nil {
		log.Println(err)
	}
	log.Println("GET:", getUsers)

	for _, getUser := range getUsers {
		log.Println("*", getUser)
	}
}

func CrateTable(conn *gorm.DB) error {

	sqlQuery := ``

	result := conn.Exec(sqlQuery)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddUser(conn *gorm.DB, newUser *User) error {

	sqlQuery := `insert into users (name, age, phone)		
						values (?, ?, ?);`

	tx := conn.Exec(sqlQuery, newUser.Name, newUser.UserAge, newUser.Phone)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func AddUsers(conn *gorm.DB, users []*User) error {

	sqlQuery := `insert into users (name, age, phone)		
						values (?, ?, ?);`

	for _, user := range users {

		tx := conn.Exec(sqlQuery, user.Name, user.UserAge, user.Phone)
		if tx.Error != nil {
			return tx.Error
		}

	}

	return nil
}

func GetUser(conn *gorm.DB, userId int64) (*User, error) {
	var user User

	sqlQwery := `select *from users where id = ?;`

	tx := conn.Raw(sqlQwery, userId).Scan(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &user, nil
}

func GetUsers(conn *gorm.DB, age int64) ([]*User, error) {
	var users []*User

	sqlQwery := `select *from users where age < ?;`

	tx := conn.Raw(sqlQwery, age).Scan(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

type User struct {
	Id        int64 `json:"id" gorm:"column:id"`
	Name      string
	UserAge   int64 `gorm:"column:age"`
	Phone     string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
