package algorithm_competition

import (
	"testing"
)

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{[]int{18, 19, 5, 5, 18, 19, 5, 6, 12, 19, 13, 4, 16, 11, 4, 16, 10, 8, 12, 8, 2, 1, 8, 17, 4, 18, 3, 5, 16, 2, 16, 12, 17, 16, 7, 16, 2, 17, 19, 9, 1, 20, 17, 17, 4, 6}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArray(tt.args.nums); got != tt.want {
				t.Errorf("divideArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_halveArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{[]int{5, 19, 8, 1}}, 3},
		{"t2", args{[]int{3, 8, 20}}, 3},
		{"t3", args{[]int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halveArray(tt.args.nums); got != tt.want {
				t.Errorf("halveArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumSubsequenceCount(t *testing.T) {
	type args struct {
		text    string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"t1", args{"abdcdbc", "ac"}, 4},
		{"t2", args{"aabb", "ab"}, 6},
		{"t3", args{"vnedkpkkyxelxqptfwuzcjhqmwagvrglkeivowvbjdoyydnjrqrqejoyptzoklaxcjxbrrfmpdxckfjzahparhpanwqfjrpbslsyiwbldnpjqishlsuagevjmiyktgofvnyncizswldwnngnkifmaxbmospdeslxirofgqouaapfgltgqxdhurxljcepdpndqqgfwkfiqrwuwxfamciyweehktaegynfumwnhrgrhcluenpnoieqdivznrjljcotysnlylyswvdlkgsvrotavnkifwmnvgagjykxgwaimavqsxuitknmbxppgzfwtjdvegapcplreokicxcsbdrsyfpustpxxssnouifkypwqrywprjlyddrggkcglbgcrbihgpxxosmejchmzkydhquevpschkpyulqxgduqkqgwnsowxrmgqbmltrltzqmmpjilpfxocflpkwithsjlljxdygfvstvwqsyxlkknmgpppupgjvfgmxnwmvrfuwcrsadomyddazlonjyjdeswwznkaeaasyvurpgyvjsiltiykwquesfjmuswjlrphsdthmuqkrhynmqnfqdlwnwesdmiiqvcpingbcgcsvqmsmskesrajqwmgtdoktreqssutpudfykriqhblntfabspbeddpdkownehqszbmddizdgtqmobirwbopmoqzwydnpqnvkwadajbecmajilzkfwjnpfyamudpppuxhlcngkign", "rr"}, 496},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSubsequenceCount(tt.args.text, tt.args.pattern); got != tt.want {
				t.Errorf("maximumSubsequenceCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
