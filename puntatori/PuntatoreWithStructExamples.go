package puntatori

import (
	"Example/logger"
)

var loggerInstance = logger.GetLoggerInstance()

type Persona struct {
	Nome    string
	Cognome string
	Eta     int
}

type Impiegato struct {
	Persona
	Ruolo string
}

func compleanno(ptr *Persona) {
	ptr.Eta++
}

// come si passa un puntatore a una funzione in GO? &

func EsempiPuntatoriStruct() {

	p := Persona{Nome: "Andrea", Cognome: "Cavallo", Eta: 35}

	loggerInstance.Println("Persona creata come Struct: ", p)
	//accetta puntatore alla struttura persona
	// all'interno viene incrementata l'eta
	loggerInstance.Println("Prima del compleanno ", p.Nome, " ha : ", p.Eta)
	compleanno(&p)
	loggerInstance.Println("Dopo il compleanno ", p.Nome, " ha : ", p.Eta)
	// copia dell'argomento e non l'argomento originale
	// quando passiamo un puntatore la funzione riceve un riferimento non una copia
	// se modifico il valore a cui punta il puntatore stiamo modificando il valore dell'arg originale
	impiegato := Impiegato{Persona: Persona{Nome: "Andrea", Cognome: "Cavallo", Eta: 200}, Ruolo: "developer"}
	// anche con strutture annidate accedi direttamente
	loggerInstance.Println("Esempio di accesso a una stuct annidata easy :", impiegato.Nome)
}
