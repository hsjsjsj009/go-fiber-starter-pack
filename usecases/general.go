package usecases

import "go-fiber-starter-pack/db"

type UC interface {
	BeginTx(db.SQLTx)
}
