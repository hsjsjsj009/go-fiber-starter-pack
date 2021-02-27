package gormpsql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

	_ "github.com/lib/pq"
)

// Connection ...
type Connection struct {
	Host                    string
	DbName                  string
	User                    string
	Password                string
	Port                    int
	Location                *time.Location
	DBMaxConnection         int
	DBMAxIdleConnection     int
	DBMaxLifeTimeConnection int
}

// Connect ...
func (c Connection) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
		c.Host,c.User,c.Password,c.DbName,c.Port,c.Location.String())
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	dbSql,err := db.DB()
	if err != nil {
		return nil, err
	}

	dbSql.SetMaxOpenConns(c.DBMaxConnection)
	dbSql.SetMaxIdleConns(c.DBMAxIdleConnection)
	dbSql.SetConnMaxLifetime(time.Duration(c.DBMaxLifeTimeConnection) * time.Second)

	return db, err
}
