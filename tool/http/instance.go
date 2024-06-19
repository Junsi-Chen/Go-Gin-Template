package http

type Instance struct {
	Mode         string `toml:"mode"`
	Address      string `toml:"address"`
	CorsUrl      string `toml:"cors_url"`
	Port         string `toml:"port"`
	ReadTimeout  int    `toml:"readTimeout"`
	WriteTimeout int    `toml:"writeTimeout"`
}
