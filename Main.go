package main

import (
	"Example/concorrenza"
	_ "Example/errorHandling"
	"Example/logger"
	_ "Example/mappe"
)

var loggerInstance = logger.GetLoggerInstance()

func main() {

	loggerInstance.Println("******************************************************")
	//puntatori.EsempiPuntatoriStruct()
	//	directories.EsempioConDirectories()
	go concorrenza.Say("Saluta con go-routine")
	concorrenza.Say("Saluta senza go-routine")

}
