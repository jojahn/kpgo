package books

import (
	"fmt"
	"testing"
)

func TestBookToString(t *testing.T) {
	p := Page{"a","b","c"}
	p2 := Page{"d","e","f"}
	b := Book{p, p2}
	fmt.Println(b)
}

func TestIndexToString(t *testing.T) {
	i := make(Index)
	pages := i["a"]
	i["a"] = append(pages, 2)
	fmt.Println(i)
}

func TestCreateIndex(t *testing.T) {
	p1 := Page{"Hello", "our", "World"}
	p2 := Page{"Hello", "Go", "Language"}
	p3 := Page{"Hello", "C++", "Language"}
	book := Book{p1,p2,p3}
	i := CreateIndex(book)
	fmt.Println(i)

	if !Contains(i["Hello"], 0) || !Contains(i["Hello"], 1) || !Contains(i["Hello"], 2) {
		t.Error("Hello is not included in all pages")
	}
}
