package main

import (
	"encoding/json"
	"log"
)

const data = `[{"name" : "shashank", "age":23, "loaction":"kulai honnakatte"},
{"name" : "shashank", "age":23, "loaction":"kulai honnakatte"},
{"name" : "rithu", "age":26, "loaction":"grama"},
{"name" : "jithu", "age":29, "loaction":"kudla"},
{"name" : "motu", "age":21, "loaction":"hosabettu"},
{"name" : "kotu", "age":34, "loaction":"surathkak"},
{"name" : "botu", "age":28, "loaction":"ksna"},
{"name" : "notu", "age":34, "loaction":"bala"},
{"name" : "lootu", "age":39, "loaction":"kaikamba"}]`

func GetData() (people []Person) {
	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		log.Fatal(err)
	}

	return people
}
