package singolang

import (
	"fmt"
	"os"
)

const senv string = "SINGULARITYENV"

type EnvOptions struct {
	EnvVars map[string]string
	PrependPath []string
	AppendPath []string
	ReplacePath	string
}

func DefaultEnvOptions() *EnvOptions {
	return &EnvOptions {
		EnvVars: make(map[string]string),
		PrependPath: []string{},
		AppendPath: []string{},
	}
}

func (opts *EnvOptions) processEnvVars() error {
	var err error
	for k, v := range opts.EnvVars {
		varName := fmt.Sprintf("SINGULARITYENV_%s", k)
		err = os.Setenv(varName, v)
	}
	return err
}

func (opts *EnvOptions) processPathMod() error {
	var err error
	for _, prepath := range opts.PrependPath {
		err = prependPath(prepath)
	}
	for _, apppath := range opts.AppendPath {
		err = appendPath(apppath)
	} 
	return err
}

func prependPath(path string) error {
	err := os.Setenv(fmt.Sprintf("%s_PREPEND_PATH", senv), path)
	return err
}

func appendPath(path string) error {
	err := os.Setenv(fmt.Sprintf("%s_APPEND_PATH", senv), path)
	return err
}

func (opts *EnvOptions) unsetAll() {
	for k := range opts.EnvVars {
		varName := fmt.Sprintf("%s_%s", senv, k)
		os.Unsetenv(varName)
	}
	os.Unsetenv(fmt.Sprintf("%s_PREPEND_PATH", senv))
	os.Unsetenv(fmt.Sprintf("%s_APPEND_PATH", senv))
}