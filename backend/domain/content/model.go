package content

import (
	"time"

	"gorm.io/gorm"
	"github.com/example/gin-backend/domain/user"
	"github.com/example/gin-backend/domain/like"
	"github.com/example/gin-backend/domain/comment"
)

type Content struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	UserID    uint           `json:"user_id"`
	User      user.User      `gorm:"foreignKey:UserID" json:"user"` // 投稿者情報（JOIN用）
	Likes     []like.Like    `json:"likes"`                  // いいね一覧
	Comments  []comment.Comment `json:"comments"`            // コメント一覧
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
