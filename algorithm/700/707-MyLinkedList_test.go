//设计链表的实现。您可以选择使用单链表或双链表。单链表中的节点应该具有两个属性：val 和 next。val 是当前节点的值，next 是指向下一个节点的指针
///引用。如果要使用双向链表，则还需要一个属性 prev 以指示链表中的上一个节点。假设链表中的所有节点都是 0-index 的。
//
// 在链表类中实现这些功能：
//
//
// get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
// addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
// addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
// addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val 的节点。如果 index 等于链表的长度，则该节点将附加
//到链表的末尾。如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
// deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。
//
//
//
//
// 示例：
//
// MyLinkedList linkedList = new MyLinkedList();
//linkedList.addAtHead(1);
//linkedList.addAtTail(3);
//linkedList.addAtIndex(1,2);   //链表变为1-> 2-> 3
//linkedList.get(1);            //返回2
//linkedList.deleteAtIndex(1);  //现在链表是1-> 3
//linkedList.get(1);            //返回3
//
//
//
//
// 提示：
//
//
// 所有val值都在 [1, 1000] 之内。
// 操作次数将在 [1, 1000] 之内。
// 请不要使用内置的 LinkedList 库。
//
//
// Related Topics 设计 链表 👍 609 👎 0

package algorithm_700

import (
	"fmt"
	"strings"
	"testing"
)

func TestConstructor(t *testing.T) {
	obj := Constructor()
	param_1 := obj.Get(0)
	fmt.Println(param_1)
	obj.AddAtHead(1)
	printlnLink(&obj)
	obj.AddAtTail(3)
	printlnLink(&obj)
	obj.AddAtIndex(1, 2)
	printlnLink(&obj)
	obj.DeleteAtIndex(1)
	printlnLink(&obj)

	return

	obj2 := Constructor()
	obj2.AddAtHead(7)
	obj2.AddAtHead(2)
	obj2.AddAtHead(1)
	printlnLink(&obj2)

	obj2.AddAtIndex(3, 0)
	printlnLink(&obj2)
	obj2.DeleteAtIndex(2)
	printlnLink(&obj2)

	obj2.AddAtHead(6)
	obj2.AddAtTail(4)
	printlnLink(&obj2)

	fmt.Println(obj2.Get(4))
	obj2.AddAtHead(4)
	printlnLink(&obj2)
	obj2.AddAtIndex(5, 0)
	printlnLink(&obj2)
	obj2.AddAtHead(6)
	printlnLink(&obj2)

	return

	var egFunc = []string{
		"MyLinkedList", "addAtHead", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "addAtTail", "deleteAtIndex",
		"addAtHead", "addAtHead", "get", "addAtTail", "addAtHead", "get", "addAtTail", "addAtIndex", "addAtTail", "addAtHead", "addAtHead",
		"addAtHead", "get", "addAtIndex", "addAtHead", "get", "addAtHead", "deleteAtIndex", "addAtHead", "addAtTail", "addAtTail", "addAtIndex",
		"addAtTail", "addAtHead", "get", "addAtTail", "deleteAtIndex", "addAtIndex", "deleteAtIndex", "addAtHead", "addAtTail", "addAtHead",
		"addAtHead", "addAtTail", "addAtTail", "get", "get", "addAtHead", "addAtTail", "addAtTail", "addAtTail", "addAtIndex", "get", "addAtHead",
		"addAtIndex", "addAtHead", "addAtTail", "addAtTail", "addAtIndex", "deleteAtIndex", "addAtIndex", "addAtHead", "addAtHead",
		"deleteAtIndex", "addAtTail", "deleteAtIndex", "addAtIndex", "addAtTail", "addAtHead", "get", "addAtIndex", "addAtTail", "addAtHead",
		"addAtHead", "addAtHead", "addAtHead", "addAtHead", "addAtHead", "deleteAtIndex", "get", "get", "addAtHead", "get", "addAtTail",
		"addAtTail", "addAtIndex", "addAtIndex", "addAtHead", "addAtTail", "addAtTail", "get", "addAtIndex", "addAtHead", "deleteAtIndex",
		"addAtTail", "get", "addAtHead", "get", "addAtHead", "deleteAtIndex", "get", "addAtTail", "addAtTail",
	}

	var egParam = [][]int{
		{}, {38}, {66}, {61}, {76}, {26}, {37}, {8}, {5}, {4}, {45}, {4}, {85}, {37}, {5}, {93}, {10, 23}, {21}, {52}, {15}, {47}, {12}, {6, 24}, {64}, {4},
		{31}, {6}, {40}, {17}, {15}, {19, 2}, {11}, {86}, {17}, {55}, {15}, {14, 95}, {22}, {66}, {95}, {8}, {47}, {23}, {39}, {30}, {27}, {0}, {99}, {45},
		{4}, {9, 11}, {6}, {81}, {18, 32}, {20}, {13}, {42}, {37, 91}, {36}, {10, 37}, {96}, {57}, {20}, {89}, {18}, {41, 5}, {23}, {75}, {7}, {25, 51}, {48},
		{46}, {29}, {85}, {82}, {6}, {38}, {14}, {1}, {12}, {42}, {42}, {83}, {13}, {14, 20}, {17, 34}, {36}, {58}, {2}, {38}, {33, 59}, {37}, {15}, {64},
		{56}, {0}, {40}, {92}, {63}, {35}, {62}, {32},
	}

	var obj1 MyLinkedList
	for idx, fn := range egFunc {
		switch fn {
		case "MyLinkedList":
			obj1 = Constructor()
		case "addAtHead":
			obj1.AddAtHead(egParam[idx][0])

		case "addAtTail":
			obj1.AddAtTail(egParam[idx][0])

		case "addAtIndex":
			obj1.AddAtIndex(egParam[idx][0], egParam[idx][1])

		case "deleteAtIndex":
			obj1.DeleteAtIndex(egParam[idx][0])

		case "get":
			printlnLink(&obj1)
			fmt.Println(fn, egParam[idx][0])
			fmt.Println(obj1.Get(egParam[idx][0]))
		}
	}
}

func printlnLink(obj *MyLinkedList) {
	var nums []string
	var tmp = obj.head
	for tmp != nil {
		nums = append(nums, fmt.Sprintf("%d", tmp.val))
		tmp = tmp.next
	}
	fmt.Println(strings.Join(nums, "->"))
}
