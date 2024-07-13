# Let's GO

In questo progetto si trova una raccolta di esempi di base per essere introdotti al linguaggio GO

in particolare la divisione per package è per argomenti e gli argomenti trattati sono

- array
- cicli
- directories
- errorHandling
- files
- interfaces
- mappe
- puntatori
- strutture

L'idea e' quella di raccogliere definizioni e esempi pratici.


### Moduli Definizione

- Un modulo è una collezione di pacchetti che vengono distribuiti insieme, permettendo una migliore gestione delle versioni del software.
- Un modulo è definito da un file `go.mod` nella directory radice del modulo.
- Il file `go.mod` elenca la versione di Go e le dipendenze del modulo.

#### Come creare un nuovo modulo

Per creare un nuovo modulo, esegui il seguente comando:

```sh
go mod init miomodulo
```
Questo comando creerà un nuovo file go.mod nella directory corrente con il seguente contenuto:

```text
module miomodulo
go 1.20
```
Il file go.mod identifica la directory come radice del modulo.

#### Scaricare dipendenze
Per scaricare una dipendenza, utilizza il comando:

```sh
go get github.com/gin-gonic/gin
```

Per rimuovere dipendenze non utilizzate e aggiornare il file go.mod, esegui:

```sh
go mod tidy

```

## File go.mod e go.sum Cosa sono??

- go.sum : contiene le somme di controllo crittografiche delle versioni esatte di ogni modulo di cui il nostro progetto dipende, garantisce che le  future installazioni siano esattamente le stesse cosi da non avere modifiche inattese ee percio' non e' consigliato modificare a mano questo file   
- go.mod : ci permette di gestire le versioni delle nostre dipendemnze.. 

`go get` aggiorna anche il file go.sum
`go mod tidy` rimuove tutte le dipendenze non necessarie dal file go.mod


## Occhio agli errori nelle importazioni uno dei piu comuni 

- Importazione Ciclica, se uno o piu' pacchetti si importano a vicenda
- Altro

```sh
go build
```
Il comando go build compila i pacchetti Go specificati nei file sorgente e crea un eseguibile.

Esempio di utilizzo:

```sh
go build main.go
```
Questo comando compila il file sorgente main.go e genera un file eseguibile chiamato main (o main.exe su Windows) nella stessa directory. Puoi eseguire l'applicazione digitando ./main nel terminale.

Nota: go build non produce alcun output se la compilazione ha successo. Se si verificano errori di compilazione, verranno stampati nel terminale.
```bash
go test
```
Il comando go test esegue test sul codice scritto in Go. Go ha un pacchetto di testing nativo chiamato testing, che fornisce funzionalità per scrivere test unitari e benchmark.

Esempio di test unitario per una funzione Add che somma due numeri:

```go

package main

import "testing"

func Add(a int, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := Add(1, 2)
    if result != 3 {
        t.Errorf("Add(1, 2) = %d; want 3", result)
    }
}
```
Per eseguire questo test, utilizza il comando:

```sh
go test
```
Questo comando eseguirà tutti i test nel pacchetto corrente. Se il test ha successo, vedrai un output simile a:

```text
Copia codice
PASS
ok  	myapp	0.002s
```
Se il test fallisce, go test mostrerà un messaggio di errore:

```text
--- FAIL: TestAdd (0.00s)
main_test.go:10: Add(1, 2) = 4; want 3
FAIL
exit status 1
FAIL	myapp	0.002s
```

È possibile eseguire test specifici passando il loro nome al comando go test con l'opzione -run. Ad esempio, go test -run TestAdd eseguirà solo il test TestAdd.

Package Logger
Questo pacchetto fornisce un'implementazione di un logger singleton in Go.

```go
package logger

import (
    "log"
    "os"
    "sync"
)

// Logger è una struttura che incapsula un'istanza di log.Logger.
type Logger struct {
    *log.Logger
}

// Variabili di package-level utilizzate per gestire l'istanza singleton.
var (
    loggerInstance *Logger
    loggerOnce     sync.Once
)

// GetLoggerInstance restituisce l'istanza singleton del logger.
// Se l'istanza non esiste, la crea inizialmente in modo thread-safe usando sync.Once.
func GetLoggerInstance() *Logger {
    loggerOnce.Do(func() {
        loggerInstance = &Logger{
            log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile),
        }
    })
    return loggerInstance
}
```

```go

package main

import (
    "github.com/tuo-username/logger"
)

func main() {
    // Ottenere l'istanza singleton del logger
    log := logger.GetLoggerInstance()
    // Utilizzare il logger per scrivere un messaggio di log
    log.Println("Questo è un messaggio di log")
}
```

In questo esempio, otteniamo e utilizziamo l'istanza singleton del logger per scrivere messaggi di log. Usando questo pattern, si assicura che ci sia una sola istanza del logger in tutta l'applicazione, evitando problemi di concorrenza e duplicazione.

## CONCORRENZA - PARALLELISMO, pecurialità di GO 


La concorrenza è un termine utilizzato in informatica per descrivere l'esecuzione di più task che possono essere gestiti in ordine indipendente. Si riferisce all'abilità di un programma di gestire più task in un ordine indipendente.

Immagina di dover eseguire diverse operazioni che richiedono molto tempo, come l'accesso a un database o il download di file da internet. Se eseguissi queste operazioni una alla volta, il programma sarebbe lento e inefficiente. Con la concorrenza, possiamo avviare più operazioni simultaneamente. Nota che questo non implica necessariamente che i task vengano eseguiti nello stesso momento, ma rende il programma più veloce ed efficiente.

A volte la concorrenza viene erroneamente scambiata con il parallelismo. Ad esempio, un singolo thread potrebbe gestire la concorrenza attraverso tecniche come il time slicing, ovvero eseguendo una parte di un task, interrompendosi per eseguire una parte di un altro task, e così via.

Il parallelismo, invece, riguarda l'esecuzione simultanea di più task.

## MODELLO DI CONCORRENZA DI GO: CSP

Tony Hoare nel 1978 : Definisce il CSP 
**Communicating Sequential Process (CSP)** è un modello che permette a due o più processi indipendenti di comunicare tra loro attraverso dei canali.

- In termini piu semplici il CSP è un modello che permette a due o piu' processi indipendenti di comunicare tra loro attraverso dei canali

*** Non comunicare facendo condividere la memoria; condividi la memoria comunicando" ***

I Processi indipendenti sono chiamati goroutines e i canali sono utilizzati per comunicare tra queste goroutines
Questa comuinicazione è sicura per la concorrenza, il che significa che 2 go-routines possono accedere in modo sicuro alla stessa variabile di memoria attraverso un canale
Significa che due go routines possono accedre in modok sicuro alla stessa variabile
Una carattersitica chiave del modello CSP in GO è che i canali non solo trasportano i dati, ma anche il cotnrollo
quindi quando una GO routine invia un dato ad un'altra goRoutine atrtaverso un canale
La goRoutine riceve sia i dati sia il controllo
Quindi questa sincronizzazione fa parte del design di GO e permette anche di evitare molte altre difficoltà

### Cosa sono le GO-routines?
Una goroutine è un thread leggere gestito dal rutine di go
ricorda il thread è la piu piccola unità di elaborazione che puo' essere gestita dal sistema operativo
Le goRoutines sono simili ai thread ma richiedono meno risorse e sono gestite dal runtime di GO invce che direttamente dal sistema operativo

come avviamo una go-routine?
usiamo la parola chiave go

- go myFunction()

### Esempio di Concorrenza in Go


## Per i javisti qualche paragone pratico 

** Come creo una HashMap<String,Object> map = new HashMap<>(); **

// Creazione di una mappa con chiavi di tipo stringa e valori di tipo interface{}
var myMap map[string]interface{} = make(map[string]interface{}



#### Author 
- Andrea Cavallo