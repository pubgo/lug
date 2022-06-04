package app

import "github.com/pubgo/xerror"

type Cfg struct {
	Debug     bool   `yaml:"debug"`
	Addr      string `yaml:"addr"`
	Advertise string `yaml:"advertise"`
}

func (c *Cfg) Check() (err error) {
	defer xerror.RecoverErr(&err)

	if c.Addr == "" {
		c.Addr = ":8080"
	}

	return
}
