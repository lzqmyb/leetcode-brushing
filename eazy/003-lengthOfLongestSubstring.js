/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    if (s.length == 0) {
        return 0;
    }
    let max = 0
    let start = 0;
    let keyIndex = {}

    for (let i = 0; i < s.length; i++) {
        let curStr = s[i]
        // console.log("index: " +i)
        if (keyIndex[curStr] != null && keyIndex[curStr] >= start) {
            console.log("value: " + keyIndex[curStr])
            start = keyIndex[curStr] + 1
        } else {
            let length = i - start + 1

            // console.log(i + ":" + length)
            if (max < length) {
                max = length
            }
        }
        keyIndex[curStr] = i

    }
    return max;
};

/**
 * @param {string} s
 * @return {number}
 */
var offcialLengthOfLongestSubstring = function(s) {
    let num=0,j=0,t=0
    for(let i=0;i<s.length;i++){
        t = s.slice(j,i).indexOf(s[i]);
        if(t == -1){
            num=Math.max(num,i-j+1)
        }else{
            j=j+t+1
        }
    }
    return num
};

// let number = lengthOfLongestSubstring("abcabcbb");
let number = lengthOfLongestSubstring("tmmzuxt");
console.log(number);
