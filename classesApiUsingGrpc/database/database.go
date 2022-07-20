package database

import (
	"fmt"
	"time"

	pb "github.com/zahraakaraki/MyApp/protos"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func GetClassesFromDB() []*pb.ReturnedType {

	dsn := "host=localhost user=postgres password=karaki217 port=5432 dbname=stc sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("err")
		panic(err)
	}

	sqlDB1, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB1.SetMaxIdleConns(100)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB1.SetMaxOpenConns(5000)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB1.SetConnMaxLifetime(time.Hour)

	var classes []*pb.Class
	var students []*pb.Student
	var teachers []*pb.Teacher
	var result []*pb.ReturnedType

	db.Find(&classes)

	for _, class := range classes {
		returnedtype := &pb.ReturnedType{}
		returnedtype.Classname = class.Name

		db.Where("classid = ?", class.Id).Find(&students)
		db.Where("classid = ?", class.Id).Find(&teachers)

		returnedtype.Student = students
		returnedtype.Teacher = teachers

		result = append(result, returnedtype)

	}

	return result
}
