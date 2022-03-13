package algorithm_competition

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	var del bool
	l, r := 0, len(list)-1
	for l <= r {
		if list[l] != list[r] {
			if del {
				return false
			}
			if l+1 <= r && list[l+1] == list[r] {
				l++
				del = true
			} else if r-1 >= l && list[l] == list[r-1] {
				r--
				del = true
			} else {
				return false
			}
		}
		l++
		r--
	}

	return true
}

//
func maxInvestment(product []int, limit int) int {
	var pMap = make(map[int]int)
	var maxNum int
	for i := range product {
		pMap[product[i]]++
		if maxNum < product[i] {
			maxNum = product[i]
		}
	}

	var res int64
	for limit > 0 && maxNum > 0 {
		if pMap[maxNum] > limit {
			res += int64(limit * maxNum)
			limit = 0
		} else {
			res += int64(pMap[maxNum] * maxNum)
			limit -= pMap[maxNum]
			maxNum--
			pMap[maxNum] += pMap[maxNum+1]
		}
	}

	return int(res % (1000000007))
}

func coopDevelop(skills [][]int) int {
	var res int64
	var sMap = make(map[[4]int]int)

	for k, v := range skills {
		var skill = [4]int{}
		for i := range v {
			skill[v[i]-1] = v[i]
		}
		sMap[skill]++

		if k+1 > sMap[skill] {
			res += int64(k + 1 - sMap[skill])
		}
	}
	return int(res % 1000000007)
}

type DiscountSystem struct {
	Config []ActCfg
}

type ActCfg struct {
	ActId      int
	PriceLimit int
	Discount   int
	Number     int
	UserLimit  int
	UserCost   map[int]int
}

func Constructor() DiscountSystem {
	return DiscountSystem{}
}

func (this *DiscountSystem) AddActivity(actId int, priceLimit int, discount int, number int, userLimit int) {
	cfg := ActCfg{
		ActId:      actId,
		PriceLimit: priceLimit,
		Discount:   discount,
		Number:     number,
		UserLimit:  userLimit,
		UserCost:   make(map[int]int),
	}
	for i := range this.Config {
		if this.Config[i].ActId == actId {
			this.Config[i] = cfg
			return
		}
	}

	this.Config = append(this.Config, cfg)
	sort.Slice(this.Config, func(i, j int) bool {
		return this.Config[i].ActId < this.Config[j].ActId
	})
}

func (this *DiscountSystem) RemoveActivity(actId int) {
	var idx = -1
	for i := range this.Config {
		if this.Config[i].ActId == actId {
			idx = i
			break
		}
	}
	var list []ActCfg
	if idx >= 0 && idx < len(this.Config) {
		list = this.Config[:idx]
		list = append(list, this.Config[idx+1:]...)
	}
	this.Config = list
}

func (this *DiscountSystem) Consume(userId int, cost int) int {
	var costMax, aid int
	for _, cfg := range this.Config {
		if cost >= cfg.PriceLimit &&
			cfg.Number > 0 &&
			costMax < cfg.Discount {
			if _, ok := cfg.UserCost[userId]; !ok || cfg.UserCost[userId] < cfg.UserLimit {
				costMax, aid = cfg.Discount, cfg.ActId
			}
		}
	}

	if costMax > 0 {
		for k, cfg := range this.Config {
			if cfg.ActId == aid {
				this.Config[k].Number--
				this.Config[k].UserCost[userId]++
			}
		}
	}

	return cost - costMax
}

/**
 * Your DiscountSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddActivity(actId,priceLimit,discount,number,userLimit);
 * obj.RemoveActivity(actId);
 * param_3 := obj.Consume(userId,cost);
 */
