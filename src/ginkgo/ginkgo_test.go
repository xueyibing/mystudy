package ginkgo

import (
	"fmt"
	. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"

"testing"
)

type Book struct {
	Title string
	Author string
	Pages int
}
func TestBooks(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books")
}

var _ = Describe("Book", func() {
	var (
		longBook  Book
		shortBook Book
	)

	BeforeEach(func() {
		longBook = Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Context("context", func() {
		It("it", func() {
			fmt.Println("hahaha")
			fmt.Println(longBook)
			fmt.Println(shortBook)
		})
	})
})