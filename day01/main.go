package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "math"
  "sort"
  "errors"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
  file, err := os.Open("input.txt")
  check(err)
  defer file.Close()

  leftList := []int{}
  rightList := []int{}
  currentRow := 0

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    values := strings.Split(line, "   ")

    leftValue, err := strconv.Atoi(values[0])
		check(err)
		rightValue, err := strconv.Atoi(values[1])
		check(err)

    leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
		currentRow++
  }

  totalDistance, err := getTotalDistance(leftList, rightList)
  check(err)

  similarityScore := getSimilarityScore(leftList, rightList)

  fmt.Println("Total Distance:", totalDistance)
  fmt.Println("Similarity Score:", similarityScore)
}

func getTotalDistance(leftList, rightList []int) (totalDistance int, err error){
  sort.Ints(leftList)
  sort.Ints(rightList)

	if len(leftList) != len(rightList) {
    return 0, errors.New("Slices are not the same length")
	}

  totalDistance = 0

  for i := range len(leftList) {
    totalDistance += int(math.Abs(float64(leftList[i] - rightList[i]))) 
  }
  
  return totalDistance, nil
}

func getSimilarityScore(leftList, rightList []int) (similarityScore int) {

  similarityScore = 0

  var count int 
  for _, target := range leftList {
    count = 0
    for _, num := range rightList {
      if num == target {
        count ++
      }
    }
    similarityScore += target * count
  }

  return similarityScore
}
