package db

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/nguyentrungtin/go-echo-boilerplate/config"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

// AutoMigrate is used to auto migrate new scheme to the database using gorm
func AutoMigrate() *gorm.DB {
	// NOTE: DONT RUN AUTOMIGRATE USING models..., it'll make migration not correct! Run migrate each model individually instead
	//return db.AutoMigrate(models...) // DONT USE THIS

	// Because gorm will create many-to-many table automatically, so, if we want to create many-to-many table with custom field,
	// we need to migrate first, then migrate normal table later

	db.AutoMigrate(&User{})
	db.AutoMigrate(&PasswordReset{})

	return db
}

// Migrate is used to migrate database using go-migrate
func Migrate() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_NAME))
	if err != nil {
		fmt.Println(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		fmt.Println(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Println(err)
	}

	// Because `key` is a reserved word in SQL and we cannot use backtick in struct field.
	// So we need to exec raw SQL to create table
	//m.Migrate(20200204235132)

	// err = m.Force(20200204235640)
	lib.Handle(err)

	err = m.Up()
	lib.Handle(err)
}
