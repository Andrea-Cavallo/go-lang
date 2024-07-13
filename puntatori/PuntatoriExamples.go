package puntatori

import "fmt"

// Esempio base di dichiarazione e utilizzo di puntatori
/**
Spiegazione dei Puntatori
Un puntatore è una variabile che conserva l'indirizzo di memoria di un'altra variabile. I puntatori sono utili per modificare i dati alla fonte piuttosto che lavorare su una copia. In Go, i puntatori sono dichiarati usando l'asterisco (*) e l'operatore di dereferenziazione (*) è utilizzato per accedere o modificare il valore a cui il puntatore punta.
*/
// Puntatori dichiarazione
func DichiarazionePuntatore() {
	var p *int                                     // Dichiarazione di un puntatore a un intero, attualmente non punta a nulla
	fmt.Println("Il valore del puntatore p è:", p) // Stampa il valore del puntatore (nil)
}

// Assegnazione di un indirizzo a un puntatore
func AssegnazioneIndirizzo() {
	var x int = 5
	var p *int
	p = &x                                         // Assegna l'indirizzo di x al puntatore p
	fmt.Println("Il valore di x è:", x)            // Stampa il valore di x
	fmt.Println("Il valore del puntatore p è:", p) // Stampa l'indirizzo di x
}

// Utilizzo di un puntatore per modificare il valore di una variabile
func ModificaValore() {
	var x int = 10
	var p *int = &x                               // Assegna l'indirizzo di x al puntatore p
	fmt.Println("Il valore originale di x è:", x) // Stampa il valore originale di x
	*p = 20                                       // Modifica il valore di x attraverso il puntatore p
	fmt.Println("Il nuovo valore di x è:", x)     // Stampa il nuovo valore di x
}

func esempio() {
	var x int = 5
	p := &x                                        // Assegna l'indirizzo di x al puntatore p
	fmt.Println("Il valore di x è:", x)            // Stampa il valore di x
	fmt.Println("Il valore del puntatore p è:", p) // Stampa l'indirizzo di x
	fmt.Println("Il valore a cui punta p è:", *p)  // Stampa il valore a cui punta p, cioè il valore di x
}

func cambioValore(ptr *int) {
	*ptr = 100

}
