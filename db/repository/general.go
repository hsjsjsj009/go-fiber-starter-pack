package repository

import (
	"go-fiber-starter-pack/db"
)

type DbInstance interface {
	RegisterTx(tx db.SQLTx)
	GetWriteInstance() db.SQLInstance
	GetReadInstance() db.SQLInstance
}

type dbInstance struct {
	db db.SQLDb
	tx db.SQLTx
}

func(g *dbInstance) RegisterTx(tx db.SQLTx) {
	g.tx = tx
}

func(g *dbInstance) GetWriteInstance() db.SQLInstance {
	if g.tx != nil {
		return g.tx
	}
	return g.db
}

func(g *dbInstance) GetReadInstance() db.SQLInstance {
	return g.db
}



