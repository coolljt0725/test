package main

import (
	"fmt"
)

func rec_sub(x []byte) {
	if len(x) < 1 {
		return
	}
	i := len(x) - 1
	if x[i] > 0 {
		x[i] = x[i] - 1
		return
	} else {
		x[i] = 9
		rec_sub(x[:i])
	}
}

func compare(a []byte, b []byte) int {
	var j int
	for j = range a {
		if a[j] > 0 {
			break
		}
	}
	sa := a[j:]
	la := len(sa)
	lb := len(b)
	if la < lb {
		return -1
	}
	var i int
	if la == lb {
		for i = 0; i < la; i++ {
			if sa[i] < b[i] {
				return -1
			} else if sa[i] > b[i] {
				break
			}
		}
		if i == la {
			return 0
		}
	}
	return 1
}

func sub(a []byte, b []byte) int {
	ans := compare(a, b)
	if ans < 1 {
		return ans
	}
	la := len(a)
	lb := len(b)
	res := 0

	for compare(a, b) >= 0 {
		for i := lb - 1; i >= 0; i-- {
			if a[la-lb+i] < b[i] {
				a[la-lb+i] += 10
				rec_sub(a[:la-lb+i])
			}
			a[la-lb+i] = a[la-lb+i] - b[i]
		}
		res++
	}

	return res
}

func main() {
	var (
		str1 string
		str2 string
	)
	fmt.Scanf("%s", &str1)
	fmt.Scanf("%s", &str2)
	a := []byte(str1)
	b := []byte(str2)
	c := make([]byte, 100)

	for i := range a {
		a[i] = a[i] - byte('0')
	}

	for i := range b {
		b[i] = b[i] - byte('0')
	}

	n := len(str1) - len(str2)
	if n < 0 {
		fmt.Printf("0\n")
		return
	}
	for i := 0; i < n; i++ {
		b = append(b, 0)
	}
	var i int
	for i = 0; i <= n; i++ {
		ans := sub(a, b[:len(b)-i])
		if ans > 0 {
			c[i] = (byte)(ans) + (byte)('0')
		} else {
			c[i] = (byte)('0')
		}
	}
	var k int
	for k = 0; k < i; k++ {
		if c[k] > (byte)('0') {
			break
		}
	}
	if k == i {
		fmt.Printf("0\n")
	} else {
		fmt.Printf("%s\n", c[k:])
	}
}
