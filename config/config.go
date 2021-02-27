package config

import (
	"github.com/hsjsjsj009/go-beans"
	log "github.com/sirupsen/logrus"
	"go-fiber-starter-pack/package/env"
	"go-fiber-starter-pack/package/gormpsql"
	"go-fiber-starter-pack/package/str"
	"gorm.io/gorm"
	"time"
)

func LoadDependenciesAndConfig() (res *beans.ProviderContainer,envConfig *env.Data) {

	res = beans.NewContainer()
	envConfig = env.LoadAllEnv("./.env")

	locZone := envConfig.GetWithDefault("TIME_ZONE","Asia/Jakarta")
	loc,err := time.LoadLocation(locZone)
	if err != nil {
		log.Fatal(err.Error())
	}
	time.Local = loc

	dbConn := gormpsql.Connection{
		Host:                    envConfig.EnvMap["DATABASE_HOST"],
		DbName:                  envConfig.EnvMap["DATABASE_DB"],
		User:                    envConfig.EnvMap["DATABASE_USER"],
		Password:                envConfig.EnvMap["DATABASE_PASSWORD"],
		Port:                    str.StringToInt(envConfig.EnvMap["DATABASE_PORT"]),
		DBMaxConnection:         str.StringToInt(envConfig.EnvMap["DATABASE_MAX_CONNECTION"]),
		DBMAxIdleConnection:     str.StringToInt(envConfig.EnvMap["DATABASE_MAX_IDLE_CONNECTION"]),
		Location: 				 loc,
		DBMaxLifeTimeConnection: str.StringToInt(envConfig.EnvMap["DATABASE_MAX_LIFETIME_CONNECTION"]),
	}

	dbPsql,err := dbConn.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	res.AddProviderSingleton(func() *gorm.DB {
		return dbPsql
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