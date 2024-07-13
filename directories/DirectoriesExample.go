package directories

import (
	"Example/logger"
	"io/ioutil"
	"log"
	"os"
)

var loggerInstance = logger.GetLoggerInstance()

//  cosa serve il param 0755?
// rappresentazione in ottale dei permessi di un file o directory che stai creando
//  nei sistemi unix-like si us sono i permessi dei file suddivisi in tre categorie
// 1 - permesso per il proprietario del file
// 2 - permesso per il gruppo del file
// 3 - permesso per gli altri utenti
// 4 - lettura , 2 scrittura, 1 esecuzione
// 0755
// 7 (permesso per il propr del file) che in ottale e' 111 in binario puo leggere scrivere e eseguire 4+2+1
// 5 (permesso per il gruppo del file)- il gruppo legge eseguire ma non puo escrivere 4+0+1
// 5 (permesso per gli altri utenti) - solo leggere ed eseguire

func readFilesFromDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)

	} else {
		// _ non ci interessa nessun indice
		for _, file := range files {
			loggerInstance.Println(file.Name())
		}
	}
}

// creo directory
// creo file
// poi leggo e stampo contenuto
// gestione exceptions

func EsempioConDirectories() {

	loggerInstance.Println("CREO DIRECTORY")
	err := os.Mkdir("testDir", 0755)
	if err != nil {
		log.Fatal(err)
	}
	loggerInstance.Println("CREO FILE DI TXT NELLA DIRECTORY")

	file, err := os.Create("testDir/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	loggerInstance.Println("SCRIVO SUL FILE DI TESTO")

	_, err = file.WriteString("Ciao Andrea")
	if err != nil {
		log.Fatal(err)
	}
	loggerInstance.Println("CHIUDO IL FILE DI TESTO PER EVITARE DANNI")

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	loggerInstance.Println("LEGGO IL FILE DI TESTO DALLA DIRECTORY")

	data, err := (ioutil.ReadFile("testDir/test.txt"))
	if err != nil {
		log.Fatal(err)
	}

	loggerInstance.Println(string(data))

}
