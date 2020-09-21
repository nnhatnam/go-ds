package binaryheap

import "fmt"

func Example()  {
	maxHeap := NewHeap(15)
	maxHeap.insert(5)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(3)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(17)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(10)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(84)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(19)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(6)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(22)
	fmt.Println(maxHeap.Data)
	maxHeap.insert(9)
	fmt.Println(maxHeap.Data)
	maxHeap.Print()
	fmt.Println(maxHeap.Data)
}
