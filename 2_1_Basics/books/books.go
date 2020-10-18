package books

import "strconv"

type Page []string
type Book []Page
type Index map[string][]int

func (x Book) String() string {
	res := "Book (" + strconv.Itoa(len(x)) + " Pages)\n"
	for idx, page := range x {
		res += "[" + strconv.Itoa(idx) + "]: "
		for _, word := range page {
			res += word + " "
		}
		res += "\n"
	}
	return res
}

func (x Index) String() string {
	res := "Index (" + strconv.Itoa(len(x)) + " Entries)\n"
	for idx, pages := range x {
		res += "[" + idx + "]: "
		for _, pageNumber := range pages {
			res += strconv.Itoa(pageNumber) + ","
		}
		res += "\n"
	}
	return res
}

func CreateIndex(book Book) Index {
	index := make(Index)

	for pageNumber, page := range book {
		for _, word := range page {
			pages := index[word]
			pages = append(pages, pageNumber)
			index[word] = pages
		}
	}

	return index
}

func Contains(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
