package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	payday    time.Time
	name      string
	city      string
	dependent int64
}
