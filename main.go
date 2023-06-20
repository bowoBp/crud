package main

import (
	"crud/modules/user"
	"crud/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	router := gin.New()

	// open connection db
	dbCrud := db.GormMysql(os.Getenv("CRUD_DSN"))

	////check connection
	//checkdb, err := dbCrud.DB()

	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	////ping to database
	//errconn := checkdb.Ping()
	//if err != nil {
	//	log.Fatal(errconn)
	//}

	fmt.Println("database connected..!")

	versionRoute := router.Group("/v1")

	userHandler := user.NewRouter(dbCrud)
	userHandler.Handle(versionRoute)

	errRouter := router.Run(os.Getenv("PORT"))
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
