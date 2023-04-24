package main_test

import (
	"context"
	"database/sql"
	"net/http"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestContext_http(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// explicitly cancel
	cancel()

	// no cancel error
	req, err := http.NewRequestWithContext(ctx, "GET", "http://undefined", nil)
	assert.Nil(t, err)

	// cancel error during client.Do
	client := &http.Client{}
	_, err = client.Do(req)
	// assert.NotNil(t, err)
	// assert.ErrorContains(t, err, ": context canceled")
	assert.ErrorIs(t, err, context.Canceled)
}

func TestContext_sql(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// explicitly cancel
	cancel()

	db, err := sql.Open("mysql", "@tcp(localhost:3306)/dummyDB")
	assert.Nil(t, err)

	// cancel error during sql.DB.QueryContext
	err = db.PingContext(ctx)
	// _, err = db.QueryContext(ctx, "SELECT 1")
	// assert.NotNil(t, err)
	// assert.ErrorContains(t, err, ": context canceled")
	assert.ErrorIs(t, err, context.Canceled)
}
