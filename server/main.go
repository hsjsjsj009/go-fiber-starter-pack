package main

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	log "github.com/sirupsen/logrus"
	"go-fiber-starter-pack/config"
	"go-fiber-starter-pack/db/repository"
	"go-fiber-starter-pack/server/bootstrap"
	"time"
)

var (
	validatorDriver = validator.New()
	uni             *ut.UniversalTranslator
	logFormat       = `{"host":"${host}","pid":"${pid}","time":"${time}","request-id":"${locals:requestid}","status":"${status}","method":"${method}","latency":"${latency}","path":"${path}",` +
		`"user-agent":"${ua}","in":"${bytesReceived}","out":"${bytesSent}"}`
	translator ut.Translator
)

func main() {
	dependencyContainer,envConfig := config.LoadDependenciesAndConfig()

	translatorInit(envConfig.Get("APP_LOCALE"))

	//Register the provider functions
	repository.Register(dependencyContainer)
	dependencyContainer.AddObjectSingleton(validatorDriver)

	defer dependencyContainer.CleanUp()

	app := fiber.New()
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     logFormat + "\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Asia/Jakarta",
	}))

	boot := bootstrap.Bootstrap{
		DepContainer: dependencyContainer,
		Validator:    validatorDriver,
		Translator: translator,
		MainRouter: app,
	}
	boot.RegisterRouters()
	log.Fatal(app.Listen(envConfig.Get("APP_HOST")))

}

func translatorInit(locale string) {
	enT := en.New()
	idT := id.New()
	uni = ut.New(enT, idT)

	transEN, _ := uni.GetTranslator("en")
	transID, _ := uni.GetTranslator("id")

	_ = enTranslations.RegisterDefaultTranslations(validatorDriver, transEN)
	_ = idTranslations.RegisterDefaultTranslations(validatorDriver, transID)

	switch locale {
	case "en":
		translator = transEN
	case "id":
		translator = transID
	}
}
