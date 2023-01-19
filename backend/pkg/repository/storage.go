package repository

import (
	"fmt"
	"os"

	"D:/project/GoSelfLearning/school_management_system/backend/pkg/models/studentModel"
	"backend/pkg/models/teacherModel"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Intialmigration() {
	err = godotenv.Load(".env")
	// Loading Environment Variables
	// dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//Database Connection String
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, password, dbName, dbPort)
	// DataBase connection
	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	// defer db.Close()
	if err != nil {
		fmt.Println("Error Connecting to dataBase")
		panic(err)

	} else {
		fmt.Println("Succesfully Connected to database")
	}
	db.AutoMigrate(&studentModel.Student{})
	db.AutoMigrate(&teacherModel.Teacher{})
	db.AutoMigrate(&studentModel.AttendanceStudent{})
	db.AutoMigrate(&teacherModel.AttendanceTeacher{})
}
