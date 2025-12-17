package transaction

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

	code := m.Run()
	os.Exit(code)
}

func TestTransferSuccessful(t *testing.T) {
	tx := db.Begin()
	// 直接回滚事务，以便测试用例可以重复执行
	defer tx.Rollback()

	var account1 Account
	var account2 Account
	tx.Table("t_account").Where("account = ?", "wxq").First(&account1)
	tx.Table("t_account").Where("account = ?", "weq").First(&account2)

	amount := 100.0
	if account1.Balance < amount {
		t.Errorf("Insufficient balance")
		return
	}

	tx.Table("t_account").Where("account = ?", "wxq").Update("balance", account1.Balance-amount)
	tx.Table("t_account").Where("account = ?", "weq").Update("balance", account2.Balance+amount)

	tx.Table("t_account").Where("account = ?", "wxq").First(&account1)
	tx.Table("t_account").Where("account = ?", "weq").First(&account2)

	// 使用新事务写入交易记录
	txRecord := Transaction{From: account1.Account, To: account2.Account, Amount: amount}
	db.Table("t_transaction").Create(&txRecord)

	// 在当前事务中判断结果是否正确
	if !(account1.Balance == 400 && account2.Balance == 600) {
		t.Errorf("Failed to transfer")
		return
	}
}
