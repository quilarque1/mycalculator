package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calc struct {}

func (Calc) Operate (entrada string, operador string) int {
	entradaLimpia := strings.Split(entrada, operador)
	valor1 := parsearOperador(entradaLimpia[0])
	valor2 := parsearOperador(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println(valor1 + valor2)
		return valor1 + valor2
	case "-":
		fmt.Println(valor1 - valor2)
		return valor1 - valor2
	case "*":
		fmt.Println(valor1 * valor2)
		return valor1 * valor2
	case "/":
		fmt.Println(valor1 / valor2)
		return valor1 / valor2
	default:
		fmt.Println("Operador Invalido")
		return 0
	}
}

func parsearOperador(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func LeerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese la operacion que desea realizar")
	scanner.Scan()
	return scanner.Text()
}
