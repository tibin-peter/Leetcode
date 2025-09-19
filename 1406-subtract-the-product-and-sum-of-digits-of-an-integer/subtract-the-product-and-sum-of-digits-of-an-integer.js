/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function(n) {
    let str= n.toString().split("")
   let newarray = str.map(Number)
   let sum = newarray.reduce((total,num)=> total+num ,0)
   let mul = newarray.reduce((total,num)=> total*num ,1)
   let result=mul-sum
   return  result
};