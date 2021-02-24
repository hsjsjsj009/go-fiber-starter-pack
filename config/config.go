package config

import (
	"database/sql"
	"github.com/hsjsjsj009/go-beans"
	log "github.com/sirupsen/logrus"
	"go-fiber-starter-pack/db"
	"go-fiber-starter-pack/package/env"
	"go-fiber-starter-pack/package/jwe"
	"go-fiber-starter-pack/package/jwt"
	"go-fiber-starter-pack/package/postgresql"
	"go-fiber-starter-pack/package/str"
)

type Configs struct {
	JweCred     *jwe.Credential
	JwtCred     *jwt.Credential
	DB          *sql.DB
	EnvConfig   *env.Data
}

func LoadDependenciesAndConfig() (res *beans.ProviderContainer,envConfig *env.Data) {

	res = beans.NewContainer()
	envConfig = env.LoadAllEnv("./.env")

	dbConn := postgresql.Connection{
		Host:                    envConfig.EnvMap["DATABASE_HOST"],
		DbName:                  envConfig.EnvMap["DATABASE_DB"],
		User:                    envConfig.EnvMap["DATABASE_USER"],
		Password:                envConfig.EnvMap["DATABASE_PASSWORD"],
		Port:                    str.StringToInt(envConfig.EnvMap["DATABASE_PORT"]),
		SslMode:                 envConfig.EnvMap["DATABASE_SSL_MODE"],
		DBMaxConnection:         str.StringToInt(envConfig.EnvMap["DATABASE_MAX_CONNECTION"]),
		DBMAxIdleConnection:     str.StringToInt(envConfig.EnvMap["DATABASE_MAX_IDLE_CONNECTION"]),
		DBMaxLifeTimeConnection: str.StringToInt(envConfig.EnvMap["DATABASE_MAX_LIFETIME_CONNECTION"]),
	}

	dbPsql,err := dbConn.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	res.AddProviderSingleton(func() (db.SQLDb,beans.CleanUpFunc) {
		return dbPsql, func() {
			dbPsql.Close()
		}
	})

	res.AddProvider(func() (db.SQLTx,error) {
		return dbPsql.Begin()
	})

	//res.AddProviderSingleton(func() jwe.Credential {
	//	return jwe.Credential{
	//		KeyLocation: envConfig.EnvMap["APP_PRIVATE_KEY_LOCATION"],
	//		Passphrase:  envConfig.EnvMap["APP_PRIVATE_KEY_PASSPHRASE"],
	//	}
	//})

	//res.AddProviderSingleton(func() jwt.Credential {
	//	return jwt.Credential{
	//		Secret:           envConfig.EnvMap["TOKEN_SECRET"],
	//		ExpSecret:        str.StringToInt(envConfig.EnvMap["TOKEN_EXP_SECRET"]),
	//		RefreshSecret:    envConfig.EnvMap["TOKEN_REFRESH_SECRET"],
	//		RefreshExpSecret: str.StringToInt(envConfig.EnvMap["TOKEN_EXP_REFRESH_SECRET"]),
	//	}
	//})
	return
}