package cicli

import "fmt"

// Esempio di ciclo for classico con if e break
func ForLoopExample() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Interrompe il ciclo quando i è uguale a 5
		}
		fmt.Println(i)
	}
}

// Esempio di ciclo for range con if e break
func ForRangeLoopExample() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range numbers {
		if value == 5 {
			break // Interrompe il ciclo quando il valore è uguale a 5
		}
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}

// Esempio di ciclo while-like con if e break
func WhileLikeLoopExample() {
	i := 0
	for {
		if i == 5 {
			break // Interrompe il ciclo quando i è uguale a 5
		}
		fmt.Println(i)
		i++
	}
}

// le etichette vanno pero' evitate il piu possibile che rendono il codice complesso
// Esempio di utilizzo di etichette con break per interrompere due cicli annidati
func BreakWithLabelExample() {
outerLoop: // etichetta per fare il break di tutto
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i: %d, j: %d\n", i, j)
			if i == 1 && j == 1 {
				break outerLoop // Interrompe entrambi i cicli quando i è 1 e j è 1
			}
		}
	}
}
