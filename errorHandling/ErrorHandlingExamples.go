package errorHandling

import (
	"errors"
	"fmt"
)

// Gli errori sono valori che possono essere restituiti dalle funzioni e successivamente controllati

// Funzione che restituisce un errore se il numero è negativo
func CheckPositiveNumber(num int) (int, error) {
	if num < 0 {
		return 0, errors.New("numero negativo: non è permesso")
	}
	return num, nil
}

// Funzione che restituisce un errore personalizzato
type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Errore %d: %s", e.Code, e.Message)
}

func CheckEvenNumber(num int) (bool, error) {
	if num%2 != 0 {
		return false, &CustomError{Code: 1001, Message: "numero dispari: non è permesso"}
	}
	return true, nil
}

// Funzione che gestisce gli errori restituiti da altre funzioni
func HandleErrors() {
	if _, err := CheckPositiveNumber(-5); err != nil {
		fmt.Println("Errore:", err)
	}

	if _, err := CheckEvenNumber(3); err != nil {
		fmt.Println("Errore:", err)
	}
}
