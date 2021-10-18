//给你一个整数数组 nums ，你可以在 nums 上执行下述操作 任意次 ：
//
//
// 如果 gcd(nums[i], nums[j]) > 1 ，交换 nums[i] 和 nums[j] 的位置。其中 gcd(nums[i], nums[
//j]) 是 nums[i] 和 nums[j] 的最大公因数。
//
//
// 如果能使用上述交换方式将 nums 按 非递减顺序 排列，返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
// 输入：nums = [7,21,3]
//输出：true
//解释：可以执行下述操作完成对 [7,21,3] 的排序：
//- 交换 7 和 21 因为 gcd(7,21) = 7 。nums = [21,7,3]
//- 交换 21 和 3 因为 gcd(21,3) = 3 。nums = [3,7,21]
//
//
// 示例 2：
//
// 输入：nums = [5,2,6,2]
//输出：false
//解释：无法完成排序，因为 5 不能与其他元素交换。
//
//
// 示例 3：
//
// 输入：nums = [10,5,9,3,15]
//输出：true
//解释：
//可以执行下述操作完成对 [10,5,9,3,15] 的排序：
//- 交换 10 和 15 因为 gcd(10,15) = 5 。nums = [15,5,9,3,10]
//- 交换 15 和 3 因为 gcd(15,3) = 3 。nums = [3,5,9,15,10]
//- 交换 10 和 15 因为 gcd(10,15) = 5 。nums = [3,5,9,10,15]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 3 * 10⁴
// 2 <= nums[i] <= 10⁵
//
// Related Topics 并查集 数组 数学 排序 👍 19 👎 0

package algorithm_1900

import "testing"

func Test_gcdSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes bool
	}{
		{"t1", args{[]int{7, 21, 3}}, true},
		{"t2", args{[]int{5, 2, 6, 2}}, false},
		{"t3", args{[]int{10, 5, 9, 3, 15}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := gcdSort(tt.args.nums); gotRes != tt.wantRes {
				t.Errorf("gcdSort() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
