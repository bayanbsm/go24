package main

import "fmt"


type Employee struct {
    name string
    id  int
}

type Manager struct {
    Employee  
    Department string
}


func (e Employee) work() {
    fmt.Printf("%s %d\n", e.name, e.id)
}
func main() {
    m := Manager{
		Employee: Employee{ 
			name: "Mimba", 
			id: 1512,
		}, 
		Department: "IT",
	}
    m.work()  
}


