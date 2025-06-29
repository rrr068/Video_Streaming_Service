package like

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserID    uint
	ContentID uint
}
