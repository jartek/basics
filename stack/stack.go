package stack

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head   *Node
	length int
}

const (
	NotFound = -1
)

func (stack *Stack) Display() []int {
	ret := []int{}
	head := stack.head
	for i := 0; i < stack.Length(); i++ {
		ret = append(ret, head.data)
		head = head.next
	}
	return ret
}

func (stack *Stack) Length() int {
	return stack.length
}

func (stack *Stack) Peek() int {
	node := stack.head
	if node == nil {
		return NotFound
	}
	return node.data
}

func (stack *Stack) Push(item int) *Stack {
	node := &Node{data: item}
	if stack.head != nil {
		node.next = stack.head
	}
	stack.length++
	stack.head = node
	return stack
}

func (stack *Stack) Pop() int {
	if stack.Length() == 0 {
		return NotFound
	}
	node := stack.head
	stack.head = node.next
	stack.length--
	return node.data
}
