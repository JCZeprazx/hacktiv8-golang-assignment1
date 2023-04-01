package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
)

type Person struct {
	ID     int
	Nama   string
	Alamat string
	Job    string
	Alasan string
}

func findPerson(name string, datas []Person) (Person, error) {
	for _, value := range datas {
		if value.Nama == name {
			return value, nil
		}
	}
	return Person{}, errors.New("Person not found")
}

func main() {
	student1 := Person{0, "John", "Jalan Dharmawangsa", "Mahasiswa", "Karena golang terbaik"}
	student2 := Person{1, "Abi", "Jalan Pradah", "Mahasiswa", "Mencoba hal baru"}
	student3 := Person{2, "Rega", "Jalan Putat Gede", "Mahasiswa", "Mencoba teknologi golang"}

	listStudent := []Person{student1, student2, student3}
	person := os.Args

	search, err := findPerson(person[1], listStudent)

	if err != nil {
		log.Fatal(err.Error())
	}

	values := reflect.ValueOf(search)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(types.Field(i).Name, "\t :", values.Field(i))
	}
}
