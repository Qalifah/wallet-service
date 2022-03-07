package config

type PostgresConfig struct {
	Database string `yaml:"database"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type BaseConfig struct {
	ServeAddr string          `yaml:"serve_addr"`
	Postgres  *PostgresConfig `yaml:"postgres"`
}
