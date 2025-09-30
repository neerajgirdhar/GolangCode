package main

import "fmt"

type vehicle interface {
	vehicleRegistrationValidTill() int
	insurancePremium() int
}

type vehicleDetails interface {
	printVehicleDetails()
}

type car struct {
	name       string
	ageofcar   int
	ccofcar    int
	typeoffuel string
}

func (car car) printVehicleDetails() {
	fmt.Printf("Vehicle Details  name of Vehicle : %s age of Vehicle %d cc  of Vehicle %d fuel type of Vehicle %s \n", car.name, car.ageofcar, car.ccofcar, car.typeoffuel)

}

func (car car) vehicleRegistrationValidTill() int {
	if car.typeoffuel == "Petrol" {
		return 15 - car.ageofcar
	}
	if car.typeoffuel == "Diesel" {
		return 10 - car.ageofcar
	}
	return 15 - car.ageofcar
}

func (car car) insurancePremium() int {
	if car.ccofcar > 1500 {
		return 250000 - car.ageofcar*10000
	}
	if car.ccofcar <= 1500 {
		return 150000 - car.ageofcar*10000
	}
	return 250000 - car.ageofcar*10000
}

type truck struct {
	name       string
	ageoftruck int
	ccoftruck  int
	typeoffuel string
}

func (truck truck) vehicleRegistrationValidTill() int {
	if truck.typeoffuel == "Petrol" {
		return 15 - truck.ageoftruck
	}
	if truck.typeoffuel == "Diesel" {
		return 10 - truck.ageoftruck
	}
	return 15 - truck.ageoftruck
}

func (truck truck) insurancePremium() int {
	if truck.ccoftruck > 3500 {
		return 700000 - truck.ageoftruck*10000
	}
	if truck.ccoftruck <= 3500 {
		return 500000 - truck.ageoftruck*10000
	}
	return 750000 - truck.ageoftruck*10000
}

func (truck truck) printVehicleDetails() {
	fmt.Printf("Vehicle Details  name of Vehicle : %s age of Vehicle %d cc  of Vehicle %d fuel type of Vehicle %s \n", truck.name, truck.ageoftruck, truck.ccoftruck, truck.typeoffuel)

}
func test(vehicle2 vehicle, vehicleDetails vehicleDetails) {
	vehicleDetails.printVehicleDetails()
	fmt.Printf("Registration Valid till %d years \n", vehicle2.vehicleRegistrationValidTill())
	fmt.Printf("Insurance Premium for this year %d  \n", vehicle2.insurancePremium())
}

func isCar(vehicle2 vehicle) (bool, string) {
	car, ok := vehicle2.(car)
	return ok, car.name
}

func whatTypeOfVehicle(vehicle2 vehicle) string {
	switch v := vehicle2.(type) {

	case car:
		{
			fmt.Printf("%t \n", v)
			return "Car"
		}
	case truck:
		return "Truck"
	default:
		return "Unknown"
	}

}
func main() {
	fmt.Println()
	fmt.Println("Example of Multiple interfaces")

	test(car{"creata", 2, 1500, "Petrol"}, car{"creata", 2, 1500, "Petrol"})
	fmt.Println()

	car1 := car{"XUV", 3, 2200, "Diesel"}
	test(car1, car1)
	fmt.Println()

	truck1 := truck{"TATA TRUCK", 3, 3700, "Diesel"}
	test(truck1, truck1)
	fmt.Println()

	truck2 := truck{"RAM TRUCK", 5, 3200, "Petrol"}
	test(truck2, truck2)
	fmt.Println()

	carortruck0, name0 := isCar(car1)
	if carortruck0 {
		fmt.Printf("Its a car and name of car is %s\n", name0)
	} else {
		fmt.Println("Its not a car")
	}

	carortruck, name := isCar(truck1)
	if carortruck {
		fmt.Printf("Its a car and name of car is %s\n", name)
	} else {
		fmt.Println("Its not a car")
	}

	carortruck2, name2 := isCar(truck2)
	if carortruck2 {
		fmt.Printf("Its a car and name of car is %s\n", name2)
	} else {
		fmt.Println("Its not a car")
	}

	fmt.Println(whatTypeOfVehicle(car1))
	fmt.Println(whatTypeOfVehicle(truck1))
	fmt.Println(whatTypeOfVehicle(truck2))
}
