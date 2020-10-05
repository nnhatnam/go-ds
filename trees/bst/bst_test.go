package bst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func treeInorder(tree *Tree) []Int {
	tempArr := []Int{}

	tree.InOderTraverse(func(v Interface){
		tempArr = append(tempArr, v.(Int))
	})
	return tempArr
}

//http://www.mathcs.emory.edu/~cheung/Courses/171/Syllabus/9-BinTree/bin-tree.html
func TestBasic(t *testing.T) {

	initalArr := []Int{50 , 30, 20 , 40 , 70 , 60 , 80}
	tree := NewTree()
	for _ , v := range initalArr {
		tree.Insert(v)
	}
	tempArr := treeInorder(tree)

	assert.Equal(t, []Int{20 , 30 , 40 , 50 , 60 , 70 , 80 } , tempArr )
	assert.Equal(t, tree.size, 7)
	tree.Remove(Int(20))
	tempArr = treeInorder(tree)

	assert.Equal(t, []Int{30 , 40 , 50 , 60 , 70 , 80 } , tempArr )
	assert.Equal(t, tree.size, 6)

	tree.Remove(Int(30))
	tempArr = treeInorder(tree)

	assert.Equal(t,  []Int{40 , 50 , 60 , 70 , 80 }, tempArr)
	assert.Equal(t, tree.size, 5)

	tree.Remove(Int(50))
	tempArr = treeInorder(tree)

	assert.Equal(t, []Int{40 , 60 , 70 , 80 } , tempArr )
	assert.Equal(t, tree.size, 4)

	fmt.Println(tempArr) //[20 30 40 50 60 70 80]
}

//https://www.cs.usfca.edu/~galles/visualization/BST.html
func TestTree_Remove(t *testing.T) {

	initalArr := []Int{50 , 30, 70 , 10 , 5, 7 , 40 , 39 , 38 , 45 , 80, 90 , 75}
	tree := NewTree()
	for _ , v := range initalArr {
		tree.Insert(v)
	}
	tempArr := treeInorder(tree)

	assert.Equal(t, []Int{5, 7, 10, 30, 38, 39, 40, 45, 50, 70, 75, 80, 90 } , tempArr )
	assert.Equal(t, tree.size, 13)

	tree.Remove(Int(7))
	tempArr = treeInorder(tree)

	assert.Equal(t, []Int{5, 10, 30, 38, 39, 40, 45, 50, 70, 75, 80, 90 } , tempArr )
	assert.Equal(t, tree.size, 12)

	tree.Remove(Int(10))
	tempArr = treeInorder(tree)

	assert.Equal(t,  []Int{5, 30, 38, 39, 40, 45, 50, 70, 75, 80, 90 }, tempArr)
	assert.Equal(t, tree.size, 11)

	tree.Remove(Int(70))
	tempArr = treeInorder(tree)

	assert.Equal(t, []Int{5, 30, 38, 39, 40, 45, 50, 75, 80, 90 } , tempArr )
	assert.Equal(t, tree.size, 10)

	search(tree.Root(), Int(50))

	fmt.Println(tempArr) //[20 30 40 50 60 70 80]
}

func Test1Node(t *testing.T) {

	initalArr := []Int{30}
	tree := NewTree()
	for _ , v := range initalArr {
		tree.Insert(v)
	}
	tempArr := treeInorder(tree)
	tree.Remove(Int(30))
	tempArr = treeInorder(tree)

	assert.Equal(t, []Int{ } , tempArr )
	assert.Equal(t, tree.size, 0)

	fmt.Println(tempArr) //[20 30 40 50 60 70 80]
}
