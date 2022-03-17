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

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	var maxKey, minKey string
	allOne := Constructor()
	allOne.Inc("hello")
	allOne.Inc("world")
	allOne.Inc("leet")
	allOne.Inc("code")
	allOne.Inc("ds")
	allOne.Inc("leet")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	
	allOne.Inc("ds")
	allOne.Dec("leet")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)


	allOne.Dec("ds")
	allOne.Inc("hello")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	
	allOne.Inc("hello")
	allOne.Inc("hello")
	allOne.Dec("world")
	allOne.Dec("leet")
	allOne.Dec("code")
	allOne.Dec("ds")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	
	allOne.Inc("new")
	allOne.Inc("new")
	allOne.Inc("new")
	allOne.Inc("new")
	allOne.Inc("new")
	allOne.Inc("new")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	minKey = allOne.GetMinKey()
	fmt.Println("minKey", minKey)

	return




	allOne.Inc("hello")
	allOne.Inc("hello")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	minKey = allOne.GetMinKey()
	fmt.Println("minKey", minKey)

	allOne.Inc("leet")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	minKey = allOne.GetMinKey()
	fmt.Println("minKey", minKey)

	// allOne.Dec("hello")
	allOne.Dec("hello")
	allOne.Dec("leet")
	allOne.Inc("leet")
	allOne.Inc("leet")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	minKey = allOne.GetMinKey()
	fmt.Println("minKey", minKey)

	allOne.Inc("world")
	allOne.Inc("world")
	allOne.Inc("world")
	allOne.Inc("world")
	maxKey = allOne.GetMaxKey()
	fmt.Println("maxKey", maxKey)
	minKey = allOne.GetMinKey()
	fmt.Println("minKey", minKey)
}
