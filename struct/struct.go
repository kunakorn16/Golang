package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	employeeID:   "101",
	// 	employeeName: "Pond",
	// 	phone:        "0983145470",
	// }
	// employeeList[1] = employee{
	// 	employeeID:   "102",
	// 	employeeName: "Jay",
	// 	phone:        "0652271924",
	// }
	// employeeList[2] = employee{
	// 	employeeID:   "103",
	// 	employeeName: "Guy",
	// 	phone:        "0873840856",
	// }
	// fmt.Println("employee = ", employeeList)

	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "Pond",
		phone:        "0983145470",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "Jay",
		phone:        "0652271924",
	}
	employee3 := employee{
		employeeID:   "103",
		employeeName: "Guy",
		phone:        "0873840856",
	}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)

	fmt.Println("employee = ", employeeList)
}
