package repository

import (
	"go-fiber-starter-pack/db"
	"go-fiber-starter-pack/db/repository/model"
	"time"
)

type ExampleRepo interface {
	DbInstance
	ExampleRead() ([]model.Example,error)
	ExampleWrite(*model.Example) error
}

type exampleImplementation struct {
	dbInstance
}

func NewExampleRepo(db db.SQLDb) ExampleRepo {
	return &exampleImplementation{
		dbInstance{db: db},
	}
}

func (e *exampleImplementation) ExampleWrite(_ *model.Example) error {
	_ = e.GetWriteInstance()
	return nil
}

func(e *exampleImplementation) ExampleRead() ([]model.Example,error) {
	_ = e.GetReadInstance()
	return []model.Example{
		{
			ID: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}, nil
}