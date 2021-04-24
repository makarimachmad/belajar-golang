package main

import "fmt"

type User struct {
	Id int
	Nama string
	Alamat string
	Status bool
}
//embeded struct
type Grup struct {
	Nama string
	Admin User
	Users []User
	Status bool
}

func main(){
	
	user1 := User{}
	user1.Id = 1
	user1.Nama = "makarim"
	user1.Alamat = "Yogyakarta"
	user1.Status = true

	fmt.Println(user1)

	user2 := User{2, "widyanto", "surabaya", true}
	fmt.Println(user2)

	user3 := dsiplayuser(user2)
	fmt.Println(user3)

	//embeded struct
	users := []User{user1, user2}
	grup := Grup{"Bahasa Pemrograman", user1, users, true}
	displaygrup(grup)

	//method
	usermethod := user1.display()
	fmt.Println(usermethod)

	//quiz
	grup.grupmethod()
	

}

//struct sebagai parameter
func dsiplayuser(user User) string{
	hasil := fmt.Sprintf("nama: %s, alamat: %s, status: %t", user.Nama, user.Alamat, user.Status)
	return hasil
}

//embeded struct
func displaygrup(grup Grup) {
	fmt.Printf("Nama Grup: ", grup.Nama)
	fmt.Println("")
	fmt.Printf("Jumlah Anggota: %d", len(grup.Users))
	fmt.Println("")

	fmt.Println("Nama Anggota: ")
	for _, user := range grup.Users{
		fmt.Println(user.Nama)
	}
}

//method
func (user User) display() string{
	return fmt.Sprintf("nama: %s, alamat: %s, status: %t", user.Nama, user.Alamat, user.Status)
}

//quiz
func (grup Grup) grupmethod(){
	fmt.Printf("Nama Grup: ", grup.Nama)
	fmt.Println("")
	fmt.Printf("Jumlah Anggota: %d", len(grup.Users))
	fmt.Println("")

	fmt.Println("Nama Anggota: ")
	for _, user := range grup.Users{
		fmt.Println(user.Nama)
	}
} 