package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/smtp"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("cannot open database", db)
	}

	from := "SENDER EMAIL ADDRESS"
	auth := smtp.PlainAuth("", from, "PASSWORD", "localhost")

	err = smtp.SendMail("localhost:1025", auth, from, []string{"RECEIVER MAIL"}, []byte("YOUR MESSAGE"))
	if err != nil {
		log.Println("host")
	}
	fmt.Println("Email sent successfully")
}
