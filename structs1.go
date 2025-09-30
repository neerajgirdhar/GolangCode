package main

import (
	"fmt"
)

type AuthenticationInfo struct {
	username string
	password string
}

func (authInfo AuthenticationInfo) getBasicAuthInfo() string {
	return fmt.Sprintf("Basic Authentication %s:%s", authInfo.username, authInfo.password)
}
func main() {
	fmt.Println()
	fmt.Println("Methods on structs it used when you have to compute some values from the properties of struts this is another syntax and is used in interfaces")
	authInfo := AuthenticationInfo{username: "John", password: "123456"}
	fmt.Println(authInfo.getBasicAuthInfo())
}
