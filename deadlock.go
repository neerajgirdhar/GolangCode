package main

import ( 
"fmt"
"errors"
)


type Person1 struct{
	AADHAR string
	PAN	string
	DL	string
	NAME string
}
 var person11 =  Person1{"AD111111111111","PN111111","DL1111111","111111"}
 var person21 =  Person1{"AD222222222222","PN222222","DL2222222","222222"}
 var person31 =  Person1{"AD333333333333","PN333333","DL3333333","333333"}
 var person41 =  Person1{"AD444444444444","PN444444","DL4444444","444444"}

var listOfPeople1 = make(map[string] Person1)

func fetchDetailsFromMap1(aadhar string) (Person1,error){
	var personDummy =  Person1{"XXXXX","XXXXX","XXXXX","XXXXX"}
for id, value := range listOfPeople1 {
		if(id == aadhar){
			return value,nil
		}
	}

	return  personDummy , errors.New("Person with Aadhar not Found")

}

func validateEmployeeDetails1(aadhar,pan,dl,name string) bool{
	
	nameChannel := make(chan bool)
	panChannel := make(chan bool)
	aadharChannel := make(chan bool)
	dlChannel := make(chan bool)
	
	 pesronFound,err := fetchDetailsFromMap1(aadhar)
	if(err !=nil){
		fmt.Println("EMPLOYEE DETAILS NOT FOUND WITH THIS AADHAR")
		return false
	}
	fmt.Println("EMPLOYEE DETAILS FOUND WITH THIS AADHAR")

	
		if(pesronFound.AADHAR ==aadhar) {
			aadharChannel <- true
		} else{
			aadharChannel <- false
		 	
		}


		if(pesronFound.NAME==name) {
			nameChannel <- true
		} else{
			nameChannel <- false
		 	
		}


		if(pesronFound.PAN ==pan) {
			panChannel <- true
		} else{
			panChannel <- false
		 	
		}
	
	
		if(pesronFound.DL ==dl) {
			dlChannel <- true
		} else{
			dlChannel <- false
		 	
		}
	

	isAadharValidated := <- aadharChannel
	isNameValidated := <- nameChannel
	isDLValidated := <- dlChannel
	isPanValidated := <- panChannel

	if(!isNameValidated){
		fmt.Println("Name Validation Failed")
	}

	if(!isDLValidated){
		fmt.Println("DL Validation Failed")
	}

	if(!isPanValidated){
		fmt.Println("PAN Validation Failed")
	}
	

	return isAadharValidated&&isNameValidated&&isDLValidated&&isPanValidated

}
func main() {


	listOfPeople1["AD111111111111"] = person11
	listOfPeople1["AD222222222222"] = person21
	listOfPeople1["AD333333333333"] = person31
	listOfPeople1["AD444444444444"] = person41

	validated := validateEmployeeDetails1("AD111111111111","PN111111","DL1111111","111111")
	
	if(validated){
		fmt.Println("Person Details Validated succesully")
	}else{
		fmt.Println("PERSON DETAILS ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	validated1 := validateEmployeeDetails1("AD111111111111","PN111112","DL1111111","111111")

	if(validated1){
		fmt.Println("Person Details Validated succesully")
	}else{
	 fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	

	validated2 := validateEmployeeDetails1("AD111111111112","PN111111","DL1111111","111111")
	if(validated2){
		fmt.Println("Person Details Validated succesully")
	}else{
		fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}

fmt.Println()
	
	validated3 := validateEmployeeDetails1("AD111111111111","PN111112","DL1111112","111111")

	if(validated3){
		fmt.Println("Person Details Validated succesully")
	}else{
	 fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

}
