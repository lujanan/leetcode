// 2241. 设计一个 ATM 机器

package algorithm_2200

import "math"

var moneyVal = []int{20, 50, 100, 200, 500}

type ATM struct {
	money []int
}

func Constructor() ATM {
	return ATM{make([]int, 5)}
}

func (this ATM) Deposit(banknotesCount []int) {
	for i := 0; i < len(banknotesCount); i++ {
		this.money[i] += banknotesCount[i]
	}
}

func (this ATM) Withdraw(amount int) []int {
	var res = make([]int, 5)
	for i := 4; i >= 0 && amount > 0; i-- {
		if this.money[i] <= 0 {
			continue
		}

		res[i] = int(math.Min(float64(this.money[i]), float64(amount/moneyVal[i])))
		amount -= res[i] * moneyVal[i]
	}

	if amount > 0 {
		return []int{-1}
	}
	for i := 0; i < len(res); i++ {
		this.money[i] -= res[i]
	}
	return res
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
