package main

import (
	"errors"
	"fmt"
)

func returnErr(a, b int) error {
	if a == b {
		err := errors.New("Error in returnError() func")
		return err
	} else {
		return nil
	}

}

func main() {
	err := returnErr(1, 2)
	if err == nil {
		fmt.Println("returnErr() ended normally!!")
	} else {
		fmt.Println(err)
	}

	err = returnErr(10, 10)
	if err == nil {
		fmt.Println("returnErr() ended normally")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error in returnError() func!!!!" {
		fmt.Println("!!")
	}
}
