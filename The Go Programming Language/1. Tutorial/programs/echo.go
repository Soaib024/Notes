// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

// Echo 1
/*
func main(){
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i];
		sep = " "
	}
	fmt.Println(s)
}

*/

// Echo 2
/*
func main(){
	s, sep := "", " "
	for _, arg := range os.Args[1:]{
		s += sep + arg
	}

	fmt.Println(s)
}
*/

func main(){
	fmt.Println(strings.Join(os.Args[1:], " "))
}

