package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nguyentrungtin/go-echo-boilerplate/config"
	"github.com/nguyentrungtin/go-echo-boilerplate/model"
)

type (
	User          = model.User
	PasswordReset = model.PasswordReset
)

var (
	db  *gorm.DB
	err error
)

var models = []interface{}{&User{}, &PasswordReset{}}

// Init is used to init the connection to the database
func Init() {
	db, err = gorm.Open("mysql", fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci", config.MYSQL_URL))

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.Set("gorm:table_options", "COLLATE=utf8mb4_unicode_ci")

	if err != nil {
		panic(err)
	}
	fmt.Println("Connect to database successfully")
}

// Session is used to get the database connection session
func Session() *gorm.DB {
	return db
}

// LogMode is used to enable/disable the SQL Command on the stdout
func LogMode(enable bool) *gorm.DB {
	return db.LogMode(enable)
}

// Close is used to close the database connection
func Close() {
	db.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection to database closed!")
	}
}
