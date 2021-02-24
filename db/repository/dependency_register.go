package repository

import "github.com/hsjsjsj009/go-beans"

func Register(container *beans.ProviderContainer) {
	container.AddProvider(NewExampleRepo)
}
