package mongodb

type MongoCfg struct {
	URI           string `mapstructure:"uri" validate:"required"`
	Host          string `mapstructure:"host" validate:"required"`
	Port          int    `mapstructure:"port" validate:"required"`
	Username      string `mapstructure:"username" validate:"required"`
	Password      string `mapstructure:"password" validate:"required"`
	DatabaseName  string `mapstructure:"database_name" validate:"required"`
	EnableTracing bool   `mapstructure:"enable_tracing"`
}
