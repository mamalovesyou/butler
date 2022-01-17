package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfig struct {
	Name   string
	Nested *NestedTestConfig
}

type NestedTestConfig struct {
	Name        string
	ComplexName string
}

var yamlExample = []byte(`
name: butler
nested:
	name: butler
    complexName: heybutler
`)

func setUpTest(t *testing.T) {
	fs := afero.NewMemMapFs()

	err := fs.Mkdir("/etc/butler", 0o777)
	require.NoError(t, err)

	file, err := fs.Create("/etc/butler/config.yml")
	require.NoError(t, err)

	_, err = file.Write(yamlExample)
	require.NoError(t, err)

	file.Close()

	v := viper.New()
	v.SetFs(fs)

	file, err = fs.Open("/etc/butler/config.yml")
	require.NoError(t, err)
	bytes := []byte{}
	_, err = file.Read(bytes)
	fmt.Println(string(bytes))
}

func TestReadConfig(t *testing.T) {

	t.Run("No env set", func(t *testing.T) {
		setUpTest(t)

		cfg := &TestConfig{}
		err := ReadConfig("/etc/butler/config.yml", "", cfg)

		assert.Equal(t, "butler", cfg.Name)
		assert.Equal(t, "butler", cfg.Nested.Name)
		assert.Equal(t, "heybutler", cfg.Nested.ComplexName)
		assert.NoError(t, err)
	})

	t.Run("With env set", func(t *testing.T) {
		setUpTest(t)

		os.Setenv("NESTED_COMPLEX_NAME", "override")

		cfg := &TestConfig{}
		err := ReadConfig("/etc/butler/config.yml", "", cfg)

		assert.Equal(t, "butler", cfg.Name)
		assert.Equal(t, "butler", cfg.Nested.Name)
		assert.Equal(t, "override", cfg.Nested.ComplexName)
		assert.NoError(t, err)
	})
}
