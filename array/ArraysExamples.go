package array

// Esempio per spiegare gli array in Go
func ArrayExample() [5]int {
	// Inizializza un array di 5 interi
	var arr [5]int
	arr[0] = 1 // Imposta il primo elemento a 1
	arr[1] = 2 // Imposta il secondo elemento a 2
	// Gli altri elementi sono inizializzati automaticamente a 0

	return arr
}

// Esempio per spiegare le slices in Go
func SliceExample() []int {
	// Inizializza una slice con 3 elementi
	slice := []int{1, 2, 3}

	return slice
}

// Esempio per spiegare l'append nelle slices in Go
func AppendValues() []int {
	var numSlice []int                      // Inizializza una slice vuota
	val1, val2 := 3, 4                      // Definisce due valori da aggiungere
	numSlice = append(numSlice, val1, val2) // Usa append per aggiungere i valori alla slice

	return numSlice
}

// Esempio per spiegare la funzione make con le slices
func InitializeSliceWithMake() []int {
	// Inizializza una slice di int con una lunghezza di 5 elementi
	numSlice := make([]int, 5)
	// Gli elementi sono inizializzati automaticamente a 0

	return numSlice
}

// Esempio per dimostrare la differenza tra la lunghezza e la capacità di una slice
func LengthCapacityExample() ([]int, int, int) {
	// Inizializza una slice con 2 elementi e capacità di 5
	slice := make([]int, 2, 5)
	slice[0] = 1 // Imposta il primo elemento a 1
	slice[1] = 2 // Imposta il secondo elemento a 2

	length := len(slice)   // Ottiene la lunghezza della slice
	capacity := cap(slice) // Ottiene la capacità della slice

	return slice, length, capacity
}
