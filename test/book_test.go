package golang_ginkgo_best_practices_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/leobastiani/golang-ginkgo-best-practices/pkg/books"
)

var _ = Describe("Book", func() {
  var foxInSocks, lesMis *books.Book

  BeforeEach(func() {
    lesMis = &books.Book{
      Title:  "Les Miserables",
      Author: "Victor Hugo",
      Pages:  2783,
    }

    foxInSocks = &books.Book{
      Title:  "Fox In Socks",
      Author: "Dr. Seuss",
      Pages:  24,
    }
  })

  Describe("Categorizing books", func() {
    Context("with more than 300 pages", func() {
      It("should be a novel", func() {
        Expect(lesMis.Title).To(Equal("Les Miserables"))
      })
    })

    Context("with fewer than 300 pages", func() {
      It("should be a short story", func() {
        Expect(foxInSocks.Title).To(Equal("Fox In Socks"))
      })
    })
  })
})
