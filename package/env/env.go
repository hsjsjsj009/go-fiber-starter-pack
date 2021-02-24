package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Data struct {
	EnvMap map[string]string
}

func(d *Data) GetWithDefault(key,defaultValue string) string {
	if data,ok := d.EnvMap[key];ok {
		return data
	}
	return defaultValue
}

func(d *Data) Get(key string) string {
	return d.GetWithDefault(key,"")
}

func GetFromEnvWithDefault(key, defaultValue string) string {
	data := os.Getenv(key)
	if data == "" {
		data = defaultValue
	}
	return defaultValue
}

func loadDotEnv(file ...string) (res map[string]string) {
	res, err := godotenv.Read(file...)
	if err != nil && len(file) > 0 {
		log.Fatal("Error loading .env file")
	}
	return res

}

func LoadAllEnv(dotEnvFile ...string) *Data {
	envMap := loadDotEnv(dotEnvFile...)
	listEnv := os.Environ()
	for _,env := range listEnv {
		splitEnv := strings.Split(env,"=")
		envMap[splitEnv[0]] = splitEnv[1]
	}
	return &Data{EnvMap: envMap}
}
