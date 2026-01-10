package matcher

import (
	"github.com/vincbro/suddig/configs"
)

type Matcher struct {
	config configs.Config
}

func New(config configs.Config) *Matcher {
	return &Matcher{config: config}
}
