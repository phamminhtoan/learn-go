package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id int
	animalType string
	nickname string
	zone int
	age int
}

func main()  {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/Dino")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	// Query multi row
	rows, err := db.Query("select * from animals")
	defer rows.Close()
	var animals []animal
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)

	// Query a single row
	row := db.QueryRow("select * from Dino.animals where age > 3")
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)


	// Insert a row
	//result, err := db.Exec("insert into Dino.animals (animal_type, nickname, zone, age) values (?, ?, ?, ?)","cat","Ha",5,16 )
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())

	//update a row
	//age := 20
	//id :=2
	//result, err := db.Exec("update Dino.animals set age = ? where id = ?", age,id)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())



}