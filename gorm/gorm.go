package gorm

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type DBConfig struct {
	DsName   string
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func NewMysql(conf DBConfig) (*sql.DB, error) {
	link := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open(mysql.Open(link), &gorm.Config{})

	if err != nil {
		return nil, errors.New("mysql link failed")
	}
	return db, nil
}

func TestGorm() {
	fmt.Println("============= gorm ===============")
}
