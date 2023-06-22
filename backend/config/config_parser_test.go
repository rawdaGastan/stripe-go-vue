// Package config for config details
package config

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rightConfig = `
{
	"port": ":3000",
	"stripe": {
			"publisher": "pk",
			"secret": "sk"
	},
	"database": {
			"file": "testing.db"
	},
	"version": "v1"
}
	`

func TestReadConfFile(t *testing.T) {
	t.Run("read config file ", func(t *testing.T) {
		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(rightConfig), 0644)
		assert.NoError(t, err)

		data, err := ReadConfFile(configPath)
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
	})

	t.Run("change permissions of file", func(t *testing.T) {
		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(rightConfig), fs.FileMode(os.O_RDONLY))
		assert.NoError(t, err)

		data, err := ReadConfFile(configPath)
		assert.Error(t, err)
		assert.Empty(t, data)
	})

	t.Run("no file exists", func(t *testing.T) {
		data, err := ReadConfFile("./testing.json")
		assert.Error(t, err)
		assert.Empty(t, data)
	})
}

func TestParseConf(t *testing.T) {

	t.Run("can't unmarshal", func(t *testing.T) {
		config := `{testing}`

		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(config), 0644)
		assert.NoError(t, err)

		_, err = ReadConfFile(configPath)
		assert.Error(t, err)
	})

	t.Run("parse config file", func(t *testing.T) {
		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(rightConfig), 0644)
		assert.NoError(t, err)

		got, err := ReadConfFile(configPath)
		assert.NoError(t, err)

		expected := Configuration{
			Port: ":3000",
			StripeKeys: StripeKeys{
				Publisher: "pk",
				Secret:    "sk",
			},
			Database: DB{
				File: "testing.db",
			},
			Version: "v1",
		}

		assert.NoError(t, err)
		assert.Equal(t, got.Port, expected.Port)
		assert.Equal(t, got.StripeKeys, expected.StripeKeys)
		assert.Equal(t, got.Database, expected.Database)
		assert.Equal(t, got.Version, expected.Version)
	})

	t.Run("no file", func(t *testing.T) {
		_, err := ReadConfFile("config.json")
		assert.Error(t, err)
	})

	t.Run("no server configuration", func(t *testing.T) {
		config :=
			`
{
	"port": ""
}
	`

		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(config), 0644)
		assert.NoError(t, err)

		_, err = ReadConfFile(configPath)
		assert.Error(t, err, "port configuration is required")
	})

	t.Run("no database configuration", func(t *testing.T) {
		config :=
			`
{
	"port": ":3000",
	"database": {
		"file": ""
	}
}
	`

		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(config), 0644)
		assert.NoError(t, err)

		_, err = ReadConfFile(configPath)
		assert.Error(t, err, "database configuration is required")
	})

	t.Run("no stripe configuration", func(t *testing.T) {
		config :=
			`
{
	"port": ":3000",
	"database": {
        "file": "testing.db"
    },
	"stripe": {
			"publisher": "",
			"secret": ""
	}
}
	`

		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(config), 0644)
		assert.NoError(t, err)

		_, err = ReadConfFile(configPath)
		assert.Error(t, err, "stripe configuration is required")
	})

	t.Run("no version configuration", func(t *testing.T) {
		config :=
			`
{
	"port": ":3000",
	"database": {
        "file": "testing.db"
    },
	"stripe": {
			"publisher": "pk",
			"secret": "sk"
	}
	"version": ""
}
	`

		dir := t.TempDir()
		configPath := filepath.Join(dir, "/config.json")

		err := os.WriteFile(configPath, []byte(config), 0644)
		assert.NoError(t, err)

		_, err = ReadConfFile(configPath)
		assert.Error(t, err, "version is required")
	})
}
