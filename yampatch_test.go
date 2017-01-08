package yampatch_test

import (
	. "github.com/hyenaspots/yampatch"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yampatch", func() {
	var target, delta, desiredResult string

	BeforeEach(func() {
		target = `---
key: 1

key2:
  nested:
    super_nested: 2
  other: 3

array: [4,5,6]

items:
- name: item7
- name: item8`

		delta = `---
- type: replace
path: /key
value: 10

- type: replace
path: /new_key?
value: 10

- type: replace
path: /key2/nested/super_nested
value: 10

- type: replace
path: /key2/nested?/another_nested/super_nested
value: 10

- type: replace
  path: /array/0
  value: 10

- type: replace
  path: /array/-
  value: 10

- type: replace
  path: /array2?/-
  value: 10

- type: replace
  path: /items/name=item7/count
  value: 10

- type: replace
  path: /items/name=item8/count
  value: 10

- type: replace
  path: /items/name=item9?/count
  value: 10`

		desiredResult = `---
key: 10

key2:
	nested:
		super_nested: 10
	other: 3
	another_nested:
		super_nested: 10

array: [10,5,6,10]

array2: [10]

items:
- name: item7
	count: 10
- name: item8
- name: item9
	count: 10`

	})

	Context("Given a valid target string and a valid delta string with multiple operations", func() {
		It("Returns the post-op string", func() {
			actualResult, err := ApplyOps(target, delta)

			Expect(actualResult).To(Equal(desiredResult))
			Expect(err).To(BeNil())
		})
	})

	Context("Given a valid target string and an invalid delta string", func() {
		BeforeEach(func() {
			delta = `---
- type: replace
path: /key_not_there
value: 10`
			desiredResult = ""
		})

		It("Returns an error", func() {
			result, err := ApplyOps(target, delta)

			Expect(result).To(Equal(desiredResult))
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Given a valid target string and a valid delta string with only one operation", func() {
		BeforeEach(func() {
			target = `key: 1`

			delta = `- type: replace
			path: /key
			value: 10`
		})

		It("Returns the post-op string", func() {
			desiredResult = `key: 10`
			actualResult, err := ApplyOps(target, delta)

			Expect(actualResult).To(Equal(desiredResult))
			Expect(err).To(BeNil())
		})
	})
})
