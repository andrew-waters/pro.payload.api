package deps

import "github.com/PayloadPro/pro.payload.api/configs"

// Config wrapped in a container
type Config struct {
	App *configs.AppConfig
	DB  *configs.DatabaseConfig
}

// Setup the config
func (c Config) Setup() {
	c.App.Setup()
	c.DB.Setup()
}
