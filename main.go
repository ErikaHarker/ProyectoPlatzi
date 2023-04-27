package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var countID int = 1

type Employee struct {
	Age            int
	Name, LastName string
	Occupation     string
}

func (u Employee) FullName() string {
	return u.Name + " " + u.LastName
}

func (u Employee) String() string {
	return fmt.Sprintf("%s , ocupacion: %s", u.FullName, u.Occupation)
}

func cargaDatosEjemplo() map[int]Employee {
	listEmployee := make(map[int]Employee)
	employee := Employee{150, "Frailejon", "Perez", "ambientalista"}
	employee2 := Employee{50, "Frank", "Molina", "destective"}
	employee3 := Employee{53, "Frederick", "Starks", "psicoanalista"}
	listEmployee[1] = employee
	listEmployee[2] = employee2
	listEmployee[3] = employee3

	countID = 4

	return listEmployee
}

func AddEmployee(name string, lastName string, occupation string, age int, listEmployee map[int]Employee) map[int]Employee {
	employee := Employee{150, name, lastName, occupation}
	listEmployee[countID] = employee

	countID++

	return listEmployee
}

func (e *Employee) SetEmployee(name string, lastName string, occupation string, age int) {
	e.Name = name
	e.LastName = lastName
	e.Occupation = occupation
	e.Age = age
}

func printAllEmployee(listEmployee map[int]Employee) {
	for key, value := range listEmployee {
		fmt.Println("PERSONA ID: ", key)
		fmt.Println("Nombre completo: ", value.FullName())
		fmt.Println("Edad: ", value.Occupation)
	}
}

func NameSlow(name string, channel chan string) {
	nameChars := strings.Split(name, "")
	for _, nChar := range nameChars {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(nChar)
	}
	fmt.Print("\n")
	channel <- ".. FIN .."
}

func main() {
	var option string
	var op int = -1
	var err error
	var listEmployee map[int]Employee
	listEmployee = cargaDatosEjemplo()
	for op != 0 {
		fmt.Print("Menu de empleados:\n")
		fmt.Print("1.) Consultar lista de empleados\n")
		fmt.Print("2.) Agregar lista de empleados\n")
		fmt.Print("3.) Editar empleado\n")
		fmt.Print("4.) Eliminar lista de empleados\n")
		fmt.Print("5.) Mostrar nombre de empleado lento\n")
		fmt.Print("\n0.) Salir\n")

		fmt.Scanln(&option)
		op, err = strconv.Atoi(option)
		if err == nil {
			switch op {
			case 0:
				break
			case 1:
				printAllEmployee(listEmployee)
			case 2:
				var name, lastName, occupation, ageStr string
				var age int
				fmt.Println("Escriba los siguientes datos del nuevo empleado: ")
				fmt.Print("Nombre: ")
				fmt.Scanln(&name)
				fmt.Print("Apellido: ")
				fmt.Scanln(&lastName)
				fmt.Print("Edad: ")
				fmt.Scanln(&ageStr)
				age, err = strconv.Atoi(ageStr)
				for err != nil {
					fmt.Print("Error!!, ingrese solo numeros, Edad: ")
					fmt.Scanln(&ageStr)
					age, err = strconv.Atoi(ageStr)
				}
				fmt.Print("Ocupacion: ")
				fmt.Scanln(&occupation)
				listEmployee = AddEmployee(name, lastName, occupation, age, listEmployee)
				fmt.Println("... Se agrego correctamente el usuario ...")
			case 3:
				var name, lastName, occupation, ageStr, idEmployeeStr string
				var age, idEmployee int
				fmt.Print("Escriba el ID del empleado a Modificar: ")
				fmt.Scanln(&idEmployeeStr)
				idEmployee, err = strconv.Atoi(idEmployeeStr)
				for err != nil {
					fmt.Print("Error!!, ingrese solo numeros, ID: ")
					fmt.Scanln(&idEmployeeStr)
					idEmployee, err = strconv.Atoi(idEmployeeStr)
				}
				employee, ok := listEmployee[idEmployee]
				if ok {
					fmt.Println("Escriba los siguientes datos del empleado: ")
					fmt.Print("Nombre: ")
					fmt.Scanln(&name)
					fmt.Print("Apellido: ")
					fmt.Scanln(&lastName)
					fmt.Print("Edad: ")
					fmt.Scanln(&ageStr)
					age, err = strconv.Atoi(ageStr)
					for err != nil {
						fmt.Print("Error!!, ingrese solo numeros, Edad: ")
						fmt.Scanln(&ageStr)
						age, err = strconv.Atoi(ageStr)
					}
					fmt.Print("Ocupacion: ")
					fmt.Scanln(&occupation)
					employee.SetEmployee(name, lastName, occupation, age)
					listEmployee[idEmployee] = employee

				} else {
					fmt.Println("El empleado ", idEmployee, " no existe")
				}
			case 4:
				var idEmployeeStr string
				var idEmployee int
				fmt.Print("Escriba el ID del empleado a eliminar: ")
				fmt.Scanln(&idEmployeeStr)
				idEmployee, err = strconv.Atoi(idEmployeeStr)
				for err != nil {
					fmt.Print("Error!!, ingrese solo numeros, ID: ")
					fmt.Scanln(&idEmployeeStr)
					idEmployee, err = strconv.Atoi(idEmployeeStr)
				}
				employee, ok := listEmployee[idEmployee]
				if ok {
					delete(listEmployee, idEmployee)
					fmt.Println(employee, " Fue eliminado")
				} else {
					fmt.Println("El empleado ", idEmployee, " no existe")
				}
			case 5:
				var idEmployeeStr string
				var idEmployee int
				fmt.Print("Escriba el ID del empleado: ")
				fmt.Scanln(&idEmployeeStr)
				idEmployee, err = strconv.Atoi(idEmployeeStr)
				for err != nil {
					fmt.Print("Error!!, ingrese solo numeros, ID: ")
					fmt.Scanln(&idEmployeeStr)
					idEmployee, err = strconv.Atoi(idEmployeeStr)
				}
				employee, ok := listEmployee[idEmployee]
				if ok {
					channel := make(chan string)
					go NameSlow(employee.Name, channel)
					msg := <-channel
					fmt.Println(msg)
				} else {
					fmt.Println("El empleado ", idEmployee, " no existe")
				}
			default:
				fmt.Print("La opcion no es valida\n")
				continue
			}
		} else {
			fmt.Println("La opcion no es valida")
			op = -1
		}
	}

}
