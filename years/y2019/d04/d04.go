package y2019d04

// return all valid passwords within the given range that meet the following criteria:
// 1. 6 - digit number
// 2. atleast 2 adjacent digits are the same
// 3. from left to right, a digit never decreases
func RunDay04Pt1(min, max int) int {
    validPwds := 0
    for i := min; i < max; i++ {
        if isValid(i, sixDigit, adjacentValues, neverDecreased) { validPwds++ }
    }
    return validPwds
}

func RunDay04Pt2(min, max int) int {
    validPwds := 0
    for i := min; i < max; i++ {
        if isValid(i, sixDigit, adjacentValues, neverDecreased, onlyAdjacentValues) { validPwds++ }
    }
    return validPwds
}

func isValid(pwd int, validation ...func(int) bool) bool {
    valid := true
    for _, validate := range validation {
        valid = valid && validate(pwd)
    }
    return valid
}

func sixDigit(pwd int) bool { return pwd > 99_999 && pwd < 1_000_000 }

func adjacentValues(pwd int) bool {
    currPwd := pwd
    curr := pwd % 10
    for currPwd > 0 {
        nextPwd := currPwd / 10
        next := nextPwd % 10
        if curr == next { return true }
        currPwd = nextPwd
        curr = next
    }
    return false
}

func onlyAdjacentValues(pwd int) bool {
    currPwd := pwd
    curr := pwd % 10
    seenCount := 0
    for currPwd > 0 {
        nextPwd := currPwd / 10
        next := nextPwd % 10
        if curr == next { 
            seenCount++
        } else {
            if seenCount == 1 { return true }
            seenCount = 0
        }
        currPwd = nextPwd
        curr = next
    }
    return seenCount == 1
}

func neverDecreased(pwd int) bool {
    currPwd := pwd
    curr := pwd % 10
    for currPwd > 0 {
        nextPwd := currPwd / 10
        next := nextPwd % 10
        if next > curr { return false }
        currPwd = nextPwd
        curr = next
    }
    return true
}

// Password is valid if:
// 1. 6 - digit number
// 2. atleast 2 adjacent digits are the same
// 3. from left to right, a digit never decreases
