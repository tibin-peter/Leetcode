/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    left,right:=1,n
    for left<=right{
        center:=left+(right-left)/2
        if guess(center)==0{
            return center
        }else if guess(center)==1{
            left=center+1
        }else{
            right=center-1
        }
    }
    return -1
}