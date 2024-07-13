package files

import (
	"Example/logger"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var loggerInstance = logger.GetLoggerInstance()

//GO ha un supporto top quality per la scrittura e lettura dei file
// infatti c'e il pacchetto OS che fornisce funzioni

func creaNuovoFile() {

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	// uso defer per assicurarmi che venga chiuso alla fine della funzione indipendentemente da dove c'e' l uscita della fx
	// per evitare perdite di risorse
	defer file.Close()

}

// deprecato perche lento e fa cagare
func leggiFile() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

// come leggo file di grandi dimensioni?
// li devi leggere a pezzi !!!!
// usi pacchetto bufio in maniera bufferizzata
// elaborpo un pezzo del file e poi il successivo
func leggiFileDiGrandiDim() {

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)

	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		loggerInstance.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// scritture su file come?

func scriviSuFile() {

	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// se file esiste gia' il contenuto viene cancellato e riscritto
	// questo perche aprendo in modalita' scrittura e' cosi
	// se lo apri in modalita' append aggiungi utilizzando os.Append anzichè os.Create
	_, err = file.WriteString("Daje zio")
	if err != nil {
		log.Fatal(err)
	}
	// ignoro numero di Bye scritti  se writeString restituisce un err
}

// Gestione di Errori - approccio Univoco di GO
// a differenza di altri linguaggi Go restituisce un ERRORRE come un secondo valore
// APPROCCIO CHE OBBLIGA a gestirlo immediatamente anzichè rimandarne la gestione
// file non esiste
// non hai permessi di scrittura, lettura - esempio

func esempioWhithErrore() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		loggerInstance.Println(string(data))
	}

}
