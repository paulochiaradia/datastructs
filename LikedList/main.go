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
	if ll.headNode == nil {
		return nil
	}

	node := ll.headNode
	for node.nextNode != nil {
		node = node.nextNode
	}
	return node
}

func (ll *LinkedList) AddToEnd(property int) {
	node := &Node{
		property: property,
		nextNode: nil,
	}
	lastNode := ll.LastNode()
	if lastNode == nil {
		ll.headNode = node
	} else {
		lastNode.nextNode = node
	}
}

func (ll *LinkedList) NodeWithValue(value int) (*Node, error) {
	node := ll.headNode
	for node != nil {
		if node.property == value {
			return node, nil
		}
		node = node.nextNode
	}
	return nil, errors.New("node doesn't exist")
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
	if ll.headNode != nil {
		ll.headNode = ll.headNode.nextNode
	}
}

func (ll *LinkedList) DeleteEnd() {
	if ll.headNode == nil {
		return
	}

	if ll.headNode.nextNode == nil {
		ll.headNode = nil
		return
	}

	node := ll.headNode
	var previous *Node
	for node.nextNode != nil {
		previous = node
		node = node.nextNode
	}
	previous.nextNode = nil
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
	var length int
	for node := ll.headNode; node != nil; node = node.nextNode {
		length++
	}
	return length
}

func (ll *LinkedList) HasCycle() bool {
	slow := ll.headNode
	fast := ll.headNode

	for slow != nil && fast != nil && fast.nextNode != nil {
		slow = slow.nextNode
		fast = fast.nextNode.nextNode

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	listaLigada := LinkedList{}
	listaLigada.AddToEnd(1)
	listaLigada.AddToEnd(2)
	listaLigada.AddToEnd(3)
	listaLigada.AddToEnd(4)
	listaLigada.PrintList()
	//listaLigada.DeleteSearch(8)
	//fmt.Println("------")
	//listaLigada.PrintList()
	//fmt.Println("------")
	//fmt.Println(listaLigada.LenghtList())
	listaLigada.headNode.nextNode.nextNode.nextNode = listaLigada.headNode

	if listaLigada.HasCycle() {
		fmt.Println("List contains a cycle")
	} else {
		fmt.Println("List does not contain a cycle")
	}
}
