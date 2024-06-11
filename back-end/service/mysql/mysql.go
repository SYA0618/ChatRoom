package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	fmt.Println(password)
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_host := os.Getenv("DB_HOST")
	host := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, db_host, db_port, db_name)
	d, err := sql.Open("mysql", host)
	db = d
	return err
}

func GetDB() *sql.DB {
	return db
}

func Insert(*sql.DB) {

}

// func Read() {
// 	var users []models.User
// 	rows, err := db.Query("SELECT * FROM account")
// 	defer rows.Close()
// 	if err != nil {
// 		fmt.Printf("Query failed,err:%v\n", err)
// 		return
// 	}

// 	for rows.Next() {
// 		var user models.User
// 		err = rows.Scan(&user.Id, &user.User, &user.Password)
// 		if err != nil {
// 			fmt.Printf("Scan failed,err:%v\n", err)
// 			return
// 		}

// 		users = append(users, user)
// 	}
// 	fmt.Println(users)
// }
