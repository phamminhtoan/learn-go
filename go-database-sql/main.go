package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type user struct {
	id   int
	name string
}

type application struct {
	db *sql.DB
}

func (app *application) getUser(id int) {
	rows, err := app.db.Query("select id, name from users where id = ?", id)
	check(err)
	defer rows.Close()

	user := &user{}
	for rows.Next() {
		err := rows.Scan(&user.id, &user.name)
		check(err)
		log.Println(user.id, user.name)
	}
	err = rows.Err()
	check(err)
	fmt.Println(user)

}

func (app *application) getAllUser() {
	users := []user{}
	user := user{}
	rows, err := app.db.Query("select * from users")
	check(err)
	for rows.Next() {
		err := rows.Scan(&user.id, &user.name)
		check(err)
		users = append(users, user)
	}
	fmt.Println(users)
}

func (app *application) insertUser(id, name string) {
	stmt := `INSERT INTO users VALUES(?,?)`
	res, err := app.db.Exec(stmt, id, name)
	check(err)
	lastID, err := res.LastInsertId()
	check(err)
	fmt.Println(lastID)
}

func main() {
	db, err := sql.Open("mysql", "root:123456789@/hello")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

	app := &application{db: db}
	// app.getUser(0)
	app.insertUser("3", "minh thach")
	app.getAllUser()

}
