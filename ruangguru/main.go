package main

import (
	// "fmt"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var Wait sync.WaitGroup
var Counter int = 0

func main() {

	// submission
	// for routine := 1; routine <= 2; routine++ {
	// 	// fmt.Println("sebelum wait add: ", routine)
	// 	Wait.Add(1)
	// 	// fmt.Println("setelah wait add: ", routine)
	// 	go Routine(routine)
	// 	// Routine(routine)
	// 	// fmt.Println("setelah go routine add: ", routine)
	// }
	// // fmt.Printf("sebelum Final Counter: %d\n", Counter)
	// Wait.Wait()
	// fmt.Printf("Final Counter: %d\n", Counter)

	//1.
	// abc := []int{5,7,8,3}
	// fmt.Println(Pertama(abc))

	//2.
	// x := []int{11,10,10,5,10,15,20,10,7,11}
	// fmt.Println(Foo(x, 8,18,3,6))
	// fmt.Println(Foo(x,10,20,0,9))
	// fmt.Println(Foo(x,8,18,6,3))
	// fmt.Println(Foo(x,20,10,0,9))
	// fmt.Println(Foo(x,20,10,0,9))
	// fmt.Println(Foo(x,6,7,8,8))
	// n:=0
	// for n != 10{
	// 	fmt.Println("n: ", n)
	// 	n = n+1
	// }

	//3
	// fmt.Println(h(1, "fruits"))
	// fmt.Println(h(2, "fruits"))
	// fmt.Println(h(5, "fruits"))
	// fmt.Println(h(pow(2, 2), "fruits"))
	// fmt.Println(h(pow(2, 1), "fruits"))

	dict := map[int][]int{}

	scanner := bufio.NewScanner(os.Stdin)
	dari := bufio.NewScanner(os.Stdin)
	sampai := bufio.NewScanner(os.Stdin)
	jarak := bufio.NewScanner(os.Stdin)
	fmt.Print("berapa banyak: ")
	scanner.Scan()
	i:=0
	banyak,_ := strconv.Atoi(scanner.Text())
	for i < banyak{
		fmt.Println("CASE ",i+1)
		fmt.Print("Bilangan dari: ")
		dari.Scan()
		dari,_ := strconv.Atoi(dari.Text())
		fmt.Print("Bilangan sampai: ")
		sampai.Scan()
		sampai,_ := strconv.Atoi(sampai.Text())
		fmt.Print("Jarak: ")
		jarak.Scan()
		jarak,_ := strconv.Atoi(jarak.Text())
		for dari <= sampai{
			if dari % jarak == 0 {
				dict[i] = append(dict[i], dari)
			}
			dari = dari + 1
		}
		i += 1
	}
	for j,_ := range dict{
		fmt.Println("Case", j+1, ": ", len(dict[j]))
	}
}

func Routine(id int) {

	for count := 0; count < 2; count++ {
		value := Counter
		value++
		Counter = value
	}

	Wait.Done()
}

func Pertama(nilai []int) int{
	
	var s = 0
	for _, k := range nilai{
		s = s + k
	}
	return s
}

func Foo(x []int, a int, b int, i int, j int) int{
	k := j
	ct := 0
	
	for k > i-1{
		if x[k] <= b && x[k] > a{
			ct = ct +1
		}
		k = k-1
	}
	return ct
}

func g(str string) string{
	i:=0
	new_str := ""
	for i < len(str)-1{
		new_str = new_str + string(str[i+1])
		i = i+1
	}
	return new_str
}

func f(str string) string{
	if len(str) == 0{
		return ""
	}else if len(str) == 1{
		return str
	}else{
		return f(g(str)) + string(str[0])
	}
}

func h(n int, str string) string{
	
	for n!=1{
		if n%2 == 0{
			n = n%2
		}else{
			n = 3*n + 1
		}
		str = f(str)
	}
	return str
}

func pow(x int, y int) int{
	if y==0{
		return 1
	}else{
		return x * pow(x, y-1)
	}
}