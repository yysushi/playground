package books_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	It("", func() {
		Expect("Book").To(Equal("Book"))
		// Expect("Book").To(Equal("Book2"))
		ExpectWithOffset(1, "Book").To(Equal("Book2"))
	})
})
