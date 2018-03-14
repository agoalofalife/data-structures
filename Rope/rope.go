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
func (r *Rope) merge(nodeMerge *Rope) *Rope {
	newNode := new(Rope)
	newNode.left = r
	newNode.right = nodeMerge
	newNode.char = ""
	newNode.weight = r.weight + nodeMerge.weight
	return newNode
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

	nodeSecond := new(Rope)
	nodeSecond.create(" world")

	node.merge(nodeSecond)

	log.Println(node)
}