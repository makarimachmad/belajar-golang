package main

import "fmt"
import "strconv"

func main(){
	fmt.Println(BinerCount("1111"))
	fmt.Println(Identic("abba"))
	fmt.Println(Identic("aaabbbba"))
	fmt.Println(Identic("bbbbaaab"))
	fmt.Println(Identic("aaaabbbb"))
	fmt.Println(Identic("aabbbbab"))
}

func BinerCount(S string) int {
    nilai := 0
    tampung := []int{}
    if i, err := strconv.ParseInt(S, 2, 64); err != nil {
        fmt.Println(err)
    } else {
        nilai = int(i)
    }
    for {
        if nilai == 0 {
		break
	    }  
        if nilai % 2 == 0 {
	    nilai = nilai/2
            tampung = append(tampung, nilai)
        }
        if nilai % 2 == 1 {
	    nilai = (nilai-1)/2
            tampung = append(tampung, nilai)
        }      
    }
    return len(tampung)
}

func Identic(s string) bool{
	fmt.Println("len(s): ", len(s))
	for i:= 1; i<=len(s); i++{
		fmt.Println("s[i-1]: ", string(s[i-1]), " s[i]: ", string(s[i]))
		if string(s[i-1]) == "b" && string(s[i]) == "a"{
			return false
		}
	}
	return true
}