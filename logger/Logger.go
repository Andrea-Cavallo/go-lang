// Package logger fornisce una semplice implementazione di un logger singleton.
// Questo assicura che ci sia una sola istanza del logger in tutta l'applicazione,
// garantendo l'unicità e l'accesso globale al logger.

package logger

import (
	"log"
	"os"
	"sync"
)

// Logger è una struttura che incapsula un'istanza di log.Logger.
// Questo permette di estendere le funzionalità di log.Logger se necessario.
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
