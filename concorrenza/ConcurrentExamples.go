package concorrenza

import (
	"fmt"
	"sync"
	"time"
)

func Say(s string) {
	// nella funzione main la invoco 2 volte la prima volta con GO
	// la esegue in parallelo rispetto all 'esecuzione del codice cosi vedremo che vengono stampati alternativaemnte
	// non ci sara' un ordine nelle stampe la chiamata con go non blocca il flusso
	// SE IL MAIN TERMINA TUTTE LE GO ROUTINES vengono chiuse
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

// come syncronizzo le go-routines?
// come gestisco eveentuali go-routines orfane?

// esempi con package Sync

// def della funzione worked definita da ogni go-routine
// id che identifica la go-routine
// puntatore a un Sync gruppo
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mi assicuro che viene chiamato quando la fx finisce
	// indipendentemente da come finisce
	fmt.Println("Worker ", id, " starting")

	time.Sleep(time.Second) // simulo lavoro come threa.sleep
	fmt.Println("Worker ", id, " ending")
}

// creo un gruppo syncWaitGroup
// ciclo for per 5 go-routine segnalo al gruppo con l'add che entra una go-routine
// e poi chiamo il worker con id corrente e il puntatore
// alla fine blocco cosi mi assicuro che il programma principale non termini prima
// avvio 5 routine simulo lavoro che dura 1 sec e mi assicuro con il sync.WaitGroup che non termina
func FiveWorkers() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
}

// MUTUAL EXCLUSION AVOID RACE CONDITIONS
// GARANTISCE CHE 1 SOLA GO ROUTINE PUO ACCEDERE ALLA MAPPA
// Creo un tipo SafeCounter che utilizza un MUTEX per proteggere l'accesso
// simultaneo alla mappa v
type SafeCounter struct {
	v     map[string]int
	mutex sync.Mutex
}

// incrementa il valore alla mappa utilizzata
func (c *SafeCounter) Inc(key string) {
	c.mutex.Lock()
	c.v[key]++
	c.mutex.Unlock()
}

// restituisce la value
func (c *SafeCounter) Value(key string) int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.v[key]
}

// NEL MAIN LANCIO 10000 GO ROUTINES
func ExampleWithMutex() {
	c := &SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc("somekey")
			wg.Done()
		}()
	}
	fmt.Println(c.Value("somekey"))
}
