package data

import (
	"github.com/inoth/toybox-layout/internal/data/user"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(user.NewUserRepo)
