package main

import "fmt"

func main() {

	fmt.Println("========= CALCULADORA CIENTIFICA v1.0 ========")

	var historial string
	var total int = 0
	var continuar string

	for {
		var operacion string
		var num1 int
		fmt.Println("Ingrese el primer numero: ")
		fmt.Scanln(&num1)
		var num2 int
		fmt.Println("Ingrese el segundo numero")
		fmt.Scanln(&num2)
		fmt.Println("ingrese la operacion (+, -, *, / , ^ , !): ")
		fmt.Scanln(&operacion)

		var resultados string

		switch operacion {
		case "+":
			resultados = fmt.Sprintf("%d + %d = %d", num1, num2, num1+num2)
		case "-":
			resultados = fmt.Sprintf("%d - %d = %d", num1, num2, num1-num2)
		case "*":
			resultados = fmt.Sprintf("%d * %d = %d", num1, num2, num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Printf("error: division por cero no es permitida\n")
				return
			}
			resultados = fmt.Sprintf("%d / %d = %.2f", num1, num2, float64(num1)/float64(num2))
		case "^":
			resultado := 1
			for i := 0; i < num2; i++ {
				resultado *= num1
			}
			resultados = fmt.Sprintf("%d ^ %d = %d", num1, num2, resultado)
		case "!":
			if num1 < 0 {
				fmt.Printf("error: no existe factorial de numeros negativos\n")
				return
			}
			factorial := 1
			for i := 1; i <= num1; i++ {
				factorial *= i
			}
			resultados = fmt.Sprintf("%d! = %d", num1, factorial)

		default:
			fmt.Println("Operacion no valida")
		}

		total++
		historial += fmt.Sprintf("%d) %s\n", total, resultados)

		fmt.Println("Desea realizar otra operacion? (s/n)")
		fmt.Scanln(&continuar)

		if continuar == "n" {
			break
		}
	}
	fmt.Println("====== Historial de operaciones =====")
	fmt.Print(historial)
	fmt.Printf("Total de operaciones realizadas: %d\n", total)
	fmt.Println("Gracias por usar la calculadora cientifica v1.0")

}
