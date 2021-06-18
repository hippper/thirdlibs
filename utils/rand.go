package utils

import (
	"math/rand"
	"time"
)

var Rander = rand.New(rand.NewSource(time.Now().UnixNano())) // 线程不安全

// 包含min, max
func RandNum(min, max int) int {
	return Rander.Intn(max-min+1) + min
}

/** Returns a random shuffling of the array. */
func Shuffle(src []string) []string {
	nums := make([]string, len(src))
	copy(nums, src)

	Rander.Shuffle(len(nums), func(i int, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}
