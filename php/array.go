package php

import (
	"math/rand"
	"reflect"
	"time"
)

// Array - Create an array
func Array(v ...interface{}) []interface{} {

	return v
}

// Count - Count all elements in an array, or something in an object
func Count(v []interface{}) int {

	return len(v)
}

// ArrayColumn — Return the values from a single column in the input array
func ArrayColumn(arrayMap map[string]map[string]interface{}, columnKey string) (r []interface{}) {

	for _, input := range arrayMap {
		if v, ok := input[columnKey]; ok {
			r = append(r, v)
		}
	}

	return
}

func ArrayCountValues(s []interface{}) map[interface{}]uint {

	r := make(map[interface{}]uint)
	for _, v := range s {
		if c, ok := r[v]; ok {
			r[v] = c + 1
		} else {
			r[v] = 1
		}
	}

	return r
}

// ArrayFill - Fill an array with values
func ArrayFill(startIndex uint, num uint, value interface{}) map[uint]interface{} {

	r := make(map[uint]interface{})
	var i uint
	for i = 0; i < num; i++ {
		r[startIndex] = value
		startIndex++
	}

	return r
}

// ArrayFillKeys - Fill an array with values, specifying keys
func ArrayFillKeys(keys []interface{}, value interface{}) map[interface{}]interface{} {

	r := make(map[interface{}]interface{})
	for _, v := range keys {
		r[v] = value
	}

	return r
}

//ArrayDiff 模拟PHP array_diff函数
func ArrayDiff(array1 []interface{}, othersParams ...[]interface{}) ([]interface{}, error) {
	if len(array1) == 0 {
		return []interface{}{}, nil
	}
	if len(array1) > 0 && len(othersParams) == 0 {
		return array1, nil
	}
	var tmp = make(map[interface{}]int, len(array1))
	for _, v := range array1 {
		tmp[v] = 1
	}
	for _, param := range othersParams {
		for _, arg := range param {
			if tmp[arg] != 0 {
				tmp[arg]++
			}
		}
	}
	var res = make([]interface{}, 0, len(tmp))
	for k, v := range tmp {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res, nil
}

//ArrayIntersect 模拟PHP array_intersect函数
func ArrayIntersect(array1 []interface{}, othersParams ...[]interface{}) ([]interface{}, error) {
	if len(array1) == 0 {
		return []interface{}{}, nil
	}
	if len(array1) > 0 && len(othersParams) == 0 {
		return array1, nil
	}
	var tmp = make(map[interface{}]int, len(array1))
	for _, v := range array1 {
		tmp[v] = 1
	}
	for _, param := range othersParams {
		for _, arg := range param {
			if tmp[arg] != 0 {
				tmp[arg]++
			}
		}
	}
	var res = make([]interface{}, 0, len(tmp))
	for k, v := range tmp {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res, nil
}

func ArrayMerge(arr ...[]interface{}) []interface{} {

	s := make([]interface{}, 0)
	for _, v := range arr {
		s = append(s, v...)
	}

	return s
}

// ArrayPush - Push one or more elements onto the end of array
func ArrayPush(s *[]string, args ...string) {

	*s = append(*s, args...)
}

// ArrayPop array_pop()
// Pop the element off the end of slice
func ArrayPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// ArrayUnshift array_unshift()
// Prepend one or more elements to the beginning of a slice
func ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
	*s = append(elements, *s...)
	return len(*s)
}

// ArrayShift array_shift()
// Shift an element off the beginning of slice
func ArrayShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}

// ArrayKeyExists array_key_exists()
func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[key]
	return ok
}

// ArrayCombine array_combine()
func ArrayCombine(s1, s2 []interface{}) map[interface{}]interface{} {
	if len(s1) != len(s2) {
		panic("the number of elements for each slice isn't equal")
	}
	m := make(map[interface{}]interface{}, len(s1))
	for i, v := range s1 {
		m[v] = s2[i]
	}
	return m
}

// ArrayKeys - get keys of map data as a Array
// in php,the keys you want always is string or number
// here,let it be string
func ArrayKeys(data map[string]interface{}) []string {
	if len(data) < 1 {
		return []string{}
	}
	var resData []string
	for index := range data {
		resData = append(resData, index)
	}
	return resData
}

// ArrayValues array_values()
func ArrayValues(elements map[interface{}]interface{}) []interface{} {
	i, vals := 0, make([]interface{}, len(elements))
	for _, val := range elements {
		vals[i] = val
		i++
	}
	return vals
}

// ArrayReverse - Return an array with elements in reverse order
func ArrayReverse(s []interface{}) []interface{} {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// InArray in_array()
// haystack supported types: slice, array or map
func InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
}

// ArrayFlip - Exchanges all keys with their associated values in an array
func ArrayFlip(arrayMap map[interface{}]interface{}) map[interface{}]interface{} {

	r := make(map[interface{}]interface{})
	for i, v := range arrayMap {
		r[v] = i
	}

	return r
}

// ArraySlice array_slice()
func ArraySlice(s []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(s)) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < uint(len(s)) {
		return s[offset:end]
	}
	return s[offset:]
}

// ArrayPad array_pad()
func ArrayPad(s []interface{}, size int, val interface{}) []interface{} {
	if size == 0 || (size > 0 && size < len(s)) || (size < 0 && size > -len(s)) {
		return s
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(s)
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = val
	}
	if size > 0 {
		return append(s, tmp...)
	}
	return append(tmp, s...)
}

// ArrayRand array_rand()
func ArrayRand(elements []interface{}) []interface{} {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]interface{}, len(elements))
	for i, v := range r.Perm(len(elements)) {
		n[i] = elements[v]
	}
	return n
}