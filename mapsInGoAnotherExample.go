package main

import (
	//"errors"
	"errors"
	"fmt"
)

var mapofSeniorCitizen = make(map[string]seniors)

func populateMapOfCitizens(seniorCitizen []seniors) (map[string]seniors, error) {
	if len(seniorCitizen) == 0 {
		return nil, errors.New("No  citizen found")
	}

	mapofSeniorCitizen1 := make(map[string]seniors)

	for _, value := range seniorCitizen {
		mapofSeniorCitizen1[value.id] = value
	}

	return mapofSeniorCitizen1, nil
}

func deletePeopleBelowSixty(allpeople map[string]seniors) (map[string]seniors, error) {

	if len(allpeople) == 0 {
		return nil, errors.New("Nothing in map")
	}
	for id, value := range allpeople {
		if value.age < 60 {
			delete(allpeople, id)
		}

	}
	return allpeople, nil
}

type seniors struct {
	id     string
	name   string
	age    int
	gender string
}

func main() {
	seniorCitizen := []seniors{
		seniors{"OA1001", "John", 80, "Male"},
		seniors{"OA1002", "Gini", 70, "Female"},
		seniors{"OA1003", "Peter", 80, "Male"},
		seniors{"OA1004", "Phebe", 55, "Female"},
	}
	allpeople, error := populateMapOfCitizens(seniorCitizen)
	fmt.Println("List of All Citizens:")
	for key, value := range allpeople {
		fmt.Printf("key: %s, value: %v\n", key, value)
	}
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("List of Senior  Citizens:")
		finalListOfSenior, error1 := deletePeopleBelowSixty(allpeople)
		if error1 != nil {
			fmt.Println(error1)
		}
		for key, value := range finalListOfSenior {
			fmt.Printf("key: %s, value: %v\n", key, value)
		}
	}

}
