package strutture

import "Example/logger"

var loggerInstance = logger.GetLoggerInstance()

// DEFINIZIONE DI STRUCT:
// collezione di dati di vari tipi raggruppati insieeme
// definisci struttura cerchio
type Cerchio struct {
	Raggio float64
}

// definiamo due metodi per la struttura Cerchio
// NOTA: ( c Cerchio ) è il Ricevitore del metodo non un Puntatore!
// come definisci un metodo con un ricevitore di valore in GO
// ( v tipo ) nomeMetodo() {}
func (c Cerchio) Area() float64 {
	return c.Raggio * 3.14 * c.Raggio
}

func (c Cerchio) Circonferenza() float64 {
	return c.Raggio * 3.14 * 2
}

func EsempiConCerchio() {

	c := Cerchio{Raggio: 10}
	loggerInstance.Println("L'area del cerchio: ", c.Area())
	loggerInstance.Println("Circonferenza del Cerchio: ", c.Area())

}
