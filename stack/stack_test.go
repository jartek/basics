package stack_test

import (
	. "github.com/jartek/basics/stack"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	var (
		stack Stack
	)

	Context("Empty Stack", func() {
		BeforeEach(func() {
			stack = Stack{}
		})

		Describe("Display", func() {
			It("should return an empty array", func() {
				Expect(stack.Display()).To(Equal([]int{}))
			})
		})

		Describe("Length", func() {
			It("should have 0 length", func() {
				Expect(stack.Length()).Should(BeZero())
			})
		})

		Describe("Peek", func() {
			It("should return -1", func() {
				Expect(stack.Peek()).To(Equal(NotFound))
			})
		})

		Describe("Push", func() {
			It("should increment length", func() {
				Expect(stack.Length()).Should(BeZero())
				stack.Push(1)
				Expect(stack.Length()).To(Equal(1))
			})

			It("should insert element at head", func() {
				stack.Push(1)
				Expect(stack.Peek()).To(Equal(1))
				stack.Push(2)
				Expect(stack.Peek()).To(Equal(2))
			})
		})
		Describe("Pop", func() {
			It("should return an error", func() {
				Expect(stack.Pop()).To(Equal(NotFound))
			})
		})
	})

	Context("Existing Stack", func() {
		BeforeEach(func() {
			stack = Stack{}
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
		})

		Describe("Display", func() {
			It("should return an array with stack elements", func() {
				Expect(stack.Display()).To(Equal([]int{3, 2, 1}))
			})
		})

		Describe("Length", func() {
			It("should have a length 3", func() {
				Expect(stack.Length()).Should(Equal(3))
			})
		})

		Describe("Peek", func() {
			It("should return the element at head", func() {
				Expect(stack.Peek()).To(Equal(3))
			})
		})

		Describe("Push", func() {
			It("should push to the existing stack", func() {
				stack.Push(4)
				Expect(stack.Display()).To(Equal([]int{4, 3, 2, 1}))
			})
		})

		Describe("Pop", func() {
			It("should decrement length", func() {
				stack.Pop()
				Expect(stack.Length()).To(Equal(2))
			})

			It("should pop the head element", func() {
				Expect(stack.Pop()).To(Equal(3))
				Expect(stack.Display()).To(Equal([]int{2, 1}))
			})
		})
	})
})
