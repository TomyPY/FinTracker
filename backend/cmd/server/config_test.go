package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	dbName := "fintracker"
	dbHost := "localhost:3306"
	dbPasswordEnv := "DB_PASSWORD_ENV"
	dbUser := "root"

	myConf, err := NewConfig("")
	cfg := myConf.cfg

	assert.NotNil(t, cfg)
	assert.Equal(t, nil, err)
	assert.Equal(t, dbName, cfg.DB.Name)
	assert.Equal(t, dbHost, cfg.DB.Host)
	assert.Equal(t, dbPasswordEnv, cfg.DB.PasswordEnv)
	assert.Equal(t, dbUser, cfg.DB.User)

}
