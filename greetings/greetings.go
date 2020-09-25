package greetings
import (
	"errors"
	"fmt"
)

//fungsi hello mengembalikan nilai nama yg ada pada pesan
func Hello(nama string) (string, error){

	if nama == "" {
		return "", errors.New("gak ada namanya")
	}

	//kalo ada namanya masuk ke message
	//di greeting message
	message := fmt.Sprintf("Asalamualaikum, %v. selamat datang", nama)
	return message, nil
}