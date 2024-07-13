package mappe

import (
	"fmt"
	"reflect"
	"testing"
)

// RICORDA CHE SULLE MAPPE L'ORDINE NON VIENE MANTENUTO

// Esempio per spiegare l'inizializzazione di una mappa
func InitializeMap() map[string]int {
	// Inizializza una mappa con chiavi di tipo stringa e valori di tipo intero
	var myMap map[string]int
	// Le mappe devono essere inizializzate usando make o assegnando una mappa letterale
	myMap = make(map[string]int)

	return myMap
}

// Esempio per aggiungere elementi a una mappa
func AddElementsToMap() map[string]int {
	myMap := make(map[string]int)
	myMap["a"] = 1 // Aggiunge un elemento con chiave "a" e valore 1
	myMap["b"] = 2 // Aggiunge un elemento con chiave "b" e valore 2

	return myMap
}

// Esempio per accedere agli elementi di una mappa
func AccessMapElement() int {
	myMap := map[string]int{"a": 1, "b": 2}
	value := myMap["a"] // Accede al valore associato alla chiave "a"

	return value
}

// Esempio per cancellare un elemento da una mappa
func DeleteMapElement() map[string]int {
	myMap := map[string]int{"a": 1, "b": 2}
	delete(myMap, "a") // Cancella l'elemento con chiave "a"

	return myMap
}

// Esempio per iterare su una mappa
func IterateMap() map[string]int {
	myMap := map[string]int{"a": 1, "b": 2}

	for key, value := range myMap {
		// Itera su ogni coppia chiave/valore nella mappa
		fmt.Printf("Chiave: %s, Valore: %d\n", key, value)
	}

	return myMap
}

// Funzione per creare il catalogo di libri
func CatalogoLibri() map[string]int {
	// Inizializza la mappa con i titoli dei libri come chiavi e il numero di pagine come valori
	catalogo := InitializeMap()

	// Aggiunge alcuni libri al catalogo
	catalogo["Il Signore degli Anelli"] = 1200
	catalogo["Harry Potter e la Pietra Filosofale"] = 223
	catalogo["Il Nome della Rosa"] = 512
	catalogo["1984"] = 328
	catalogo["Il Grande Gatsby"] = 180

	return catalogo
}

func TestInitializeMap(t *testing.T) {
	result := InitializeMap()
	// Verifica che il tipo di dati restituito sia corretto
	if reflect.TypeOf(result).Kind() != reflect.Map {
		t.Errorf("InitializeMap() error: expected %v, got %v", reflect.Map, reflect.TypeOf(result).Kind())
	}
	// Verifica che la mappa restituita sia vuota
	if len(result) != 0 {
		t.Errorf("InitializeMap() error: expected %v, got %v", 0, len(result))
	}
}

func TestAddElementsToMap(t *testing.T) {
	result := AddElementsToMap()
	expected := map[string]int{"a": 1, "b": 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AddElementsToMap() error: expected %v, got %v", expected, result)
	}
}

func TestAccessMapElement(t *testing.T) {
	result := AccessMapElement()
	expected := 1
	if result != expected {
		t.Errorf("AccessMapElement() error: expected %v, got %v", expected, result)
	}
}

func TestDeleteMapElement(t *testing.T) {
	result := DeleteMapElement()
	expected := map[string]int{"b": 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("DeleteMapElement() error: expected %v, got %v", expected, result)
	}
}

func TestCatalogoLibri(t *testing.T) {
	result := CatalogoLibri()
	expected := map[string]int{
		"Il Signore degli Anelli":             1200,
		"Harry Potter e la Pietra Filosofale": 223,
		"Il Nome della Rosa":                  512,
		"1984":                                328,
		"Il Grande Gatsby":                    180,
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CatalogoLibri() error: expected %v, got %v", expected, result)
	}
}
