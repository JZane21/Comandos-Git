package main

import (
	"fmt"
	"math"
)

// Compilacion completa:
// go build archivo.go
// ./archivo

// Compilacion rapida:
// go run archivo.go

func main() {

	fmt.Println("Hola mundooo, Golang esta instalado")

	// sieve(10)
	// collatz(10)
	// slice := []int{5, 2, 3, 1, 4}
	// fmt.Println(ordenamiento(slice))
	// fmt.Println(factorial(5))
	//fmt.Println(sumPrimes(50000))
	count := 100000
	fmt.Println(sumPrimes2(count)) //100000
}

func factorial(n int) int {
	fmt.Println("\nFactorial:")
	fmt.Println(n)
	aux := 1
	for i := 2; i <= n; i++ {
		aux *= i
	}
	return aux
}

func ordenamiento(n []int) []int {
	fmt.Println("\nBubble Sort:")
	fmt.Println(n)
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n); j++ {
			if n[i] > n[j] {
				aux := n[i]
				n[i] = n[j]
				n[j] = aux
			}
		}
	}
	return n
}

func collatz(n int) {
	fmt.Println("\nCollatz:")
	fmt.Println(n)
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		fmt.Println(n)
	}
}

func sieve(n int) {
	fmt.Println("\nSieve:")
	fmt.Println("Numeros primos hasta: ", n)
	slice := []int{}
	for i := 0; i <= n; i++ {
		slice = append(slice, i)
	}

	slice[1] = 0
	size := len(slice)

	for i := 2; i < size; i++ {
		if slice[i] != 0 {
			fmt.Println(slice[i])
			for j := i * i; j < size; j += i {
				slice[j] = 0
			}
		}
	}

}

func sumPrimes(numero int) int {
	fmt.Println("\nSuma de numeros Primos Version 1:")

	n := numero
	sum := 0
	contador := 0
	lastValue := 0
	lastNP := 0
	for contador < n {
		slice := []int{}
		for i := 0; i <= numero; i++ {
			slice = append(slice, i)
		}

		slice[1] = 0
		size := len(slice)
		lastNP = slice[len(slice)-1]
		for i := 2; i < size; i++ {
			if slice[i] != 0 {
				sum += slice[i]
				contador++
				if contador == n {
					lastValue = slice[i]
					i = size
				}
				for j := i * i; j < size; j += i {
					slice[j] = 0
				}
			}
		}
		if contador < n {
			contador = 0
			sum = 0
			numero = numero + 100
		}
	}

	sieve(lastValue)
	fmt.Println("Ultimo valor del Slice: ", lastNP)
	fmt.Print("\nLa suma es: ")
	return sum

}

func sumPrimes2(numero int) int {
	fmt.Println("\nSuma de numeros Primos Version 2")
	n := numero
	numero = numero*(int(math.Sqrt(float64(numero)))) + 2*int(math.Sqrt(float64(numero)))
	sum := 0
	contador := 0
	lastValue := 0
	lastNP := 0
	slice := []int{}
	for i := 0; i <= numero; i++ {
		slice = append(slice, i)
	}
	slice[1] = 0
	size := len(slice)
	lastNP = slice[len(slice)-1]
	for i := 2; i < size; i++ {
		if slice[i] != 0 {
			sum += slice[i]
			contador++
			if contador == n {
				lastValue = slice[i]
				i = size
			}
			for j := i * i; j < size; j += i {
				slice[j] = 0
			}
		}
	}
	if false {
		sieve(lastValue)
	}
	fmt.Println("Ultimo valor del Slice: ", lastNP)
	fmt.Print("\nLa suma es: ")
	return sum

}
