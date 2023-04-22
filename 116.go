/*
	You are given a perfect binary tree where all leaves are on the same level, and every parent has two children. The binary tree has the following definition:

	Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.

	Initially, all next pointers are set to NULL.

	Example 1:

	Input: root = [1,2,3,4,5,6,7]
	Output: [1,#,2,3,#,4,5,6,7,#]
	Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
	Example 2:

	Input: root = []
	Output: []


	Constraints:

	The number of nodes in the tree is in the range [0, 212 - 1].
	-1000 <= Node.val <= 1000


	Follow-up:

	You may only use constant extra space.
	The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.
*/

package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := []*Node{root}
	for len(queue) != 0 {
		l := len(queue)
		for i, j := 0, 1; i < l-1 && j < l; i, j = i+1, j+1 {
			queue[i].Next = queue[j]
		}
		for i := 0; i < l; i++ {
			if queue[i].Left == nil {
				break
			}
			queue = append(queue, queue[i].Left, queue[i].Right)
		}
		queue = queue[l:]
	}
	return root
}

func main() {
	root := &Node{Val: 1, Next: nil, Left: &Node{Val: 2}, Right: &Node{Val: 3}}

	println(connect(root).Left.Next.Next)
}
