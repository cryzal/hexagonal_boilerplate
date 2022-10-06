package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type OutletModel struct {
	ID        *int64    `gorm:"column:id;primary_key"`
	Code      *string   `gorm:"column:code"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone_number"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (OutletModel) TableName() string {
	return "outlet"
}
func (_o *OutletModel) BeforeCreate(trx *gorm.DB) (err error) {
	code := uuid.New().String()
	_o.Code = &code
	_o.CreatedAt = time.Now()
	return nil
}
