package main

import (
	"fmt"
	"strings"
)
func main() {

	
	
	 firstString := "THIS IS NEERAJ"
	 secondString := "i work in infosys"
     firstString =  strings.Replace(firstString,"E","e",2)


   slice := []string{firstString , secondString}
   firstString = strings.Join(slice,".")
firstString =  strings.Replace(firstString,"I","K",3)

		fmt.Printf(firstString)
}
