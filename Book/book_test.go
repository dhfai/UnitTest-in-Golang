package Book

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/jedib0t/go-pretty/table"
)

func TestSearchBooks(t *testing.T) {
	InitDatabase()

	tests := []struct {
		keyword string
		want    []Book
	}{
		{
			keyword: "Class",
			want: []Book{
				{"Classroom of the elite", "Shōgo Kinugasa", "Slice of Life"},
			},
		},
		{
			keyword: "Funiculi",
			want: []Book{
				{"Funiculi Funicula", "Toshikazu Kawaguchi", "Romance"},
			},
		},
		{
			keyword: "Fitzgerald",
			want: []Book{
				{"The Great Gatsby", "F. Scott Fitzgerald", "Fiction"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.keyword, func(t *testing.T) {
			got := SearchBooks(tt.keyword)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchBooks(%q) = %v, want %v", tt.keyword, got, tt.want)
			}
		})
	}

	fmt.Println("\nTest Summary:")
	tbl := table.NewWriter()
	tbl.SetOutputMirror(os.Stdout)
	tbl.AppendHeader(table.Row{"Keyword", "Expected", "Actual", "Result"})
	for _, tt := range tests {
		got := SearchBooks(tt.keyword)
		expected := fmt.Sprintf("%v", tt.want)
		actual := fmt.Sprintf("%v", got)
		result := "Pass"
		if !reflect.DeepEqual(got, tt.want) {
			result = "Fail"
		}
		tbl.AppendRow([]interface{}{tt.keyword, expected, actual, result})
	}
	tbl.Render()
}
