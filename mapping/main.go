package main
import (
    "fmt"
    "encoding/json"
)

type MyAddress struct {
    House string
    School string
}
type Student struct {
    Id int64
    Name string
    Scores float32
    Address MyAddress
    Labels []string
}

func Test() {

    dict := make(map[string]interface{})
    dict["id"] = 201902181425       // int
    dict["name"] = "jackytse"       // string
    dict["scores"] = 123.456        // float
    dict["address"] = map[string]string{"house":"my house", "school":"my school"}   // map
    dict["labels"] = []string{"aries", "warmhearted", "frank"}      // slice

    jsonbody, err := json.Marshal(dict)
    if err != nil {
        // do error check
        fmt.Println(err)
        return
    }

    student := Student{}
    if err := json.Unmarshal(jsonbody, &student); err != nil {
        // do error check
        fmt.Println(err)
        return
    }

    fmt.Printf("%#v\n", student)
}

func main() {
    Test()
}