package luhn

import (
  "strings"
  "strconv"
)

func Valid(num string) bool {
  if (len(num) == 0) {
    return false
  }

  digits := convDigits(num)
  revdigits := reverse(digits)

  sum := 0
  for i, n := range revdigits {
    if i % 2 == 1 {
      n = digitalRoot(n * 2)
    }
    sum += n
  }

  return sum % 10 == 0
}

func reverse(arr1 []int) ([]int) {
  length := len(arr1)
  arr2 := make([]int, length)

  for i := range arr1 {
    arr2[length - i - 1] = arr1[i]
  }

  return arr2
}

func convDigits(s string) (digits []int) {
  pieces := strings.Split(s, "")
  digits = make([]int, len(pieces))

  for i := range pieces {
    digits[i], _ = strconv.Atoi(pieces[i])
  }
  return
}

func digitalRoot(num int) int {
  if num != 0 {
    num = 1 + ((num - 1) % 9)
  }

  return num
}
