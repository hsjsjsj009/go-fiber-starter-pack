package usecases

import "gorm.io/gorm"

type UC interface {
	BeginTx(*gorm.DB)
}
