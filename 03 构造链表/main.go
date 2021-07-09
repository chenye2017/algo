package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val int
	Next *Node
	Prev *Node
}


/*
 需要实现的方法
get(index)
addAtHead(val)
addAtTail(val)
addAtIndex(index,val)
deleteAtIndex(index)
*/

type Link struct {
	Size int
	Head *Node
	Tail *Node
}

func (l *Link) Init() {
	l.Size = 0
	l.Head = new(Node)
	l.Tail = new(Node)
	l.Head.Next = l.Tail
	l.Tail.Prev = l.Head
}

func (l *Link) String() string {
	result := make([]int, 0)

	current := l.Head.Next

	for {
		if current == l.Tail {
			break
		} else {
			result = append(result, current.Val)
			current = current.Next
		}
	}

	strs := make([]string, 0, len(result))
	for _, v := range  result {
		strs = append(strs, strconv.Itoa(v))
	}
	return  strings.Join(strs, ",")
}

func (l *Link) AddAtHead(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
		Prev: nil,
	}

	originNext := l.Head.Next

	l.Head.Next = node
	node.Prev = l.Head
	node.Next = originNext
	originNext.Prev = node
	l.Size++
}

func (l *Link) AddAtTail(val int) {
	node := &Node{
		Val:  val,
		Next: nil,
		Prev: nil,
	}

	originPrev := l.Tail.Prev

	l.Tail.Prev = node
	node.Next = l.Tail
	originPrev.Next = node
	node.Prev = originPrev

	l.Size++
}

// index 是 1 开始
func (l *Link)Get(index int) (*Node, error) {

	if index < 1 || index > l.Size {
		return  nil, errors.New(" index 不合法")
	}

	next := l.Head
	for i:=0;i < index; i++ {
		next = next.Next
	}
	return  next, nil
}

// 在指定位置添加元素
func (l *Link)AddAtIndex(index int, val int) error{
	if index == 0 {
		l.AddAtHead(val)
		return nil
	}
	if index == l.Size + 1 {
		l.AddAtTail(val)
		return nil
	}

	indexNode,err := l.Get(index)
	if err != nil {
		return err
	}

	addNode := &Node{
		Val:  val,
		Next: nil,
		Prev: nil,
	}
	pre := indexNode.Prev

	pre.Next = addNode
	addNode.Prev = pre
	addNode.Next = indexNode
	indexNode.Prev = addNode

	l.Size++
	return  nil
}

func (l *Link)DeleteAtIndex(index int) error{
	if index < 0 || index > l.Size + 1{
		return errors.New("超出限制")
	}

	node, err := l.Get(index)

	if err != nil {
		return err
	}

	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev

	node.Next = nil
	node.Prev = nil

	l.Size--
	return  nil
}


/**
案例2：删除倒数第N个节点
 */

func (l *Link)DeleteRevertIndex(index int) error {
	if index == 0 || index > l.Size {
		return errors.New("index 不合法")
	}

	firstStart, err := l.Get(index)

	if err != nil {
		return fmt.Errorf("%w 先走队列出错", err)
	}

	firstNext := firstStart
	LastNext := 0
	for {
		firstNext = firstNext.Next
		LastNext++

		if firstNext == l.Tail {
			_ = l.DeleteAtIndex(LastNext)
			break
		}
	}
	return nil
}


func main()  {
	l := new(Link)

	l.Init()

	l.AddAtHead(0)
	l.AddAtHead(1)
	l.AddAtHead(4)

	l.AddAtTail(5)

	fmt.Println(l)

	tmp,_ := l.Get(4)
	fmt.Println(tmp)

	l.AddAtIndex(0, -10)

	fmt.Println(l, "=====")


	l.AddAtIndex(l.Size+1,10)

	fmt.Println(l, "=====")

	l.AddAtIndex(1,-100)

	fmt.Println(l, "-----加法结束-----")

/*	l.DeleteAtIndex(1)

	fmt.Println("减法开始", l)

	l.DeleteAtIndex(l.Size)

	fmt.Println(l)

	l.DeleteAtIndex(3)

	fmt.Println(l)*/


	l.DeleteRevertIndex(l.Size)

	fmt.Println(l, "-----revert end------")

	l.DeleteRevertIndex(1)

	fmt.Println(l, "-----revert 1------")

	l.DeleteRevertIndex(3)

	fmt.Println(l, "-----revert 3------")
}
