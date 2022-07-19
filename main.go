package main

import (
	"fmt"
	"sync"
)

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	var wg sync.WaitGroup

	wg.Add(len(n))
	for i := 0; i < len(n); i++ {
		sl := n[i]
		indexSlice := i + 1 // Роблю це для того щоб був гарний вигляд типу слайс 1 слайс 2 і 3 а не слайс 0 і так далі.
		go func() {
			defer wg.Done()
			sum(sl, indexSlice)
		}()
	}
	wg.Wait()

}

func sum(arr []int, i int) { // Додав аргумент i для того щоб виводити номер слайсу.
	var sum int
	for _, num := range arr {
		sum += num
	}
	fmt.Printf("slice %v: %v\n", i, sum)

}

//for i, elem := range n {
//	wg.Add(1)
//	go func() {
//		defer wg.Done()
//		sum(elem, i+1)
//	}()
//}
//
//wg.Wait()

//вивід
//slice 3: 49
//slice 3: 49
//slice 3: 49