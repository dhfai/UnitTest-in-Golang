package main

import (
	"Tugas-1/Book"
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func main() {
	Book.InitDatabase()

	fmt.Print("Masukkan kata kunci pencarian: ")
	var keyword string
	fmt.Scanln(&keyword)

	searchResults := Book.SearchBooks(keyword)

	if len(searchResults) == 0 {
		fmt.Println("Tidak ditemukan buku yang sesuai dengan kata kunci pencarian.")
	} else {
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"#", "Title", "Author", "Genre"})

		for i, book := range searchResults {
			t.AppendRow([]interface{}{i + 1, book.Title, book.Author, book.Genre})
		}

		fmt.Println("\nHasil Pencarian")
		fmt.Println("Keyword yang dicari:", keyword)
		fmt.Println("--------------------------")
		t.Render()
	}
}
