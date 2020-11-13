package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {

    //creating struct specifying field names
    emp1 := Employee{
        firstName: "Himanshu",
        age:       24,
        salary:    10000,
        lastName:  "Kumar",
    }

    //creating struct without specifying field names
    emp2 := Employee{"John", "Kumar", 29, 11800}

    fmt.Println("Employee 1", emp1)
    fmt.Println("Employee 2", emp2)
}