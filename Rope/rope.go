package main

import "log"

type Rope struct {
	char string
	weight int // length
	left *Rope // node left
	right *Rope // node right
}

func (r *Rope) create(char string) *Rope {
	r.char = char
	r.weight = len(char)
	return r
}
// merge nodes
func merge(nodeFirst *Rope, nodeSecond *Rope) *Rope {
	newNode := new(Rope)
	newNode.left = nodeFirst
	newNode.right = nodeSecond
	newNode.char = ""
	newNode.weight = nodeFirst.weight + nodeSecond.weight
	return newNode
}
func split(node *Rope, i int) (*Rope, *Rope) {
	tree1, tree2 := new(Rope), new(Rope)
	if node.left != nil {
		if node.left.weight >= i {
			tree1, tree2.left = split(node.left, i)
			tree2.right = node.right
			tree2.weight = tree2.left.weight + tree2.left.weight
		} else {
			tree1.right, tree2 = split(node.right, i - node.left.weight)
			tree1.left = node.left
			tree1.weight = tree1.left.weight + tree1.right.weight
		}
	} else {
		tree1.char = node.char[0:i]
		tree2.char = node.char[i:len(node.char)]
		tree1.weight = i
		tree2.weight = len(node.char) - i
	}
	return  tree1, tree2
}
// get char to index
func (r *Rope) get(i int) string {
	// check node is leaf or not
	if r.left != nil {
		// node is composite
		if r.left.weight >= i {
			return r.get(i)
		} else {
			return r.get(r.left.weight)
		}
	} else {
		// node is leaf, return char
		return string(r.char[i])
	}
}


func main(){
	node := new(Rope)
	node.create("Hello")
	//
	//nodeSecond := new(Rope)
	//nodeSecond.create(" world")
	//
	//nodeMerge := merge(node, nodeSecond)

	log.Println(split(node, 2))
}