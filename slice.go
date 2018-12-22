package go_libs_lawtech

import (
	"math/rand"
	"time"
)

/*
useful functions for slice
*/

type reducetype func(interface{}) interface{}
type filtertype func(interface{}) bool

func SliceRandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index, _ := range list {
		list[index] += min
	}
	return list
}

func SliceMerge(slice1, slice2 []interface{}) []interface{} {
	return append(slice1, slice2...)
}

func InSlice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func SliceReduce(slice []interface{}, a reducetype) (dSclie []interface{}) {
	for _, v := range slice {
		dSclie = append(dSclie, a(v))
	}
	return
}

func SliceRand(a []interface{}) interface{} {
	randNum := rand.Intn(len(a))
	return a[randNum]
}

func SliceSum(intSlice []int64) (sum int64) {
	for _, v := range intSlice {
		sum += v
	}
	return
}

func SliceFilter(slice []interface{}, a filtertype) (ftSlice []interface{}) {
	for _, v := range slice {
		if a(v) {
			ftSlice = append(ftSlice, v)
		}
	}
	return
}

func SliceDiff(slice1, slice2 []interface{}) (diffSlice []interface{}) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffSlice = append(diffSlice, v)
		}
	}
	return
}

func SliceChunk(slice []interface{}, size int) (chunkSlice [][]interface{}) {
	if size >= len(slice) {
		chunkSlice = append(chunkSlice, slice)
		return
	}
	end := size
	for i := 0; i <= len(slice)-size; i += size {
		chunkSlice = append(chunkSlice, slice[i:end])
		end += size
	}
	return
}

func SliceRange(start, end, step int64) (intSlice []int64) {
	for i := start; i <= end; i += step {
		intSlice = append(intSlice, i)
	}
	return
}

func SlicePad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}

	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}

	return slice
}

func SliceUnique(slice []interface{}) (uniqueSlice []interface{}) {
	for _, v := range slice {
		if !InSlice(v, uniqueSlice) {
			uniqueSlice = append(uniqueSlice, v)
		}
	}
	return
}

func SliceShuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}
