package main

import ( 
"fmt"
"errors"
)


type Person struct{
	AADHAR string
	PAN	string
	DL	string
	NAME string
}
 var person1 =  Person{"AD111111111111","PN111111","DL1111111","111111"}
 var person2 =  Person{"AD222222222222","PN222222","DL2222222","222222"}
 var person3 =  Person{"AD333333333333","PN333333","DL3333333","333333"}
 var person4 =  Person{"AD444444444444","PN444444","DL4444444","444444"}

var listOfPeople = make(map[string] Person)

func fetchDetailsFromMap(aadhar string) (Person,error){
	var personDummy =  Person{"XXXXX","XXXXX","XXXXX","XXXXX"}
for id, value := range listOfPeople {
		if(id == aadhar){
			return value,nil
		}
	}

	return  personDummy , errors.New("Person with Aadhar not Found")

}

func validateEmployeeDetails(aadhar,pan,dl,name string) bool{
	
	bufferedChannel := make(chan bool,4)
	
	pesronFound,err := fetchDetailsFromMap(aadhar)
	if(err !=nil){
		fmt.Println("EMPLOYEE DETAILS NOT FOUND WITH THIS AADHAR")
		return false
	}
	fmt.Println("EMPLOYEE DETAILS FOUND WITH THIS AADHAR")

		
		if(pesronFound.AADHAR ==aadhar) {
			bufferedChannel <- true
		} else{
			bufferedChannel <- false
		 	
		}
	
	
		if(pesronFound.NAME==name) {
			bufferedChannel <- true
		} else{
			bufferedChannel <- false
		 	
		}
	

		if(pesronFound.PAN ==pan) {
			bufferedChannel <- true
		} else{
			bufferedChannel <- false
		 	
		}
	

		if(pesronFound.DL ==dl) {
			bufferedChannel <- true
		} else{
			bufferedChannel <- false
		 	
		}
	

	isValidated ,ok := <- bufferedChannel
	if(ok){
		fmt.Println("bufferedChannel is open")
	}else{
		fmt.Println("bufferedChannel is closed")
	}
	
	close(bufferedChannel)
	isValidated =true
	 for v := range bufferedChannel {
        if(!v){
			isValidated =false
			break
		}
     

    }

	
	
	
	

	return isValidated

}
func main() {


	listOfPeople["AD111111111111"] = person1
	listOfPeople["AD222222222222"] = person2
	listOfPeople["AD333333333333"] = person3
	listOfPeople["AD444444444444"] = person4

	validated := validateEmployeeDetails("AD111111111111","PN111111","DL1111111","111111")
	
	if(validated){
		fmt.Println("Person Details Validated succesully")
	}else{
		fmt.Println("PERSON DETAILS ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	validated1 := validateEmployeeDetails("AD111111111111","PN111112","DL1111111","111111")

	if(validated1){
		fmt.Println("Person Details Validated succesully")
	}else{
	 fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	

	validated2 := validateEmployeeDetails("AD111111111112","PN111111","DL1111111","111111")
	if(validated2){
		fmt.Println("Person Details Validated succesully")
	}else{
		fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}

fmt.Println()
	
	validated3 := validateEmployeeDetails("AD111111111111","PN111112","DL1111112","111111")

	if(validated3){
		fmt.Println("Person Details Validated succesully")
	}else{
	 fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	
}
