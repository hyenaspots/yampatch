package yampatch_test

import (
	. "github.com/hyenaspots/yampatch"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yampatch", func() {
	var unpatchedYAML, operationsYAML, desiredResult string

	Context("Given a valid unpatchedYAML string and a valid operationsYAML string with multiple operations", func() {
		BeforeEach(func() {
			unpatchedYAML = `---
key1: 1
key2:`

			operationsYAML = `---
- type: remove
  path: /key2

- type: replace
  path: /key1
  value: 10

- type: replace
  path: /key3?
  value: 10`

			desiredResult = `key1: 10
key3: 10
`
		})

		It("Returns the post-op string", func() {
			actualResult, err := ApplyOps(unpatchedYAML, operationsYAML)

			Expect(err).To(BeNil())
			Expect(actualResult).To(Equal(desiredResult))
		})
	})

	Context("Given a valid unpatchedYAML string and an invalid operationsYAML string", func() {
		BeforeEach(func() {
			operationsYAML = `---
- type: replace
path: /key_not_there
value: 10`
			desiredResult = ""
		})

		It("Returns an error", func() {
			result, err := ApplyOps(unpatchedYAML, operationsYAML)

			Expect(err).To(HaveOccurred())
			Expect(result).To(Equal(desiredResult))

		})
	})

	Context("Given a valid unpatchedYAML string and a valid operationsYAML string with only one operation", func() {
		BeforeEach(func() {
			unpatchedYAML = `key: 1`

			operationsYAML = `
- type: replace
  path: /key
  value: 10`
		})

		It("Returns the post-op string", func() {
			desiredResult = `key: 10
`
			actualResult, err := ApplyOps(unpatchedYAML, operationsYAML)

			Expect(err).To(BeNil())
			Expect(actualResult).To(Equal(desiredResult))

		})
	})
})
