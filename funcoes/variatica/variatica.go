package main

import (
	"fmt"
)

func media(numeros ...float64) float64 {
	media := 0.0

	if len(numeros) > 0 {
		total := 0.0
		for _, num := range numeros {
			total += num
		}
		media = total / float64(len(numeros))
	}

	return media
}

func main() {
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}
