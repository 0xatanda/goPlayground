package main

import "fmt"

type authInfo struct {
	username string
	password string
}

func (aut authInfo) getBasicAuth() string {
	return fmt.Sprintf("Auth: Basic %s:%s", aut.username, aut.password)
}

func test(auth authInfo) {
	fmt.Println(auth.getBasicAuth())

}

func main() {
	test(authInfo{
		username: "0x",
		password: "atanda",
	})
}
