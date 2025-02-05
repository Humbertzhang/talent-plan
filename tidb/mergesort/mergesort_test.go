package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/pingcap/check"
)

var _ = check.Suite(&sortTestSuite{})

func TestT(t *testing.T) {
	check.TestingT(t)
}

func prepare(src []int64) {
	rand.Seed(time.Now().Unix())
	for i := range src {
		src[i] = rand.Int63()
	}
}

type sortTestSuite struct{}

func (s *sortTestSuite) TestMergeSort(c *check.C) {
	lens := []int{1, 3, 5, 7, 11, 13, 17, 19, 23, 29, 1024, 1 << 13, 1 << 17, 1 << 19, 1 << 20}


	for i := range lens {
		//fmt.Printf("Lens:%d\n", lens[i])
		src := make([]int64, lens[i])
		expect := make([]int64, lens[i])
		prepare(src)
		copy(expect, src)
		src = MergeSort(src)
		sort.Slice(expect, func(i, j int) bool { return expect[i] < expect[j] })

		//fmt.Printf("RealSrc: %+v\n", expect)
		//fmt.Printf("Mysrc:%+v\n", src)
		for i := 0; i < len(src); i++ {
			c.Assert(src[i], check.Equals, expect[i])
		}
	}

}
