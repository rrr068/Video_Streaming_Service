package comment

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID    uint
	ContentID uint
	Text      string
}
