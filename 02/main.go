package main

import (
	"fmt"
	"sort"
)

/**
 找出 数组中 和 为target 的组合
 如果顺序有关，直接把所有的固定顺序 * 2 即可
 感觉没必要做重复的数据
 */
func main()  {
	arr := []int{1,4,3,2,8,100,50}
	target := 10

	r := Sum2(arr, target)

	r1 := Sum1(arr, target)

	r2 := Sum3(arr, target)

	fmt.Println(r, r1, r2)
}

/**
暴力， 两次循环
 */
func Sum1(arr []int, target int) [][]int {
	result := make([][]int, 0)

	for i:=0; i< len(arr) ; i++ {
		for j:= i+1; j < len(arr);  j++ {
			if arr[i] + arr[j] == target {
				result = append(result, []int{arr[i], arr[j]})
			}
		}
	}

	return result
}

/**
利用map， 直接找出对应的target
会出现正序和 fan序列，可以利用 map 中存的值过滤掉另一半
 */
func Sum2(arr []int, target int) [][]int  {
	result := make([][]int, 0)
	// 先搞个 map
	m := make(map[int]struct{})
	for _, v := range arr {
		m[v] = struct{}{}
	}
	for _, v := range arr {
		diff := target - v

		if _, ok := m[diff]; ok {
			result = append(result, []int{v, diff})
		}
	}
	return  result
}

/**
 先排序，再利用双指针去做
 */
func Sum3(arr []int, target int) [][]int {
	// 大于 2 元素 才会这么做的

	sort.SliceStable(arr, func(i, j int) bool {
		if arr[i] < arr[j] {
			return true
		} else {
			return false
		}
	})


	min := 0
	max := len(arr) - 1

	result := make([][]int, 0)

	for {
		if min == max {
			break
		}

		// 值太大了
		if arr[min] + arr[max] > target {
			max--
		} else if arr[min] + arr[max] < target {
			min++
		} else {
			result = append(result, []int{arr[min], arr[max]})
			min++
		}

	}
	return result
}
