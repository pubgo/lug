package registry

var Name = "registry"

type Cfg struct {
	Driver string `json:"driver"`
}

func GetDefaultCfg() Cfg {
	return Cfg{
		Driver: "mdns",
	}
}
