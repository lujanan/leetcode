//请你设计一个用于存储字符串计数的数据结构，并能够返回计数最小和最大的字符串。
//
// 实现 AllOne 类：
//
//
// AllOne() 初始化数据结构的对象。
// inc(String key) 字符串 key 的计数增加 1 。如果数据结构中尚不存在 key ，那么插入计数为 1 的 key 。
// dec(String key) 字符串 key 的计数减少 1 。如果 key 的计数在减少后为 0 ，那么需要将这个 key 从数据结构中删除。测试用例
//保证：在减少计数前，key 存在于数据结构中。
// getMaxKey() 返回任意一个计数最大的字符串。如果没有元素存在，返回一个空字符串 "" 。
// getMinKey() 返回任意一个计数最小的字符串。如果没有元素存在，返回一个空字符串 "" 。
//
//
//
//
// 示例：
//
//
//输入
//["AllOne", "inc", "inc", "getMaxKey", "getMinKey", "inc", "getMaxKey",
//"getMinKey"]
//[[], ["hello"], ["hello"], [], [], ["leet"], [], []]
//输出
//[null, null, null, "hello", "hello", null, "hello", "leet"]
//
//解释
//AllOne allOne = new AllOne();
//allOne.inc("hello");
//allOne.inc("hello");
//allOne.getMaxKey(); // 返回 "hello"
//allOne.getMinKey(); // 返回 "hello"
//allOne.inc("leet");
//allOne.getMaxKey(); // 返回 "hello"
//allOne.getMinKey(); // 返回 "leet"
//
//
//
//
// 提示：
//
//
// 1 <= key.length <= 10
// key 由小写英文字母组成
// 测试用例保证：在每次调用 dec 时，数据结构中总存在 key
// 最多调用 inc、dec、getMaxKey 和 getMinKey 方法 5 * 10⁴ 次
//
// Related Topics 设计 哈希表 链表 双向链表 👍 203 👎 0

package algorithm_400

type AllOne struct {
	StrMap map[string]*OneLink
	Min    *OneLink
	Max    *OneLink
}

type OneLink struct {
	Val  string
	Num  int
	Pre  *OneLink
	Next *OneLink
}

func Constructor() AllOne {
	return AllOne{StrMap: make(map[string]*OneLink)}
}

func (this *AllOne) Inc(key string) {
	node, ok := this.StrMap[key]
	// 新key，直接放在最小位置
	if !ok {
		var node = &OneLink{Val: key, Num: 1}
		if this.Min == nil {
			this.Min, this.Max = node, node
		} else {
			this.Min.Pre, node.Next = node, this.Min
			this.Min = node
		}

		this.StrMap[key] = node
		return
	}

	node.Num++
	// 位置不变（右边没有元素 或 值比右边元素小）
	if !(node.Next != nil && node.Num > node.Next.Num) {
		return
	}

	// 原先是最小值，最小值变成右边元素
	if this.Min.Val == node.Val {
		this.Min = node.Next
	}

	// 左右两边元素连接
	next := node.Next
	next.Pre = node.Pre
	if node.Pre != nil {
		node.Pre.Next = next
	}
	// 当前元素单独出来
	node.Pre, node.Next = nil, nil
	// 比最大元素大，直接放到最大元素位置
	if node.Num >= this.Max.Num {
		this.Max.Next, node.Pre = node, this.Max
		this.Max = node
		return
	}

	// 找到第一个不小于的元素
	for node.Num > next.Num {
		next = next.Next
	}
	node.Pre, node.Next = next, next.Next
	if next.Next != nil {
		next.Next.Pre = node
	}
	next.Next = node
}

func (this *AllOne) Dec(key string) {
	node, ok := this.StrMap[key]
	if !ok {
		return
	}

	// 需要移除元素
	if node.Num <= 1 {
		delete(this.StrMap, key)
		if this.Min.Val == key {
			this.Min = this.Min.Next
			if this.Min != nil {
				this.Min.Pre = nil
			}
		}
		if this.Max.Val == key {
			this.Max = this.Max.Pre
			if this.Max != nil {
				this.Max.Next = nil
			}
		}

		if node.Pre != nil {
			node.Pre.Next = node.Next
		}
		if node.Next != nil {
			node.Next.Pre = node.Pre
		}
		return
	}

	node.Num--
	// // 位置不变（左边没有元素 或 值比左边元素大）
	if !(node.Pre != nil && node.Num < node.Pre.Num) {
		return
	}
	// 原先是最大值，最大值变成左边元素
	if this.Max.Val == node.Val {
		this.Max = node.Pre
	}

	// 左右两边元素连接
	pre := node.Pre
	pre.Next = node.Next
	if node.Next != nil {
		node.Next.Pre = pre
	}
	// 当前元素单独出来
	node.Pre, node.Next = nil, nil
	// 比最小元素小，直接放到最小元素位置
	if node.Num <= this.Min.Num {
		this.Min.Pre, node.Next = node, this.Min
		this.Min = node
		return
	}

	// 找到第一个不大于的元素
	for node.Num < pre.Num {
		pre = pre.Pre
	}
	node.Pre, node.Next = pre.Pre, pre
	if pre.Pre != nil {
		pre.Pre.Next = node
	}
	pre.Pre = node
}

func (this *AllOne) GetMaxKey() string {
	if this.Max != nil {
		return this.Max.Val
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.Min != nil {
		return this.Min.Val
	}
	return ""
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
