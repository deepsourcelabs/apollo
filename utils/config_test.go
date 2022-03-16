package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	rc = &redisConfig{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	lc = &loggingConfig{
		File: "./test.log",
	}

	tc = &timeoutConfig{
		Type:  "minute",
		Value: 5,
	}

	sc = &serverConfig{
		Addr: "localhost:8080",
	}

	config = &Config{
		Redis:   *rc,
		Server:  *sc,
		Logging: *lc,
		Timeout: *tc,
	}
)

func TestGetConfig(t *testing.T) {
	t.Run("must load config", func(t *testing.T) {
		got, err := GetConfig("./testdata/config_test.yaml")
		assert.Nil(t, err)

		want := config

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("must return error if path is invalid", func(t *testing.T) {
		_, err := GetConfig("./testdata/config_test.ml")
		assert.NotNil(t, err)
	})
}

func TestGetTimeout(t *testing.T) {
	t.Run("must return timeout in minutes", func(t *testing.T) {
		got := GetTimeout(config)
		want := 5 * time.Minute

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("must return timeout in seconds", func(t *testing.T) {
		config.Timeout.Type = "second"
		got := GetTimeout(config)
		want := 5 * time.Second

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("must return timeout in microsecond", func(t *testing.T) {
		config.Timeout.Type = "microsecond"
		got := GetTimeout(config)
		want := 5 * time.Microsecond

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want:%v", got, want)
		}
	})

	t.Run("must return default timeout if config not provided", func(t *testing.T) {
		got := GetTimeout(&Config{})
		want := time.Minute

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want:%v", got, want)
		}
	})
}
