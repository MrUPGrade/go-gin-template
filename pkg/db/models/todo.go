package pkg

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title       string `binding:"required"`
	Description string
	Done        bool
}
