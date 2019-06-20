package resources

import "github.com/jinzhu/gorm"

type ResourceBase struct {
	DB *gorm.DB
}
