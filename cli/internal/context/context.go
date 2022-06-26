package context

import (
	config "github.com/zoomoid/assignments/v1/internal/config"
	zap "go.uber.org/zap"
)

type AppContext struct {
	Logger        *zap.SugaredLogger
	Cwd           string
	Root          string
	Configuration *config.Configuration
}

// Read uses the context's root to read a configmap into the context's struct field
func (c *AppContext) Read() error {
	cfg, err := config.ReadConfigMap(c.Root)
	if err != nil {
		return err
	}
	c.Configuration = cfg
	return nil
}

// Write writes the context's struct field to a file at the context's root
func (c *AppContext) Write() error {
	err := config.WriteConfigMap(c.Configuration, c.Root)
	return err
}

// Clone copies all fields except the logger into a fresh context and returns a reference to it
func (c *AppContext) Clone() *AppContext {
	nc := &AppContext{
		Cwd:           c.Cwd,
		Root:          c.Root,
		Configuration: c.Configuration.Clone(),
		Logger:        c.Logger, // don't actually clone the zap logger, this is fine to be aliased by all contexts
	}
	return nc
}
