//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitInjector() (*Injector, func(), error) {
	wire.Build(
		InitDB,
		ModelSet,
		ServiceSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
