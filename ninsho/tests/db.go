package tests

import (
	"flag"
	"ninsho/internal/db"
	"ninsho/internal/users"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/v10"
)

var testDB *pg.DB

func createSchema(db *pg.DB, model interface{}) {
	err := db.Model(model).CreateTable(&pg.CreateTableOptions{
		Temp: true,
	})
	if err != nil {
		panic(err)
	}
}

func init() {
	testDB = db.Start(&pg.Options{
		Addr:     flag.String("db_addr", "", "address of the db"),
		User:     flag.String("db_user", "postgres", "user"),
		Password: flag.String("db_pass", "postgres", "passowrd of the db"),
		Database: flag.String("db_name", "ninsho_test", "name of the db"),
	})

	createSchema(testDB, (*users.User)(nil))

	_, err := db.Model(&users.User{
		Email:    "eta@hotmail.com",
		Password: "lerolero123",
	}).Insert()

	if err != nil {
		panic(err)
	}

}
