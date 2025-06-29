package user

import {
	"gorm.io/gorm"
	"github.com/example/gin-backend/domain/content"
	"github.com/example/gin-backend/domain/like"
	"github.com/example/gin-backend/domain/comment"
	"github.com/example/gin-backend/domain/follow"
}

type User struct {
	gorm.Model
	ID        	uint   `gorm:"primaryKey" json:"id"`
	Name        string
	Email       string `gorm:"unique"`
	Description string
	Password    string

	Contents []content.Content
	Likes    []like.Like
	Comments []comment.Comment

	Followers []follow.Follow `gorm:"foreignKey:FollowedID"`
	Following []follow.Follow `gorm:"foreignKey:FollowerID"`
}
