// package main

// import (

// 	"log"
// 	"runtime"
// )

// func main() {
// 	// numCPUs := runtime.NumCPU()
// 	// log.Printf("number of cpu available: %d\n",numCPUs)

// }

// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	switch runtime.GOOS{
// 	case "linux":
// 		fmt.Println("running on linux")
// 	case "darwin":
// 		fmt.Print("runing on macos")
// 	case "windows":
// 		fmt.Println("running on windows")

// 	case "unknown":
// 		fmt.Printf("uknown")
// 	}

// }

package main

import "fmt"

// func simplefunction(){
// 	fmt.Println("Simple function");
// }

// func add(a,b int) int {
// 	return a+b
// }

func divide(a, b float64) (float64, error){
	if b==0 {
		return 0, fmt.Errorf("denominator must be zero")
	}
	return a/b, nil
}


func main(){
	// fmt.Println("learn in good way");

	// var name string = "89";
	// var version = "version latest";
	// fmt.Println(name);
	// fmt.Println(version)

	// var money int = 67000
	// var currency = 3489
	// fmt.Println(money)
	// fmt.Println("currecny",currency);
	// var dimension float64 = 87.12
	// fmt.Println(dimension)

	// var decided bool = false 
	//  fmt.Println(decided)

	//  var person = 23
	//  fmt.Println(person)

	//  const pi= 67.63
	//  fmt.Println(pi)

	// person := 123
	// fmt.Println(person)

	// name := "dude"
	// fmt.Println(name)

	// fmt.Println("what is your name");
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("hello, mr.",name);

	// reader := bufio.NewReader(os.Stdin)
	// name,_ := reader.ReadString('\n')
	// fmt.Println("hello mr.",name)

	// fmt.Println("function calling start: -")
	// simplefunction()
	
	// ans := add(3,4)
	// fmt.Println("add of two number is:", ans)

fmt.Println("error handling in function")
ans, _ := divide(10,0) // for ignore error
fmt.Println(ans)


// fmt.Println("error handling in function")
// ans, err := divide(10,0)
// if err != nil{
// 	fmt.Println("Error handling")
// }
// fmt.Println("Division of two number is",ans)


}