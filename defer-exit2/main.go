package main

import (
	"fmt"
	"math"
	"sync"
)

type Circle struct {
	Diameter      float64
	Area          float64
	Circumference float64
}

func calculateCircleProperties(diameter float64, wg *sync.WaitGroup, resultCh chan<- Circle) {
	defer wg.Done()

	radius := diameter / 2
	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius

	result := Circle{
		Diameter:      diameter,
		Area:          area,
		Circumference: circumference,
	}

	resultCh <- result
}

func main() {
	// Diameter-diameter lingkaran yang akan dihitung
	diameters := []float64{5.0, 10.0, 15.0, 20.0}

	// Channel untuk hasil perhitungan
	resultCh := make(chan Circle)

	// WaitGroup untuk menunggu semua goroutine selesai
	var wg sync.WaitGroup

	// Memulai goroutine untuk menghitung setiap lingkaran
	for _, diameter := range diameters {
		wg.Add(1)
		go calculateCircleProperties(diameter, &wg, resultCh)
	}

	// Menutup channel resultCh setelah semua perhitungan selesai
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Menerima dan mencetak hasil perhitungan
	for circle := range resultCh {
		fmt.Printf("Diameter: %.2f, Luas: %.2f, Keliling: %.2f\n", circle.Diameter, circle.Area, circle.Circumference)
	}
}
