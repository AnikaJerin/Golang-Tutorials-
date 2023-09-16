// practice 1

// package is something that has various files and codes and there are different kinds of packages in golang
//  Here package main will tell the go complier that our code should be complied into an executable program in the end

package main

import "fmt" //Importing packages 

// main function is the entry point of our program. this line will execute our whole program
// which means if we run our program go will look through our files in our programs for this funtion 
// and it will execute this funtion autometically. Our program ca have ONLY ONE main func

func main(){
	fmt.Println("Hello Anika")
}