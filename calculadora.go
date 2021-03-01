package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lang struct {
	language string
}

func (l *lang) startProgram() {
	fmt.Println("Select language Spanish writing [es] or just enter to continue in English..")
	fmt.Println("Selecciona el lenguaje Español escribiendo [es] o sólo da enter para continuar en Inglés..")
	fmt.Print("Lang: ")
	leng := leerEntrada()
	if leng == "es" {
		fmt.Println("Seleccionaste Español")
		l.language = "es"
	} else {
		fmt.Println("Continue with English")
		l.language = "en"
	}
}
func (l *lang) indicationProgram() {
	if l.language == "es" {
		fmt.Println("Nota: Este programa es una calculadora que sólo puede sumar, restar, multiplicar o dividir")
	} else {
		fmt.Println("Note: This program is a calculator that only can add, subtract, multiply or divide")
	}
}

type calc struct{}

func operate(entrada []string, operador string) string {
	result := ""
	resultInt, _ := strconv.Atoi(entrada[0])
	for i := range entrada {
		j := len(entrada)
		if i+1 < j {
			operador2, _ := strconv.Atoi(entrada[i+1])
			switch operador {
			case "*":
				//fmt.Println("Multiplicando..")
				resultInt = resultInt * operador2
				break
			case "/":
				//fmt.Println("Dividiendo..")
				resultInt = resultInt / operador2
				break
			case "+":
				//fmt.Println("Sumando..")
				resultInt = resultInt + operador2
				break
			case "-":
				//fmt.Println("Restando..")
				resultInt = resultInt - operador2
				break
			default:
				fmt.Println(operador, "No está soportado")
				return "error"
			}
		}
	}
	result = strconv.Itoa(resultInt)
	return result
}

func leerEntrada() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func encontrarOperador(entrada string, signo string) int {
	index := strings.Index(entrada, signo)
	return index
}

func doOperations(entrada string, signs []string) string {
	result := "nada"
	fail := ""
	for index := range signs {
		operatorIndex := encontrarOperador(entrada, signs[index])
		if operatorIndex == -1 {
			continue
		}
		aNumbers := strings.Split(entrada, signs[index])
		for i := range aNumbers {
			if len(aNumbers[i]) > 1 {
				aNumbers[i] = doOperations(aNumbers[i], signs)
			}
		}
		result = operate(aNumbers, signs[index])
		if operatorIndex != -1 {
			break
		}
	}

	if len(fail) < 1 {
		return result
	} else {
		return fail
	}
}

func main() {
	signs := []string{"-", "+", "/", "*"}

	/* Bienvenida, selección de leguaje e información del programa*/
	lan := &lang{}
	lan.startProgram()
	lan.indicationProgram()
	fmt.Println("Signos:", signs)

	/*Leer la operación*/
	fmt.Println("Las operaciones son así 3-1+5*5/2 ")
	numero := 0
	for numero < 1 {
		fmt.Print("Ingresa la operación: ")
		entrada := leerEntrada()
		if entrada == "" {
			break
		}
		salida := doOperations(entrada, signs)
		fmt.Println("salida:", salida)
	}
	fmt.Println("Hasta pronto")
}
