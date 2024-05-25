package main_test

import (
	"context"
	"database/sql"
	"errors"
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

func TestContext_childcancel(t *testing.T) {
	parent, parentCancel := context.WithCancel(context.Background())
	child, childCancel := context.WithCancel(parent)
	select {
	case <-parent.Done():
		t.Fail()
	case <-child.Done():
		t.Fail()
	default:
	}
	childCancel()
	select {
	case <-parent.Done():
		t.Fail()
	case <-child.Done():
	default:
		t.Fail()
	}
	select {
	case <-parent.Done():
		t.Fail()
	case <-child.Done():
	default:
		t.Fail()
	}
	parentCancel()
	select {
	case <-parent.Done():
	default:
		t.Fail()
	}
}

func TestContext_parentcancel(t *testing.T) {
	parent, parentCancel := context.WithCancel(context.Background())
	child, childCancel := context.WithCancel(parent)
	select {
	case <-parent.Done():
		t.Fail()
	case <-child.Done():
		t.Fail()
	default:
	}
	parentCancel()
	select {
	case <-child.Done():
	default:
		t.Fail()
	}
	select {
	case <-parent.Done():
	default:
		t.Fail()
	}
	childCancel()
}

func TestContext_cause(t *testing.T) {
	err := errors.New("original error")
	ctx, cancelCauseFunc := context.WithCancelCause(context.Background())
	cancelCauseFunc(err)
	assert.ErrorIs(t, context.Cause(ctx), err)
	assert.NotErrorIs(t, context.Cause(ctx), context.Canceled)
	assert.NotErrorIs(t, ctx.Err(), err)
	assert.ErrorIs(t, ctx.Err(), context.Canceled)

	parentCtx, cancelFunc := context.WithCancel(context.Background())
	ctx2, cancelCauseFunc2 := context.WithCancelCause(parentCtx)
	defer cancelCauseFunc2(nil)
	cancelFunc()
	assert.ErrorIs(t, ctx2.Err(), context.Canceled)
	assert.ErrorIs(t, context.Cause(ctx2), context.Canceled)

	ctx3, cancelCauseFunc3 := context.WithCancelCause(context.Background())
	cancelCauseFunc3(nil)
	assert.ErrorIs(t, ctx3.Err(), context.Canceled)
	assert.ErrorIs(t, context.Cause(ctx3), context.Canceled)
	assert.NotErrorIs(t, context.Cause(ctx3), nil)

}
