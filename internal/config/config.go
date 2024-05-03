package config

import "errors"

type databaseConfig struct {
	Username string
	Password string
	Name     string
	Addr     string
	Port     int
}

type Config struct {
	Addr string
	Port int
	Db   databaseConfig
}

type configBuilder struct {
	addr *string
	port *int
	db   struct {
		username *string
		password *string
		name     *string
		addr     *string
		port     *int
	}
}

func (b *configBuilder) Addr(addr string) *configBuilder {
	b.addr = &addr
	return b
}

func (b *configBuilder) Port(port int) *configBuilder {
	b.port = &port
	return b
}

func (b *configBuilder) DbName(name string) *configBuilder {
	b.db.name = &name
	return b
}

func (b *configBuilder) DbUsername(username string) *configBuilder {
	b.db.username = &username
	return b
}

func (b *configBuilder) DbPassword(password string) *configBuilder {
	b.db.password = &password
	return b
}

func (b *configBuilder) DbAddr(addr string) *configBuilder {
	b.db.addr = &addr
	return b
}

func (b *configBuilder) DbPort(port int) *configBuilder {
	b.db.port = &port
	return b
}

func (b *configBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.addr == nil {
		cfg.Addr = "localhost"
	} else {
		cfg.Addr = *b.addr
	}

	if b.port == nil {
		cfg.Port = 8000
	} else if *b.port < 0 {
		return Config{}, errors.New("api port must be non-negative")
	} else {
		cfg.Port = *b.port
	}

	if b.db.username == nil {
		cfg.Db.Username = "postgres"
	} else {
		cfg.Db.Username = *b.db.username
	}

	if b.db.password == nil {
		cfg.Db.Password = ""
	} else {
		cfg.Db.Password = *b.db.password
	}

	if b.db.name == nil {
		cfg.Db.Name = "postgres"
	} else {
		cfg.Db.Name = *b.db.name
	}

	if b.db.addr == nil {
		cfg.Db.Addr = "localhost"
	} else {
		cfg.Db.Addr = *b.db.addr
	}

	if b.db.port == nil {
		cfg.Db.Port = 5432
	} else if *b.db.port < 0 {
		return Config{}, errors.New("database port must be non-negative")
	} else {
		cfg.Db.Port = *b.db.port
	}

	return cfg, nil
}

func ConfigBuilder() *configBuilder {
	return &configBuilder{}
}
