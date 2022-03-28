package main

import "fmt"

func main() {

	//array
	var bahasa[5] string
	bahasa[0] = "python"
	bahasa[1] = "golang"
	bahasa[2] = "java"
	bahasa[3] = "c++"
	bahasa[4] = "R"

	fmt.Println(bahasa)

	bahasav2 := [...] string{
		"golang",
		"python",
		"R",
		"Kotlin",
		"C++",
		"Java",
	}

	for _, bsa := range bahasav2{
		fmt.Println(bsa)
	}

	//slice
	var gamingplaystationconsoles []string

	gamingplaystationconsoles = append(gamingplaystationconsoles, "ps4")
	gamingplaystationconsoles = append(gamingplaystationconsoles, "xbox")
	gamingplaystationconsoles = append(gamingplaystationconsoles, "nintendo")

	fmt.Println(gamingplaystationconsoles)

	gamingconsoles := []string{"ps4","xbox","nintendo"}
	fmt.Println(gamingconsoles)

	for _, bsa := range gamingconsoles{
		fmt.Println(bsa)
	}

	//map
	//var mymap map[string]int
	mymap := map[string]int{}

	mymap["python"]=10
	mymap["go"]=9
	mymap["javascript"]=8
	mymap["python"]=6

	fmt.Println(mymap)

	mymapv2 := map[string]string{"golang":"cepat banget", "python":"bagus"}
	fmt.Println(mymapv2)

	mymapv3 := map[string]string{
		"golang":"cepat banget", 
		"python":"bagus",}
	fmt.Println(mymapv3)

	for key, value := range mymapv2{
		fmt.Println("key",key,"value",value)
	}
	delete(mymapv2,"golang")
	fmt.Println(mymapv2)

	value, adagk := mymapv2["java"]
	fmt.Println(value, adagk)

	//slice map
	siswa := []map[string]string{
		{"nama":"makarim","nilai":"A"},
		{"nama":"widyanto","nilai":"B"},
	}
	fmt.Println(siswa)

	for _, murid := range siswa{
		fmt.Println(murid["nama"],"-",murid["nilai"])
	}

	//quiz
	nilai := [8]int{100,80,75,92,70,93,88,67}
	var bantu int

	for _, isi := range nilai{
		bantu = bantu + isi
	}
	fmt.Println("nilai: ", bantu, "rata-rata: ", float64(bantu)/float64(len(nilai)))

	var nilaibagus []int
	for _, nb := range nilai{
		if nb > 90{
			nilaibagus = append(nilaibagus, nb)
		}
	}
	fmt.Println(nilaibagus)
}