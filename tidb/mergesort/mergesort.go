package main

import (
	"math"
	"sort"
	"sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) (srcNew []int64) {
	//fmt.Printf("init Src: :%+v\n", src)

	/************分割src开始*************/
	// gcnt Goroutine数目
	// gCnt仅取偶数
	var gCnt int

	if len(src) < 100 {
		gCnt = 1
	} else {
		gCnt = 4
	}

	lenSrc := len(src)

	startIdx := 0
	endIdx := int(math.Min(float64(startIdx+lenSrc/gCnt), float64(lenSrc)))

	slices := make([][]int64, 0)

	// gcnt路分别使用sort.Slice()进行排序
	for i := 0; i <= gCnt; i+=1 {
		slices = append(slices, src[startIdx:endIdx])
		startIdx = endIdx
		endIdx = int(math.Min(float64(startIdx+lenSrc/gCnt), float64(lenSrc)))
	}

	//fmt.Printf("slices %+v\n", slices)
	/*************分割src结束************/



	//fmt.Printf("Before Sort: %#v\n", slices)


	/*************排序每个slice开始************/
	var wg sync.WaitGroup
	// sort every slice
	for i := 0; i < len(slices); i+=1 {
		wg.Add(1)
		go func(s *[]int64) {
			defer wg.Done()
			sort.Slice(*s, func(a, b int) bool {
							return (*s)[a] < (*s)[b]
			})
		}(&slices[i])
	}
	wg.Wait()
	/**************排序每个slice结束***********/


	//fmt.Printf("After Sort: %#v\n", slices)


	/**************归并slices开始***********/
	// 两两归并
	// slices[i] 和 slice[i+1]归并到slices[i]
	// 注意处理奇数情况
	// 先不并发归并

	for len(slices) > 1 {
		slicesNext := make([][]int64, 0)
		for i := 0; i < len(slices); i+=2 {
			// 长度为奇数 ，则将最后一个直接入新slice
			if len(slices) % 2 == 1 && i == len(slices)-1 {
				slicesNext = append(slicesNext, slices[i])
			} else {
				slicesNext = append(slicesNext, MergeSlices(slices[i], slices[i+1]))
			}
		}
		slices = slicesNext
	}
	srcNew = slices[0]
	return
}

func main() {
	src := make([]int64, 0)
	for i := 20; i > 0; i-- {
		src = append(src, int64(i))
	}

	MergeSort(src)
}

func MergeSlices(s1 []int64, s2 []int64) (ms []int64) {
	ms = make([]int64, 0)
	idx1 := 0
	idx2 := 0

	// 两个中有一个小于其长度
	for idx1 < len(s1) || idx2 < len(s2) {
		//idx1到结尾了 或 idx2 到结尾了
		if idx1 == len(s1) {
			ms = append(ms, s2[idx2])
			idx2 += 1
			continue
		} else if idx2 == len(s2) {
			ms = append(ms, s1[idx1])
			idx1 += 1
			continue
		} else {
			// 比较大小
			if s1[idx1] >= s2[idx2] {
				ms = append(ms, s2[idx2])
				idx2 += 1
			} else {
				ms = append(ms, s1[idx1])
				idx1 += 1
			}
		}
	}
	return ms
}