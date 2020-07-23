package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	assert := assert.New(t)

	config := DefaultConfig()
	assert.NotNil(config)
	assert.NotNil(config.Threads)

	config.SetWorkDir("/foo")
	assert.Equal("/foo/config/node_private_key.json", config.NodePrivateKeyFile())

	config.NodePrivateKey = "other/test_key.json"
	assert.Equal("/foo/other/test_key.json", config.NodePrivateKeyFile())

	config.NodePrivateKey = "/new_root/test_key.json"
	assert.Equal("/new_root/test_key.json", config.NodePrivateKeyFile())
}

func TestValidateConfig(t *testing.T) {
	assert := assert.New(t)

	config := DefaultConfig()
	config.SetWorkDir("/foo")
	assert.NoError(config.ValidateBasic())
}

// TODO: complete testing coverage
