package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	dilbert := Employee{Name: "Dilbert"}
	fmt.Println(dilbert)

	position := &dilbert.Name
	*position = "Lucas"
	fmt.Println(*position)

}
