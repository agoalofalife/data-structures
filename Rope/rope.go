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
// get char to index
func (r *Rope) get(i int, rope *Rope) string {
	if r.left != nil {
		if r.left.weight >= i {
			return r.get(i, rope.left)
		} else {
			return  r.get(rope.left.weight, rope.right)
		}
	} else {
		return string(rope.char[i])
	}
}

func main(){
	node := new(Rope)
	node.create("Hello")
	log.Println(node)
}