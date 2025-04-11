package increasingtripletsubsequence

func increasingTriplet(nums []int) bool {
    first := int(^uint(0) >> 1) // MaxInt
    second := int(^uint(0) >> 1)

    for _, num := range nums {
        if num <= first {
            first = num
        } else if num <= second {
            second = num
        } else {
            // num > second → нашли третье число
            return true
        }
    }

    return false
}
