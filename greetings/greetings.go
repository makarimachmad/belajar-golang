package greetings
import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//fungsi hello mengembalikan nilai nama yg ada pada pesan
func Hello(nama string) (string, error){

	//kalo gk ada namanya mengembalikan nilai pesan ini
	if nama == "" {
		return nama, errors.New("gak ada namanya")
	}

	//kalo ada namanya masuk ke message
	//di greeting message
	//message := fmt.Sprintf("Asalamualaikum, %v. selamat datang", nama)
	//message := fmt.Sprintf(randomFormat(), nama)

	//untuk ujicoba tes
	message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(namas []string) (map[string]string, error){
	// pemetaan untuk beberapa nama dalam pesan
	messages := make(map[string]string)
	//perulangan menerima pembagian nama untuk mendapatkan pesan
	//dari setiap namanya
	for _, nama := range namas{
		message, err := Hello(nama)
		if err != nil {
			return nil, err
		}
		messages[nama] = message
	}
	return messages, nil
}

//init untuk memberikan nilai variabel yg digunakan
func init(){
	rand.Seed(time.Now().UnixNano())
}

//mengembalikan nilai randomformat greeting messages
//greeting masssage memilih secara random
func randomFormat() string{
	formats := []string{
		"Hi, %v. Selamat datang",
		"Bagus, senang melihat kamu, %v",
		"halo, %v senang bertemu",
	}


	//mengembalikan greeting message
	return formats[rand.Intn(len(formats))]
}