package test

import (
	"github.com/stretchr/testify/assert"
	"go_restful_api/simple"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, cleanup := simple.InitializedConnection("Database")
	assert.NotNil(t, connection)

	cleanup()
}
