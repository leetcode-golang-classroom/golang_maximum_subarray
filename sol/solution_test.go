package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for idx := 0; idx < b.N; idx++ {
		maxSubArray(nums)
	}
}
func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "nums = [-2,1,-3,4,-1,2,1,-5,4]",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			name: "nums = [1]",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "nums = [5,4,-1,7,8]",
			args: args{[]int{5, 4, -1, 7, 8}},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
