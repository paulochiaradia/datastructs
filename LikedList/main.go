package main

import (
	"errors"
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (ll *LinkedList) AddToHead(property int) {
	node := &Node{property: property, nextNode: nil}
	if ll.headNode != nil {
		node.nextNode = ll.headNode
	}
	ll.headNode = node
}

func (ll *LinkedList) PrintList() {
	for node := ll.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (ll *LinkedList) LastNode() *Node {
	node := &Node{}
	for node = ll.headNode; node.nextNode != nil; node = node.nextNode {
	}
	return node
}

func (ll *LinkedList) AddToEnd(property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}
	lastNode := ll.LastNode()
	lastNode.nextNode = node
}

func (ll *LinkedList) NodeWithValue(value int) (*Node, error) {
	node := &Node{}
	for node = ll.headNode; node != nil; node = node.nextNode {
		if node.property == value {
			return node, nil
		}
	}
	return nil, errors.New("node doesnt exist")
}

func (ll *LinkedList) AddAfter(property int, find int) {
	node, err := ll.NodeWithValue(find)
	if err != nil {
		fmt.Println("Node searched not found")
		return
	}
	nodeInsert := &Node{
		property: property,
		nextNode: nil,
	}
	nodeInsert.nextNode = node.nextNode
	node.nextNode = nodeInsert

}

func (ll *LinkedList) DeleteHead() {
	ll.headNode = ll.headNode.nextNode
}

func (ll *LinkedList) DeleteEnd() {
	node := &Node{}
	previus := &Node{}
	for node = ll.headNode; node.nextNode != nil; node = node.nextNode {
		previus = node
	}
	previus.nextNode = nil
}

func (ll *LinkedList) DeleteSearch(property int) {
	if ll.headNode == nil {
		return // Lista está vazia, nada a deletar
	}

	// Verificar se o nó a ser deletado é o headNode
	if ll.headNode.property == property {
		ll.headNode = ll.headNode.nextNode
		return
	}

	// Inicializar node e nodePrevius corretamente
	node := ll.headNode
	var nodePrevius *Node

	// Percorrer a lista para encontrar o nó a ser deletado
	for node != nil && node.property != property {
		nodePrevius = node
		node = node.nextNode
	}

	if node != nil {
		nodePrevius.nextNode = node.nextNode
		return
	}

	fmt.Println("Node searched not found")
}

func (ll *LinkedList) LenghtList() int {
	var lenght int
	var node *Node

	for node = ll.headNode; node != nil; node = node.nextNode {
		lenght++
	}
	return lenght
}
func main() {
	listaLigada := LinkedList{}
	listaLigada.AddToHead(4)
	listaLigada.AddToHead(7)
	listaLigada.AddToHead(10)
	listaLigada.AddToHead(9)
	listaLigada.PrintList()
	listaLigada.DeleteSearch(8)
	fmt.Println("------")
	listaLigada.PrintList()
	fmt.Println("------")
	fmt.Println(listaLigada.LenghtList())

}
