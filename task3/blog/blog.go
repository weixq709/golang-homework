package blog

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int        `gorm:"primaryKey"`
	Account      string     `gorm:"size:20;not null"`
	Name         string     `gorm:"size:20"`
	ArticleCount int        `gorm:"default:0"`
	CreateTime   *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Posts        []Post     `gorm:"foreignKey:UserID;references:ID"`
}

type Post struct {
	ID            int        `gorm:"primaryKey"`
	UserID        int        ``
	Title         string     `gorm:"size:100;not null"`
	Content       string     `gorm:"type:text"`
	CommentStatus string     `gorm:"type:varchar(20)"`
	Comments      []Comment  `gorm:"foreignKey:PostID;references:ID"`
	CreatedAt     *time.Time `gorm:"type:datetime(0);default:CURRENT_TIMESTAMP"`
}

type Comment struct {
	ID      int        `gorm:"primaryKey"`
	PostID  int        ``
	Content string     `gorm:"type:text"`
	Created *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (post *Post) AfterCreate(db *gorm.DB) (err error) {
	// todo 如果更新失败如何处理
	db.Model(&User{}).Where("id = ?", post.UserID).Update("article_count", gorm.Expr("article_count + ?", 1))
	return nil
}

func (comment *Comment) AfterDelete(db *gorm.DB) (err error) {
	var count int64
	db.Model(&Comment{}).Where("post_id = ?", comment.PostID).Count(&count)
	if count == 0 {
		db.Model(&Post{}).Where("id = ?", comment.PostID).Update("comment_status", "无评论")
	}
	return
}
