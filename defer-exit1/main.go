package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fungsi untuk menghitung faktorial dari suatu angka
func calculateFactorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * calculateFactorial(n-1)
}

// Fungsi untuk menghasilkan channel faktorial
func Factorial(ch <-chan int) chan int {
	factorialCh := make(chan int)

	go func() {
		defer close(factorialCh)
		for n := range ch {
			result := calculateFactorial(n)
			factorialCh <- result
		}
	}()

	return factorialCh
}

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) + 1 // Menghasilkan angka acak dari 1 hingga 100

	fmt.Printf("Menghitung faktorial dari %d angka acak:\n", n)

	inputCh := make(chan int)

	// Mengirim n angka acak ke dalam channel inputCh
	go func() {
		defer close(inputCh)
		for i := 0; i < n; i++ {
			inputCh <- rand.Intn(10) // Menggunakan angka acak dari 1 hingga 10 sebagai input faktorial
		}
	}()

	factorialCh := Factorial(inputCh)

	// Menerima dan menampilkan hasil faktorial dari channel factorialCh
	for result := range factorialCh {
		fmt.Printf("Hasil faktorial: %d\n", result)
	}
}
