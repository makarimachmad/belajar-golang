package greetings
import "fmt"

//fungsi hello mengembalikan nilai nama yg ada pada pesan
func Hello(nama string) string{
	message := fmt.Sprintf("Asalamualaikum, %v. selamat datang", nama)
	return message
}