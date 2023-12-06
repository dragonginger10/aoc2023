package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type section struct {
  sectionName string
  sectionInts [][]int
}

type mapping struct {
  destinationRange int
  sourceRange int
  rangeLength int
}

func parseIntList(str string) []int {
  var result []int
  seperated := strings.Split(str, " ")
  for _, s := range seperated {
    i, _ := strconv.ParseInt(s, 10, 0)
    result = append(result, int(i))
  }
  return result
}

func ReadInput(file string) []string {
  var result []string
	rawFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
  defer rawFile.Close()

  scanner := bufio.NewScanner(rawFile)
  for scanner.Scan() {
    readLine := scanner.Text()
    result = append(result, readLine)
  }
  return result
}

func getSection(almanacPart []string) []section {
  var result []section
  re := regexp.MustCompile(` map:`)
  var newSection section
  for _, line := range almanacPart {
    if re.Match([]byte(line)) {
      name := re.ReplaceAll([]byte(line), []byte{})
      newSection.sectionName = string(name)
      continue
    }
    if line == ""{
      result = append(result, newSection)
      newSection.sectionInts = make([][]int, 0)
      continue
    }
    parsed := parseIntList(line)
    newSection.sectionInts = append(newSection.sectionInts, parsed)
  }
  return result
}

func main() {
	almanac := ReadInput(os.Args[1])

  seeds := strings.Split(almanac[0], ": ")
  seeds = strings.Split(seeds[1], " ")

  sections := getSection(almanac[2:])

  for _, section := range sections {
    fmt.Println(section.sectionName)
    fmt.Println(section.sectionInts)
    fmt.Println("********************************************************************************")
  }

}
