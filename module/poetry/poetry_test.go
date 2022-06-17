package poetry

import (
	_ "embed"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go:embed pyproject.toml
var data []byte

func TestParseToml(t *testing.T) {
	root := &tomlTree{}
	assert.NoError(t, toml.Unmarshal(data, &root.v))
	fmt.Println(root.Get("tool", "poetry", "dependencies").v)
	fmt.Println(root.Get("tool", "poetry", "name").v)
}

//go:embed poetry.lock.py
var __lockData []byte

func TestParsePoetryLock(t *testing.T) {
	root := &tomlTree{}
	assert.NoError(t, toml.Unmarshal(__lockData, &root.v))
	pkgs := root.Get("package").AsArray()
	fmt.Println(pkgs)
}
