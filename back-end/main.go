package main

import (
	"webserver/router"
	"webserver/service/mysql"

	"github.com/joho/godotenv"
)

func main() {
	/*loading .env*/
	err_Env := godotenv.Load()
	checkErr(err_Env)
	err_mysql := mysql.Init()
	checkErr(err_mysql)
	router.Init()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
