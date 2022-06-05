package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSetDemo = wire.NewSet(NewDemoService)

// var ProviderSet = wire.NewSet(NewGreeterService)
