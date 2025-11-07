package main
import "fmt"
func removeEmployeeswithCode8000(map1 map[int]string, startswith int) map[int]string {
	fmt.Println("This is Defered function")
	for key, value := range map1 {
		if startswith == key {
			fmt.Printf("Deleting Value %s \n", value)
			delete(map1, key)

		}
	}
	return map1
}
func removeEmployees(map1 map[int]string, startswith int) map[int]string {
	fmt.Println("Deleting employee with ", startswith)
	defer removeEmployeeswithCode8000(map1, 8000)
	for key, value := range map1 {
		if startswith == key {
			fmt.Printf("Deleting Value %s \n", value)
			delete(map1, key)

		}
	}
	return map1
}
func main() {
	namesmap := make(map[int]string)
	namesmap[1001] = "Alice"
	namesmap[1002] = "Bob"
	namesmap[1003] = "Cat"
	namesmap[1004] = "David"
	namesmap[2000] = "Eve"
	namesmap[2001] = "Fred"
	namesmap[2005] = "Jack"
	namesmap[8000] = "John"

	namesmap = removeEmployees(namesmap, 1001)
	fmt.Println("Defer also called")
	fmt.Println(namesmap)
}
