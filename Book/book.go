package Book

import (
	"strings"
)

type Book struct {
	Title  string
	Author string
	Genre  string
}

var Database []Book

func InitDatabase() {
	Database = append(Database, Book{"Classroom of the elite", "Shōgo Kinugasa", "Slice of Life"})
	Database = append(Database, Book{"Funiculi Funicula", "Toshikazu Kawaguchi", "Romance"})
	Database = append(Database, Book{"The Great Gatsby", "F. Scott Fitzgerald", "Fiction"})
}

func SearchBooks(keyword string) []Book {
	var results []Book
	keyword = strings.ToLower(keyword)
	for _, book := range Database {
		if strings.Contains(strings.ToLower(book.Title), keyword) ||
			strings.Contains(strings.ToLower(book.Author), keyword) ||
			strings.Contains(strings.ToLower(book.Genre), keyword) {
			results = append(results, book)
		}
	}
	return results
}
