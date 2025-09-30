package main

import (
	"errors"
	"fmt"
)

func funcReturnMapFromSlice(employeecode []int, name []string) (map[int]employee, error) {
	employeeMap := make(map[int]employee)
	if len(employeecode) == 0 {

		return nil, errors.New("No employeecode found to create Map")
	}
	for i := 0; i < len(employeecode); i++ {
		employeeMap[employeecode[i]] = employee{
			employeeCode: employeecode[i],
			name:         name[i],
		}
	}
	return employeeMap, nil
}

type employee struct {
	employeeCode int
	name         string
}

func main() {

	fmt.Println("Map Syntax 1")
	employee := make(map[string]string)
	employee["676001"] = "Doe"
	employee["676002"] = "David"
	employee["676003"] = "Peter"

	for key, value := range employee {
		//fmt.Println(key, value)
		fmt.Printf("Key :: %s, Value :: %s, Memoryaddress :: %d \n", key, value, &value)
	}

	fmt.Println("Map Syntax 2 Initialized")
	employee2 := map[int]string{
		67601: "Peter",
		67602: "John",
	}
	fmt.Println(len(employee2))

	for key, value := range employee2 {
		//fmt.Println(key, value)
		fmt.Printf("Key :: %d, Value :: %s, Memoryaddress :: %d \n", key, value, &value)
	}

	employeenumbers := []int{1, 2, 3}
	name := []string{"John", "David", "Peter"}
	mapreturned, error := funcReturnMapFromSlice(employeenumbers, name)
	if error != nil {
		fmt.Println("Error :", error)
	} else {

		fmt.Println("Value of map returned :", mapreturned)
	}
}
