func search(nums []int, target int) int {
    left,right:=0,len(nums)-1

    for left<=right{
        center :=left+(right-left)/2//intigier division avoid pointed value

        if nums[center]==target{
            return center
        }else if nums[center]<target{//if targer greter left moves
            left=center+1
        }else{
            right=center-1//if target less tight moves
        }
    }
    return -1//if not found
}