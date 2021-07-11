package dao

import (
	"fmt"

	"github.com/IamNator/veren/internal/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Mysql ...
var Mysql *gorm.DB

//Connect creates a connection to mysql database
func Connect(user, password, host, dbName string) (*gorm.DB, error) {
	// configString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassWord, cfg.DBHost, cfg.DBName)
	//conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	conStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbName)

	db, err := gorm.Open("mysql", conStr) // "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"

	if err != nil {
		return nil, err
	}

	logger.Logger.Println("database is open")

	return db, nil
}

// func init() {
// 	c := config.Config
// 	dao, er := Connect(c.DBUser, c.DBPassWord, c.DBHost, c.DBName) //connects to database server

// 	if er != nil {
// 		log.Fatal(er.Error() + " \nunable to connect to database")
// 		return
// 	}

// 	Mysql = dao

// }
