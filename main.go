// "Package main" is the namespace declaration
// "main" is a keyword that tells GO that this project is intended to run as a binary/executable (as opposed to a Library)
package main

// importing standard libraries & third party library
import (
	"fmt"
	"os"
	"strings"

	// aliasing library names
	flag "github.com/ogier/pflag"
)

// flags
var (
	user  string
)

// "main" is the entry point of our CLI app
func main() {
	// parse flags
	flag.Parse()

	// if user does not supply flags, print usage
	if flag.NFlag() == 0 {
		printUsage()
	}
	
	// if multiple users are passed separated by commas, store them in a "users" array
	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	
	// "for... range" loop in GO allows us to iterate over each element of the array.
	// "range" keyword can return the index of the element (e.g. 0, 1, 2, 3 ...etc)
	// and it can return the actual value of the element.
	// Since GO does not allow unused variables, we use the "_" character to tell GO we don't care about the index, but
	// we want to get the actual user we're looping over to pass to the function.
	for _, u := range users {
		result := getUsers(u)
		fmt.Println(`Username:	`, result.Login)
		fmt.Println(`Name:		`, result.Name)
		fmt.Println(`Email:		`, result.Email)
		fmt.Println(`Bio:		`, result.Bio)
		fmt.Println("")
	}

}

// "init" is a special function. GO will execute the init() function before the main.
func init() {
	// We pass the user variable we declared at the package level (above).
	// The "&" character means we are passing the variable "by reference" (as opposed to "by value"),
	// meaning: we don't want to pass a copy of the user variable. We want to pass the original variable.
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

// printUsage is a custom function we created to print usage for our CLI app
func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}