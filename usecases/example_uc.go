package usecases

import (
	"go-fiber-starter-pack/db"
	"go-fiber-starter-pack/db/repository"
	"go-fiber-starter-pack/package/functioncaller"
	"go-fiber-starter-pack/package/logruslogger"
	"go-fiber-starter-pack/usecases/viewmodel"
)

type ExampleUC interface {
	UC
	ExampleCase() ([]viewmodel.ExampleVM,error)
}

type exampleUCImpl struct {
	ExampleRepo repository.ExampleRepo `bean:"autowired"`
}

func NewExampleUC() *exampleUCImpl {
	return &exampleUCImpl{}
}

func (e *exampleUCImpl) BeginTx(tx db.SQLTx){
	e.ExampleRepo.RegisterTx(tx)
}

func (e *exampleUCImpl) ExampleCase() ([]viewmodel.ExampleVM, error) {
	data,err := e.ExampleRepo.ExampleRead()
	if err != nil {
		logruslogger.Error(err.Error(),functioncaller.PrintFuncName(),"read repo")
		return nil, err
	}
	var output []viewmodel.ExampleVM
	for _,ex := range data {
		output = append(output,viewmodel.ExampleVM{
			ID:ex.ID,
			CreatedAt: ex.CreatedAt,
		})
	}
	return output,nil
}


