package main

import ( 
"fmt"
"errors"
)


type Person3 struct{
	AADHAR string
	PAN	string
	DL	string
	NAME string
}
 var person13 =  Person3{"AD111111111111","PN111111","DL1111111","111111"}
 var person23 =  Person3{"AD222222222222","PN222222","DL2222222","222222"}
 var person33 =  Person3{"AD333333333333","PN333333","DL3333333","333333"}
 var person43 =  Person3{"AD444444444444","PN444444","DL4444444","444444"}

var listOfPeople3 = make(map[string] Person3)

func fetchDetailsFromMap2(aadhar string) (Person3,error){
	var personDummy =  Person3{"XXXXX","XXXXX","XXXXX","XXXXX"}
for id, value := range listOfPeople3 {
		if(id == aadhar){
			return value,nil
		}
	}

	return  personDummy , errors.New("Person with Aadhar not Found")

}

func validateEmployeeDetails2(aadhar,pan,dl,name string) bool{
	
	bufferedChannel := make(chan bool,4)
	
	pesronFound,err := fetchDetailsFromMap2(aadhar)
	if(err !=nil){
		fmt.Println("EMPLOYEE DETAILS NOT FOUND WITH THIS AADHAR")
		return false
	}
	fmt.Println("EMPLOYEE DETAILS FOUND WITH THIS AADHAR")

	go func(){
			if(pesronFound.AADHAR ==aadhar) {
				bufferedChannel <- true
			} else{
				bufferedChannel <- false
				
			}
		}()
		
	go func(){
			if(pesronFound.NAME==name) {
				bufferedChannel <- true
			} else{
				bufferedChannel <- false
				
			}
		}()
	
		
		go func(){
		if(pesronFound.PAN ==pan) {
			bufferedChannel <- true
		} else{
			bufferedChannel <- false
			
		}
		}()

		
	
	go func(){
		if(pesronFound.DL ==dl) {
		bufferedChannel <- true
		} else{
		bufferedChannel <- false

		}
		}()
		
	

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


	listOfPeople3["AD111111111111"] = person13
	listOfPeople3["AD222222222222"] = person23
	listOfPeople3["AD333333333333"] = person33
	listOfPeople3["AD444444444444"] = person43

	validated := validateEmployeeDetails2("AD111111111111","PN111111","DL1111111","111111")
	
	if(validated){
		fmt.Println("Person Details Validated succesully")
	}else{
		fmt.Println("PERSON DETAILS ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	validated1 := validateEmployeeDetails2("AD111111111111","PN111112","DL1111111","111111")

	if(validated1){
		fmt.Println("Person Details Validated succesully")
	}else{
	 fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	

	validated2 := validateEmployeeDetails2("AD111111111112","PN111111","DL1111111","111111")
	if(validated2){
		fmt.Println("Person Details Validated succesully")
	}else{
		fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}

fmt.Println()
	
	validated3 := validateEmployeeDetails2("AD111111111111","PN111112","DL1111112","111111")

	if(validated3){
		fmt.Println("Person Details Validated succesully")
	}else{
	 fmt.Println("PERSON DETAILS PASSED ARE NOT VALID,PLEASE RECHECK DETAILS ENTERED")
	}
	fmt.Println()

	
}
