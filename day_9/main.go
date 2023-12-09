package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(file string) []string {
	var result []string
	rawFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	result = strings.Split(string(rawFile), "\n")
	return result
}

func parseListOfInt(list []string) []int {
  var result []int
  for _, s := range list {
    i, err := strconv.ParseInt(s, 10, 0)
    if err != nil {
      log.Fatal(err)
    }
    result = append(result, int(i))
  }
  return result
}

func predictValue(value []int) int {
  var result int
  var count int
  lists := make(map[int][]int, 0)
  lists[count] = value

  for {
    count++
    lists[count] = differences(lists[count-1])
    if listOfZeros(lists[count]) {
      break
    }
  }

  for {
    count--
    list := lists[count]
    result += list[len(list)-1]
    if count == 0 {
      break
    }
  }

  return result
}

func differences(list []int) []int {
  var result []int
  for i, v := range list {
    if i+1 >= len(list) {
      break
    }
    x := list[i+1] - v
    result = append(result, x)
  }
  return result
}

func listOfZeros(list []int) bool {
  for _, v := range list {
    if v != 0{
      return false
    }
  }
  return true
}

func predictEalier(value []int) int {
  var result int
  var count int
  lists := make(map[int][]int, 0)
  lists[count] = value

  for {
    count++
    lists[count] = differences(lists[count-1])
    if listOfZeros(lists[count]) {
      break
    }
  }

  for {
    count--
    list := lists[count]
    result = list[0] - result
    if count == 0 {
      break
    }
  }
  
  return result
}

func main() {
	var answer int
  var part2 int
  values := ReadInput(os.Args[1])
  for _, v := range values {
    if v == "" {
      continue
    }
    listOfValue := strings.Split(v, " ")
    nextValue := predictValue(parseListOfInt(listOfValue))
    lastValue := predictEalier(parseListOfInt(listOfValue))
    // fmt.Println(v, nextValue)
    answer += nextValue
    part2 += lastValue
  }
        
	fmt.Println(answer, part2)
}
