package main

import (
	"fmt"
	"strings"
)

type person struct {
	id     string
	fname  string
	lname  string
	age    int
	gender string
}

func populateOuterMap(persons []person) map[string]map[string]person {
	mapsFinal := make(map[string]map[string]person)
	mapsPersonStartingWithOA1 := make(map[string]person)
	mapsPersonStartingWithOA2 := make(map[string]person)

	for _, value := range persons {
		if strings.HasPrefix(value.id, "OA1") {
			mapsPersonStartingWithOA1[value.id] = value
		}
		if strings.HasPrefix(value.id, "OA2") {
			mapsPersonStartingWithOA2[value.id] = value
		}
	}
	mapsFinal["OA1"] = mapsPersonStartingWithOA1
	mapsFinal["OA2"] = mapsPersonStartingWithOA2

	return mapsFinal
}

func findEmployee(personsMap map[string]map[string]person, empcode string, lname string) {
	if strings.HasPrefix(empcode, "OA1") || strings.HasPrefix(empcode, "OA2") {
		fmt.Println("employee  found")

		if strings.HasPrefix(empcode, "OA1") {
			mapstartswithOA1 := personsMap["OA1"]
			for _, p := range mapstartswithOA1 {
				if p.lname == lname {
					fmt.Println("Employee  found Details", p)
				}
			}
		}
		if strings.HasPrefix(empcode, "OA2") {
			mapstartswithOA2 := personsMap["OA2"]
			for _, p := range mapstartswithOA2 {
				if p.lname == lname {
					fmt.Println("Employee  found Details", p)
				}
			}
		}

	} else {
		fmt.Println("employee is of Different zone")
	}

}

func main() {

	persons := []person{
		person{"OA1001", "John", "Smith", 80, "Male"},
		person{"OA2002", "Gini", "Walsh", 70, "Female"},
		person{"OA1003", "Peter", "Ambrose", 80, "Male"},
		person{"OA1004", "Phebe", "Lara", 55, "Female"},
		person{"OA2005", "Riyan", "Lara", 55, "Female"},
		person{"OA2006", "Brian", "Smith", 55, "Female"},
	}
	maps := populateOuterMap(persons)
	fmt.Println(maps)
	findEmployee(maps, "OA1003", "Ambrose")
	findEmployee(maps, "OA2005", "Lara")
	findEmployee(maps, "OA3003", "Ambrose")

}
