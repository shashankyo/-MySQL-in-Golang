package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	db_user   = "application"
	db_passwd = "application123"
	db_addr   = "192.168.2.140"
	db_db     = "test"
)

type Person struct {
	Id       int
	Name     string
	Age      int
	Location string
}

func main() {
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_passwd, db_addr, db_db)
	fmt.Println(s)
	db, err := sql.Open("mysql", s)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// err = insertData(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	people, err := getAllData(db)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(people)
	fmt.Println("we have %v people in our db \n", len(people))

	people, err = GetAllAboveAge(db, 30)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(people)
	fmt.Println("we have %v people in our db that are older than 30 \n", len(people))

	err = deleteAllAboveAge(db, 30)
	if err != nil {
		log.Fatal(err)
	}
	people, err = GetAllAboveAge(db, 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(people)
	fmt.Printf("we have %v people in our db older than 30\n", len(people))

	err = updatePersonAge(db, "shashank", 23, "kulai honnakatte")
	if err != nil {
		log.Fatal(err)
	}

	people, err = getAllData(db)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(people)
	fmt.Println("we have %v people in our db \n", len(people))
}

func insertData(db *sql.DB) error {
	people := GetData()
	for _, person := range people {
		q := "INSERT INTO `person` (name,age,location) VALUES (?,?,?)"
		insert, err := db.Prepare(q)
		defer insert.Close()

		if err != nil {
			return err
		}

		_, err := insert.Exec(person.Name, person.Age, person.Location)
		if err != nil {
			return err
		}
	}

	return nil
}

func getAllData(db *sql.DB) (people []Person, err error) {
	resp, err := db.Query("SELECT*FROM `person`")
	defer resp.Close()

	if err != nil {
		return people, err
	}

	for resp.Next() {
		var pPerson Person
		err = resp.Scan(&pPerson.Id, &pPerson.Name, &pPerson.Age, &pPerson.Location)
		if err != nil {
			return people, err
		}

		people = append(people, pPerson)
	}
	return people, nil
}

func GetAllAboveAge(db *sql.DB, age int) {
	q := "SELECT * FROM `person` WHERE `age` > ?"
	resp, err := db.Query(q, age)
	defer resp.Close()

	if err != nil {
		return people, err
	}

	for resp.Next() {
		var pPerson Person
		err = resp.Scan(&pPerson.Id, &pPerson.Name, &pPerson.Age, &pPerson.Location)
		if err != nil {
			return people, err
		}

		people = append(people, pPerson)
	}
	return people, nil
}

func deleteAllAboveAge(db *sql.DB, age int) error {
	q := "DELETE FROM `person` WHERE `age` > ?"
	drop, err := db.Prepare(q)
	defer drop.Close()

	if err != nil {
		return err
	}

	_, err = drop.Exec(age)
	if err != nil {
		return err
	}
	return nil
}

func updatePersonAge(db *sql.DB, name string, age int) {
	q := "UPDATE `person` SET `age` = ? WHERE `name` like ? "
	update, err := db.Prepare(q)
	defer update.Close()

	if err != nil {
		return err
	}

	_, err = update.Exec(age, name)
	if err != nil {
		return err
	}
	return nil
}
