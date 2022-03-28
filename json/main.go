package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dimensions struct {
	Height int
	Width int
}

type Bird struct {
	Species string
	Description string
	Dimensions Dimensions
}

type Book struct {
	Title  string
	Author string

	// extra is used for additional dynamic element marshalling
	extra func() interface{} `json:"-"`
}

type FakeBook Book

func (b *Book) SetExtra(fn func() interface{}) {
	b.extra = fn
}

func (b *Book) MarshalJSON() ([]byte, error) {
	if b.extra == nil {
		b.extra = func() interface{} { return *b }
	}

	return json.Marshal(b.extra())
}

func main() {
	ms := &Book{
		Title:  "Catch-22",
		Author: "Joseph Heller",
	}

	ms.SetExtra(func() interface{} {
		return struct {
			FakeBook
			Extra1 struct {
				Something string `json:"something"`
			} `json:"extra1"`
		}{
			FakeBook: FakeBook(*ms),
			Extra1: struct {
				Something string `json:"something"`
			}{
				Something: "yy",
			},
		}
	})

	out, err := json.MarshalIndent(ms, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))

	mb := &Book{
		Title:  "Vim-go",
		Author: "Fatih Arslan",
	}

	mb.SetExtra(func() interface{} {
		return struct {
			FakeBook
			Something string `json:"something"`
		}{
			FakeBook:  FakeBook(*mb),
			Something: "xx",
		}
	})

	out, err = json.MarshalIndent(mb, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))

	mc := &Book{
		Title:  "Another-Title",
		Author: "Fatih Arslan",
	}

	out, err = json.MarshalIndent(mc, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))

	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var birds Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Print(birds.Dimensions)
}