package main

import (
	"errors"
	"fmt"
)

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type DoubleLinkedList struct {
	headNode *Node
	endNode  *Node
}

func (dll *DoubleLinkedList) AddToHead(property int) {
	node := &Node{
		property:     property,
		nextNode:     nil,
		previousNode: nil,
	}

	if dll.headNode != nil {
		node.nextNode = dll.headNode
		dll.headNode.previousNode = node
	} else {
		dll.endNode = node
	}
	dll.headNode = node
}

func (dll *DoubleLinkedList) AddToEnd(property int) {
	node := &Node{
		property:     property,
		nextNode:     nil,
		previousNode: nil,
	}
	if dll.endNode == nil {
		dll.headNode = node
		dll.endNode = node
	} else {
		node.previousNode = dll.endNode
		dll.endNode.nextNode = node
		dll.endNode = node
	}

}

func (dll *DoubleLinkedList) PrintHeadToEnd() {
	node := &Node{}
	for node = dll.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (dll *DoubleLinkedList) PrintEndToHead() {
	node := &Node{}
	for node = dll.endNode; node != nil; node = node.previousNode {
		fmt.Println(node.property)
	}
}

func (dll *DoubleLinkedList) SearchNodeWithValue(value int) (*Node, error) {
	node := &Node{}

	for node = dll.headNode; node != nil; node = node.nextNode {
		if node.property == value {
			return node, nil
		}
	}
	return nil, errors.New("node searched not found")
}

func (dll *DoubleLinkedList) AddAfter(property int, value int) {
	node, err := dll.SearchNodeWithValue(value)
	if err != nil {
		_ = errors.New("node searched not found")
		return
	}
	nodeToBeAdd := &Node{
		property:     property,
		nextNode:     nil,
		previousNode: nil,
	}
	if node.nextNode != nil {
		node.nextNode.previousNode = nodeToBeAdd
	} else {
		dll.endNode = nodeToBeAdd // Atualiza endNode se node é o último nó
	}
	node.nextNode = nodeToBeAdd

}

func main() {
	doubleLinkedList := DoubleLinkedList{}
	//doubleLinkedList.AddToHead(5)
	//doubleLinkedList.AddToHead(10)
	//doubleLinkedList.AddToHead(60)
	//doubleLinkedList.AddToHead(20)
	doubleLinkedList.AddToEnd(55)
	doubleLinkedList.AddToEnd(102)
	doubleLinkedList.AddAfter(100, 55)
	doubleLinkedList.PrintHeadToEnd()
	fmt.Println("---------")
	doubleLinkedList.PrintEndToHead()
	//node, _ := doubleLinkedList.SearchNodeWithValue(102)
	//fmt.Println(node)
}
