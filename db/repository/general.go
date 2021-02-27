package repository

import (
	"gorm.io/gorm"
)

type DbInstance interface {
	RegisterTx(tx *gorm.DB)
}

type dbInstance struct {
	db *gorm.DB
	tx *gorm.DB
}

func(g *dbInstance) RegisterTx(tx *gorm.DB) {
	g.tx = tx
}

func(g *dbInstance) GetWriteInstance() *gorm.DB {
	if g.tx != nil {
		return g.tx
	}
	return g.db
}

func(g *dbInstance) GetReadInstance() *gorm.DB {
	return g.db
}



