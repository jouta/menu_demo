package main

import "fmt"

//链表结点定义
type LinkTableNode struct{
	data interface{}
	next *LinkTableNode
}

//链表定义
type LinkTable struct{
	length int
	head *LinkTableNode
	tail *LinkTableNode
}

func NewLinkTableNode(data interface{}) *LinkTableNode{
	return &LinkTableNode{
		data : data,
		next : nil,
	}
}

//创建LinkTable
func  CreateLinkTable() *LinkTable{
	return  &LinkTable{
		length : 0,
		head :  nil,
		tail :  nil,
	}
}

func IsEmpty(LinkTable *LinkTable) bool{
	return LinkTable.length == 0
}

//添加结点
func AddLinkTableNode(LinkTable *LinkTable, Node *LinkTableNode) bool {
	if LinkTable == nil || Node == nil {
		return false
	}
	if IsEmpty(LinkTable) {
		LinkTable.head = Node
		LinkTable.tail = Node
		LinkTable.length ++
	} else {
		LinkTable.tail.next = Node
		LinkTable.tail = Node
		LinkTable.length ++
	}
	return true
}

//删除结点
func DeleteLinkTableNode(LinkTable *LinkTable, node *LinkTableNode) bool {
	if LinkTable == nil || node == nil {
		return false
	}
	if node == LinkTable.head {
		if LinkTable.length == 1 {
			LinkTable.head = nil
			LinkTable.tail = nil
		} else {
			LinkTable.head = LinkTable.head.next
		}
		LinkTable.length--
		return true
	}

	var p *LinkTableNode = LinkTable.head
	for p.next != LinkTable.tail {
		if p.next == node {
			p = p.next.next
			LinkTable.length--
			return true
		}
		p = p.next
	}
	if p.next == node {
		LinkTable.tail = p
		LinkTable.length--
		return true
	}
	return false
}

//查找结点
func SelectLinkTableNode(LinkTable *LinkTable, node *LinkTableNode) bool {
	if LinkTable == nil || node == nil {
		return false
	}
	var p *LinkTableNode = LinkTable.head
	for p != nil {
		if p == node {
			return true
		}
		p = p.next
	}
	return false
}


func main() {

	var LinkTable *LinkTable = CreateLinkTable()
	var node1 *LinkTableNode = NewLinkTableNode("这是一个结点")
	AddLinkTableNode(LinkTable, node1)

	var node2 *LinkTableNode = NewLinkTableNode("这是另一个结点")
	AddLinkTableNode(LinkTable, node2)

	fmt.Println("长度:", LinkTable.length)
	fmt.Println("头:", LinkTable.head.data)
	fmt.Println("尾:", LinkTable.tail.data)

	fmt.Println("查找node1:",SelectLinkTableNode(LinkTable, node1))

	DeleteLinkTableNode(LinkTable, node1)
	fmt.Println("删除node1后长度:", LinkTable.length)
	fmt.Println("删除node1后头:", LinkTable.head.data)
	fmt.Println("删除node1后尾:", LinkTable.tail.data)

	fmt.Println("查找node1:",SelectLinkTableNode(LinkTable, node1))
}