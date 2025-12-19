package blog

import (
	"fmt"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	dsn := "root:123456@tcp(192.168.3.124:3306)/homework?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
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

func TestCreateTable(t *testing.T) {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		t.Errorf("CreateTable err:%v\n", err)
	}
}

func TestInitializeData(t *testing.T) {
	// 初始化数据
	userIds := make([]int, 0)
	articleIds := make([]int, 0)

	if !hasData(&User{}) {
		for i := 0; i < 2; i++ {
			userId := i + 1
			account := fmt.Sprintf("test%d", i+1)
			db.Create(&User{ID: userId, Account: account, Name: account})
			userIds = append(userIds, userId)
		}
	} else {
		// 查询已有用户
		db.Model(&User{}).Select("id").Find(&userIds)
	}

	if !hasData(&Post{}) {
		for i := 0; i < 4; i++ {
			articleId := i + 1
			userId := i%len(userIds) + 1
			title := fmt.Sprintf("article%d", i+1)
			content := fmt.Sprintf("content%d", i+1)

			db.Create(&Post{UserID: userId, Title: title, Content: content})
			articleIds = append(articleIds, articleId)
		}
	} else {
		db.Model(&Post{}).Select("id").Find(&articleIds)
	}

	if !hasData(&Comment{}) {
		l := len(articleIds)
		for i := 0; i < l*2; i++ {
			db.Create(&Comment{PostID: articleIds[i%l], Content: fmt.Sprintf("comment%d", i+1)})
		}
	}
}

func TestFindUserArticlesAndComments(t *testing.T) {
	var posts []Post
	db.Model(&Post{}).Where("user_id = ?", 1).Preload("Comments").Find(&posts)
	if len(posts) == 0 {
		t.Error("Not found posts")
		return
	}

	fmt.Println("articles:")
	for _, post := range posts {
		fmt.Println(post)
	}
}

func TestFindMostCommentsArticle(t *testing.T) {
	// select * from post
	// where id = (select post_id from (select post_id, count(post_id) from comment group by post_id order by c limit 1) t)
	var post *Post
	db.Model(&Post{}).Where("id = (?)", db.Select("post_id").Table("(?) t",
		db.Table("comment").Select("post_id, count(post_id) c").Group("post_id").Order("c").Limit(1))).Find(&post)
	if post == nil {
		t.Error("Not found post")
		return
	}
	fmt.Println(post)
}

func TestAddArticle(t *testing.T) {
	var oldCount int
	var count int
	post := &Post{UserID: 1, Title: "test add", Content: "test"}
	db.Model(&User{}).Select("article_count").Where("id = ?", post.UserID).Find(&oldCount)
	db.Create(&post)
	db.Model(&User{}).Select("article_count").Where("id = ?", post.UserID).Find(&count)

	if count != oldCount+1 {
		t.Errorf("failed to update article count, expect: %d, actual: %d", oldCount+1, count)
	}
}

func TestUpdateArticleCommentStatus(t *testing.T) {
	var commentStatus string
	db.Where("post_id = ?", 4).Delete(&Comment{PostID: 4})
	db.Model(&Post{}).Select("comment_status").Where("id = ?", 4).Find(&commentStatus)

	if commentStatus != "无评论" {
		t.Errorf("failed to update comment status, expect: 无评论, actual: %s", commentStatus)
	}
}

func hasData(model interface{}) bool {
	var count int64
	db.Model(model).Count(&count)
	return count > 0
}
