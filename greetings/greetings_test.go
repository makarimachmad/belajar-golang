package greetings

import(
	"testing"
	"regexp"
)

//fungsi validasitesnama pada sebuah nama untuk
//mengecek apakah mengembalikan nilai dengan benar 
func TestHelloName(t *testing.T){
	nama := "makarim"
	want := regexp.MustCompile(`\b`+nama+`\b`)
	msg, err := Hello("makarim")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`assalamualaikum("makarim") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// fungsi untuk gk ada nama
func TestHelloEmpty(t *testing.T){
	msg, err := Hello("")
	if msg != "" || err == nil{
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
