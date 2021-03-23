package Functions

import (
	"fmt"
)

/*
This package is for when you want to pass an argument to function which on the function side is call
 a parameter
*/

func Server() {
	port := 3000

	// Passing a argument to a function
	//StartServer(port, 4)

	// Passing a argument and getting a parameter back.
	port, err := StartServer2(port)
	fmt.Println(port, err)
	// The underscore below is used when you don't want to use all return values.  It is called a write only variable.
	_, err2 := StartServer2(port)
	fmt.Println(err2)
}

func StartServer(port int, numberOfretries int) {
	// This function is example of just passing argument to a function.
	fmt.Println("The server has started on port: ", port)
	fmt.Println("The number of retries", numberOfretries)
}

func StartServer2(port int) (int, error) {
	// This function is example of just passing argument to a function, and returning parameter to caller.  The int declared on the return block
	// is used to capture an error number if any.

	fmt.Println("The server has started on port: ", port)
	//fmt.Println("The number of retries", numberOfretries)
	return port, nil
}
