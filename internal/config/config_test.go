package config

import (
	"reflect"
	"testing"
)

type ConfigTestCase struct {
	name string
	want Config
}

var DefaultValues ConfigTestCase = ConfigTestCase{
	name: "Config contains all default values",
	want: Config{
		Addr: "localhost",
		Port: 8000,
		Db: databaseConfig{
			Username: "postgres",
			Password: "",
			Name:     "postgres",
			Addr:     "localhost",
			Port:     5432,
		},
	},
}

var ConfiguredValues ConfigTestCase = ConfigTestCase{
	name: "Config contains all set values",
	want: Config{
		Addr: "127.0.0.1",
		Port: 1234,
		Db: databaseConfig{
			Username: "username",
			Password: "password",
			Name:     "name",
			Addr:     "127.0.0.2",
			Port:     5678,
		},
	},
}

func TestBuild(t *testing.T) {
	conf, err := ConfigBuilder().Build()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(conf, DefaultValues.want) {
		t.Fatalf("%v - want: %v, got: %v", DefaultValues.name, DefaultValues.want, conf)
	}

	conf, err = ConfigBuilder().
		Addr("127.0.0.1").
		Port(1234).
		DbUsername("username").
		DbPassword("password").
		DbName("name").
		DbAddr("127.0.0.2").
		DbPort(5678).
		Build()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(conf, ConfiguredValues.want) {
		t.Fatalf("%v - want: %v, got: %v", ConfiguredValues.name, ConfiguredValues.want, conf)
	}
}
