/**
 * @param {string[]} word1
 * @param {string[]} word2
 * @return {boolean}
 */
var arrayStringsAreEqual = function(word1, word2) {
    const first=word1.slice("").join("")
    const second=word2.slice("").join("")
    return first===second
};