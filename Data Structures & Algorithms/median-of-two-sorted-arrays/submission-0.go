func sortTwoArrays(nums1 []int, nums2 []int) []int {
    nums := make([]int, 0, len(nums1)+len(nums2))
    nums = append(nums, nums1...)
    nums = append(nums, nums2...)
    sort.Ints(nums)
    return nums
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    nums := sortTwoArrays(nums1, nums2)
    center := len(nums) / 2
    if len(nums) % 2 == 0 {
        return (float64(nums[center]) + float64(nums[center-1])) / 2
    }
    
    return float64(nums[center])
}
