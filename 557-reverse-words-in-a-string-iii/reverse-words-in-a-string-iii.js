/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function(s) {
    let result=s.split(' ')
    let reverse=result.map(s=>{
       return s.split("").reverse().join('')
    })
    return reverse.join(' ')
};