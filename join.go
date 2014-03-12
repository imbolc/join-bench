package main

import (
	"fmt"
	"time"
)

func main() {
	k := 1000 * 1000
	k = 100 * 1000

	fmt.Println("K = ", k)

	t := time.Now()
	d1 := make([][]int, 10*k)
	for i := range d1 {
		d1[i] = []int{1, i, 10}
	}
	fmt.Printf("%.8f\t d1 list\n", time.Since(t).Seconds())

	t = time.Now()
	d2 := make([][]int, k)
	for i := range d2 {
		d2[i] = []int{1, i + 2*k, 5}
	}
	fmt.Printf("%.8f\t d2 list\n", time.Since(t).Seconds())

	tjoin := time.Now()
	fmt.Println("\n--- start join")

	t = time.Now()
	d2_hash := make(map[string]int, len(d2))
	for i := range d2 {
		k := fmt.Sprintf("%d-%d", d2[i][0], d2[i][1])
		d2_hash[k] = d2[i][2]
	}
	fmt.Printf("%.8f\t d2 hash\n", time.Since(t).Seconds())

	t = time.Now()
	result_hash := map[string]string{}
	for i := range d1 {
		k := fmt.Sprintf("%d-%d", d1[i][0], d1[i][1])
		d2_v, found := d2_hash[k]
		if found {
			result_hash[k] = fmt.Sprintf("%d-%d", d1[i][2], d2_v)
		} else {
			result_hash[k] = fmt.Sprintf("%d-%d", d1[i][2], 0)
		}
	}
	fmt.Printf("%.8f\t d1 part to result hash\n", time.Since(t).Seconds())

	t = time.Now()
	for i := range d2 {
		k := fmt.Sprintf("%d-%d", d2[i][0], d2[i][1])
		_, found := result_hash[k]
		if !found {
			result_hash[k] = fmt.Sprintf("%d-%d", 0, d2[i][2])
		}
	}
	fmt.Printf("%.8f\t add d2 part to result hash\n", time.Since(t).Seconds())

	fmt.Printf("--- full join time: %.8f\n\n", time.Since(tjoin).Seconds())

	t = time.Now()
	result := make([][]string, len(result_hash))
	i := 0
	for k, v := range result_hash {
		result[i] = []string{k, v}
		i++
	}
	fmt.Printf("%.8f\t format result_hast to table\n\n", time.Since(t).Seconds())

	fmt.Println(result[:10])

}
