package internal

import (
	"fmt"
	"os"
)

const (
	TokenVarName = "M3O_API_TOKEN" // nolint:gosec
	PortVarName  = "TWS_PORT"
)

type AppConfig struct {
	m3oToken string
	port     string
}

func envOrDefault(varName string, orDefault string) string {
	v := os.Getenv(varName)
	if v != "" {
		return v
	}
	return orDefault
}

func (c *AppConfig) Token() string {
	return c.m3oToken
}

func (c *AppConfig) Port() string {
	return c.port
}

func LoadEnvConfig() (*AppConfig, error) {
	token := os.Getenv(TokenVarName)
	if token == "" {
		return nil, fmt.Errorf("coult not find %s env variable", TokenVarName)
	}

	return &AppConfig{
		m3oToken: token,
		port:     envOrDefault(PortVarName, "8090"),
	}, nil
}
