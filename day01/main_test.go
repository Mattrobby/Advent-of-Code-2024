package main

import (
  "testing"
)

func TestGetTotalDistance(t *testing.T)  {
  tests := []struct {
    leftList  []int
    rightList []int
    expected  int
  }{
    {
      []int{3, 4, 2, 1, 3, 3},
      []int{4, 3, 5, 3, 9, 3},
      11,
    },
  }

  for _, test := range tests {
    result, _ := getTotalDistance(test.leftList, test.rightList)

    if result != test.expected {
      t.Errorf("getTotalDistance(%v, %v) = %d; Expected %d", test.leftList, test.rightList, result, test.expected)
    }
  }
}

func TestSimilarityScore(t *testing.T)  {
  tests := []struct {
    leftList  []int
    rightList []int
    expected  int
  }{
    {
      []int{3, 4, 2, 1, 3, 3},
      []int{4, 3, 5, 3, 9, 3},
      31,
    },
  }

  for _, test := range tests {
    result := getSimilarityScore(test.leftList, test.rightList)

    if result != test.expected {
      t.Errorf("getSimilarityScore(%v, %v) = %d; Expected %d", test.leftList, test.rightList, result, test.expected)
    }
  }
}

