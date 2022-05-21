package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenDB(ip string, port int, user string, pass string, db string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, ip, port, db)
	log.Printf("dsn=%s", dsn)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
