// 2296.设计一个文本编辑器
//请你设计一个带光标的文本编辑器，它可以实现以下功能：
//
//
// 添加：在光标所在处添加文本。
// 删除：在光标所在处删除文本（模拟键盘的删除键）。
// 移动：将光标往左或者往右移动。
//
//
// 当删除文本时，只有光标左边的字符会被删除。光标会留在文本内，也就是说任意时候 0 <= cursor.position <= currentText.
//length 都成立。
//
// 请你实现 TextEditor 类：
//
//
// TextEditor() 用空文本初始化对象。
// void addText(string text) 将 text 添加到光标所在位置。添加完后光标在 text 的右边。
// int deleteText(int k) 删除光标左边 k 个字符。返回实际删除的字符数目。
// string cursorLeft(int k) 将光标向左移动 k 次。返回移动后光标左边 min(10, len) 个字符，其中 len 是光标左边的
//字符数目。
// string cursorRight(int k) 将光标向右移动 k 次。返回移动后光标左边 min(10, len) 个字符，其中 len 是光标左边
//的字符数目。
//
//
//
//
// 示例 1：
//
//
//输入：
//["TextEditor", "addText", "deleteText", "addText", "cursorRight",
//"cursorLeft", "deleteText", "cursorLeft", "cursorRight"]
//[[], ["leetcode"], [4], ["practice"], [3], [8], [10], [2], [6]]
//输出：
//[null, null, 4, null, "etpractice", "leet", 4, "", "practi"]
//
//解释：
//TextEditor textEditor = new TextEditor(); // 当前 text 为 "|" 。（'|' 字符表示光标）
//textEditor.addText("leetcode"); // 当前文本为 "leetcode|" 。
//textEditor.deleteText(4); // 返回 4
//                          // 当前文本为 "leet|" 。
//                          // 删除了 4 个字符。
//textEditor.addText("practice"); // 当前文本为 "leetpractice|" 。
//textEditor.cursorRight(3); // 返回 "etpractice"
//                           // 当前文本为 "leetpractice|".
//                           // 光标无法移动到文本以外，所以无法移动。
//                           // "etpractice" 是光标左边的 10 个字符。
//textEditor.cursorLeft(8); // 返回 "leet"
//                          // 当前文本为 "leet|practice" 。
//                          // "leet" 是光标左边的 min(10, 4) = 4 个字符。
//textEditor.deleteText(10); // 返回 4
//                           // 当前文本为 "|practice" 。
//                           // 只有 4 个字符被删除了。
//textEditor.cursorLeft(2); // 返回 ""
//                          // 当前文本为 "|practice" 。
//                          // 光标无法移动到文本以外，所以无法移动。
//                          // "" 是光标左边的 min(10, 0) = 0 个字符。
//textEditor.cursorRight(6); // 返回 "practi"
//                           // 当前文本为 "practi|ce" 。
//                           // "practi" 是光标左边的 min(10, 6) = 6 个字符。
//
//
//
//
// 提示：
//
//
// 1 <= text.length, k <= 40
// text 只含有小写英文字母。
// 调用 addText ，deleteText ，cursorLeft 和 cursorRight 的 总 次数不超过 2 * 10⁴ 次。
//
//
//
//
// 进阶：你能设计并实现一个每次调用时间复杂度为 O(k) 的解决方案吗？
//
// Related Topics 栈 设计 链表 字符串 双向链表 模拟 👍 63 👎 0

package algorithm_2200

// leetcode submit region begin(Prohibit modification and deletion)
type TextEditor struct {
	Left, Right []byte
}

func ConstructorV2() TextEditor {
	return TextEditor{}
}

func (this *TextEditor) AddText(text string) {
	this.Left = append(this.Left, []byte(text)...)
}

func (this *TextEditor) DeleteText(k int) int {
	var num = len(this.Left)
	if num >= k {
		this.Left = this.Left[:num-k]
		return k
	}

	this.Left = this.Left[:0]
	return num
}

func (this *TextEditor) CursorLeft(k int) string {
	var i = len(this.Left) - 1
	for ; i >= 0 && k > 0; i, k = i-1, k-1 {
		this.Right = append(this.Right, this.Left[i])
	}
	this.Left = this.Left[:i+1]

	if len(this.Left) >= 10 {
		return string(this.Left[len(this.Left)-10:])
	}
	return string(this.Left)
}

func (this *TextEditor) CursorRight(k int) string {
	var i = len(this.Right) - 1
	for ; i >= 0 && k > 0; i, k = i-1, k-1 {
		this.Left = append(this.Left, this.Right[i])
	}
	this.Right = this.Right[:i+1] 

	if len(this.Left) >= 10 {
		return string(this.Left[len(this.Left)-10:])
	}
	return string(this.Left)
}

type TextEditorV2 struct {
	Val         rune
	Left, Right *TextEditorV2
}

func ConstructorV3() TextEditorV2 {
	return TextEditorV2{}
}

func (this *TextEditorV2) AddText(text string) {
	for _, v := range text {
		var te = &TextEditorV2{
			Val:   v,
			Left:  this.Left,
			Right: this,
		}
		if this.Left != nil {
			this.Left.Right = te
		}
		this.Left = te
	}
}

func (this *TextEditorV2) DeleteText(k int) int {
	var num int
	for ; k > 0 && this.Left != nil; k-- {
		if this.Left.Left != nil {
			this.Left.Left.Right = this
		}
		this.Left = this.Left.Left
		num++
	}
	return num
}

func (this *TextEditorV2) CursorLeft(k int) string {
	for ; k > 0 && this.Left != nil; k-- {
		if this.Right != nil {
			this.Right.Left = this.Left
		}
		this.Left.Right = this.Right

		this.Right, this.Left = this.Left, this.Left.Left
		this.Right.Left = this
		if this.Left != nil {
			this.Left.Right = this
		}
	}

	var tmp = this
	var chars = make([]rune, 10)
	var i = 9
	for ; i >= 0 && tmp.Left != nil; i-- {
		chars[i] = tmp.Left.Val
		tmp = tmp.Left
	}
	return string(chars[i+1:])
}

func (this *TextEditorV2) CursorRight(k int) string {
	for ; k > 0 && this.Right != nil; k-- {
		if this.Left != nil {
			this.Left.Right = this.Right
		}
		this.Right.Left = this.Left

		this.Right, this.Left = this.Right.Right, this.Right

		this.Left.Right = this
		if this.Right != nil {
			this.Right.Left = this
		}
	}

	var tmp = this
	var chars = make([]rune, 10)
	var i = 9
	for ; i >= 0 && tmp.Left != nil; i-- {
		chars[i] = tmp.Left.Val
		tmp = tmp.Left
	}
	return string(chars[i+1:])
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
//leetcode submit region end(Prohibit modification and deletion)
