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
	Name   string            `env:"NAME"`
	Nested *NestedTestConfig `envPrefix:"NESTED_"`
}

type NestedTestConfig struct {
	Name        string `env:"NAME"`
	ComplexName string `env:"COMPLEX_NAME"`
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

func TestReadConfigWithNoEnv(t *testing.T) {

	setUpTest(t)

	cfg := &TestConfig{Nested: &NestedTestConfig{}}
	err := ReadConfig("/etc/butler/config.yml", "", cfg)

	assert.Equal(t, "butler", cfg.Name)
	assert.Equal(t, "butler", cfg.Nested.Name)
	assert.Equal(t, "heybutler", cfg.Nested.ComplexName)
	assert.NoError(t, err)
}

func TestReadConfigFromEnv(t *testing.T) {
	os.Clearenv()
	os.Setenv("NAME", "butler")
	os.Setenv("NESTED_NAME", "nested")
	os.Setenv("NESTED_COMPLEX_NAME", "butler")

	cfg := &TestConfig{Nested: &NestedTestConfig{}}
	err := ReadConfig("", "", cfg)

	assert.Equal(t, "butler", cfg.Name)
	assert.Equal(t, "nested", cfg.Nested.Name)
	assert.Equal(t, "butler", cfg.Nested.ComplexName)
	assert.NoError(t, err)
}
