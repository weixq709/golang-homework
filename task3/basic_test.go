package task3

import (
	"fmt"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	dsn := "root:123456@tcp(192.168.3.124:3306)/homework?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("TestMain err:%v\n", err)
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("TestMain err:%v\n", err)
		return
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db := db.Begin()
	code := m.Run()
	db.Rollback()

	os.Exit(code)
}

func TestCreation(t *testing.T) {
	err := db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"}).Error
	if err != nil {
		t.Errorf("Create student err:%v\n", err)
	}
}

func TestAddDuplicateRecords(t *testing.T) {
	err := db.Create(&Student{ID: 1, Name: "张三", Age: 20, Grade: "三年级"}).Error
	if err == nil {
		t.Errorf("Duplicate record:%v\n", err)
	}
}

func TestFindStudents(t *testing.T) {
	var students []Student
	err := db.Model(&Student{}).Where("age > ?", 18).Find(&students).Error
	if err != nil {
		t.Errorf("Find student err:%v\n", err)
		return
	}
	fmt.Println(students)
}

func TestUpdateStudent(t *testing.T) {
	err := db.Model(&Student{}).Where("name = ?", "张三").Update("Grade", "四年级").Error
	if err != nil {
		t.Errorf("Update student err:%v\n", err)
	}
}

func TestDeleteStudents(t *testing.T) {
	err := db.Where("age < ?", 15).Delete(&Student{}).Error
	if err != nil {
		t.Errorf("Delete student err:%v\n", err)
	}
}
