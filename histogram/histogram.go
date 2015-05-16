package histogram

import (
    "fmt"
    "sort"
    "strings"
)

type IntHistogram map[int]int

func (h IntHistogram) Normalized() map[int]float64 {
    r := make(map[int]float64)
    total := 0
    for _, v := range(h) { total += v }
    for k, v := range(h) {
	r[k] = float64(v) / float64(total)
    }
    return r
}

func (h IntHistogram) String() (str string) {
    n := h.Normalized()

    keys := []int{}
    for k, _ := range(h) { keys = append(keys, int(k)) }
    sort.Ints(keys)
    for _, k := range(keys) {
	str += fmt.Sprintf("%2d (%6.2f%%): %s\n", k, 100*n[k], strings.Repeat("#", int(50 * n[k])))
    }
    return
}
