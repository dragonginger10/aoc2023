package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type blockInfo struct {
	heatLoss int
	row      int
	column   int
	dr       int
	dc       int
	steps    int
	index    int
}

type PriorityQueue []*blockInfo

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*blockInfo)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func ReadInput(file string) []string {
	var result []string
	text, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer text.Close()
	s := bufio.NewScanner(text)
	for s.Scan() {
		result = append(result, s.Text())
	}
	return result
}

func parseToInt(r rune) int {
	number, err := strconv.ParseInt(string(r), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(number)
}

func main() {
	var numbers [][]int
	visited := make(map[string]bool)
	pq := &PriorityQueue{}
	heap.Init(pq)
	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	grid := ReadInput(os.Args[1])

	for _, row := range grid {
		rowList := make([]int, 0)
		for _, ch := range row {
			rowList = append(rowList, parseToInt(ch))
		}
		numbers = append(numbers, rowList)
	}

	heap.Push(pq, &blockInfo{
		heatLoss: 0,
		column:   0,
		row:      0,
		dr:       0,
		dc:       0,
		steps:    0,
	})

	for pq.Len() > 0 {
		cb := heap.Pop(pq).(*blockInfo)

		if cb.row == len(numbers)-1 && cb.column == len(numbers[0])-1 {
			fmt.Println(cb.heatLoss)
			break
		}

		key := fmt.Sprintf("%d,%d,%d,%d,%d", cb.row, cb.column, cb.dr, cb.dc, cb.steps)

		if _, ok := visited[key]; ok {
			continue
		}

		visited[key] = true

		for _, d := range directions {
			nrd := d[0]
			ncd := d[1]
			newRow := cb.row + nrd
			newColumn := cb.column + ncd

			if 0 > newRow || newRow > len(numbers)-1 || 0 > newColumn || newColumn > len(numbers[0])-1 {
				continue
			}
			if cb.dr != -nrd && cb.dc != -ncd {
				continue
			}

			c := 1
			if cb.dr == nrd && cb.dc == ncd {
				c += cb.steps
			}
			if c > 3 {
				continue
			}

			fmt.Println(newRow, newColumn)

			heap.Push(pq, &blockInfo{
				heatLoss: cb.heatLoss + numbers[newRow][newColumn],
				row:      newRow,
				column:   newColumn,
				dr:       nrd,
				dc:       ncd,
				steps:    c,
			})
		}
	}
}
