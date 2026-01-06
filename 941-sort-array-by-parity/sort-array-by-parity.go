func sortArrayByParity(nums []int) []int {
    pointer:=0

    for i:=0;i<len(nums);i++{
        if nums[i]%2==0{
            nums[pointer],nums[i]=nums[i],nums[pointer]
            pointer++
        }
    }
    return nums
}