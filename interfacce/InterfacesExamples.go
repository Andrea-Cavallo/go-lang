package interfacce

import "fmt"

//in Go un interfaccia è una collezione di Metodi
// quando crei una struttura crei un nuovo tipo
// Le interfacce sono importanti per specificare un comportamento atteso di un oggetto
// I metodi def in una interfaccia possono essere implementati da qualsiasi tipo

type Parlante interface {
	Parla() string
	Urla() string
}

// posso fare in modo che la nostra Struttura persona implemetni questa interfaccia
// tipo che definisce un insieme di metodi
type Persona struct {
	Nome string
}
type Animale struct {
	Tipo string
}

func (p Persona) Parla() string {
	return "Ciao, sono " + p.Nome + "!"
}

func (p Persona) Urla() string {
	return "CIAOOOOOOOOOOOOOOOOOOOOO " + p.Nome + "!"
}

func (a Animale) Urla() string {
	return "ROARR " + a.Tipo + "!"
}

func saluta(p Parlante) {
	fmt.Println(p.Parla())
}

// Interfaccia vuota
// rapp qualsiasi tipo perche' tutti i tipi implementano zero metodi
// servono le type ASSERTION per capire se
func stampaLunghezza(v interface{}) {

	// se v è effettivamente una stringa allora il type assertion è ok
	// e allora restituisce due valori la stringa s e true

	if s, ok := v.(string); ok {
		fmt.Println(len(s))
	} else {
		fmt.Println("Non è una Stringa..")
	}

}
