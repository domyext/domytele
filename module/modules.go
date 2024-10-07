package module

import (
	"go.uber.org/dig"
)

func RegisterModules(container *dig.Container) {
	container.Provide(ProvideStartModule)

	container.Provide(func(
		start *StartModule,
	) []Module {
		return []Module{start}
	})
}
