func sortColors(nums []int)  {
    n:=len(nums)//bubble sort
    for i:=0;i<n-1;i++{
        swapped:=false//for check swapped once
        for j:=0;j<n-i-1;j++{//move the end to the left because largest always at the end
            if nums[j]>nums[j+1]{
                nums[j],nums[j+1]=nums[j+1],nums[j]
                swapped=true
            }
        }
        if !swapped{//not swapped it will exit with o(n)
            break
        }
    }
}