package errorHandling

import (
	"errors"
	"fmt"
	"regexp"
)

// User struct che contiene i dati dell'utente
type User struct {
	Username string
	Email    string
	Age      int
}

// ValidateUser funzione che valida i dati dell'utente
func ValidateUser(user User) error {
	if user.Username == "" {
		return errors.New("il nome utente non può essere vuoto")
	}

	if len(user.Username) < 3 {
		return errors.New("il nome utente deve avere almeno 3 caratteri")
	}

	if !isValidEmail(user.Email) {
		return errors.New("email non valida")
	}

	if user.Age < 18 {
		return errors.New("l'età deve essere almeno 18 anni")
	}

	return nil
}

// isValidEmail funzione di supporto per validare l'email
func isValidEmail(email string) bool {
	// Usa una regex semplice per validare l'email
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func main2() {
	user := User{
		Username: "john",
		Email:    "john.doe@example.com",
		Age:      20,
	}

	if err := ValidateUser(user); err != nil {
		fmt.Println("Errore di validazione:", err)
	} else {
		fmt.Println("Utente valido")
	}
}
