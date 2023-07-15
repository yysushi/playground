package main_test

import (
	"context"
	"database/sql"
	_ "embed"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yysushi/playground/sql/sqlc/sqlite"
)

//go:embed sqlite/schema.sql
var sqlSceham string

func TestAuthors(t *testing.T) {
	// setup
	source, err := os.CreateTemp("testdata", "sqlite_testdb_*.db")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(source.Name())
	sdb, err := sql.Open("sqlite3", source.Name())
	if err != nil {
		t.Fatal(err)
	}
	_, err = sdb.Exec(sqlSceham)
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()

	db := sqlite.New(sdb)
	// create
	params := sqlite.CreateAuthorParams{
		Name: "mary",
	}
	result, err := db.CreateAuthor(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}
	authorID, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(authorID)
	// get
	author, err := db.GetAuthor(context.Background(), authorID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(author)
}
