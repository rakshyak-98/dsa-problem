/**
 * @param {string} licensePlate
 * @param {string[]} words
 * @return {string}
 */
var shortestCompletingWord = function (licensePlate, words) {
  const getCounts = (str) => {
    const counts = new Array(26).fill(0);

    for (let char of str.toLowerCase()) {
      if (char >= "a" && char <= "z") {
        counts[char.charCodeAt(0) - 97]++;
      }
    }
    return counts;
  };

  let targetCounts = getCounts(licensePlate);
  let result = null;
  for (let word of words) {
    if (result !== null && word.length >= result.length) continue;

    const wordCounts = getCounts(word);
    let isMatch = true;

    for (let i = 0; i < 26; i++) {
      if (wordCounts[i] < targetCounts[i]) {
        isMatch = false;
        break;
      }
    }

    if (isMatch) {
      result = word;
    }
  }
  return result;
};

console.log(
  shortestCompletingWord("1s3 PSt", ["step", "steps", "stripe", "stepple"]),
);

console.log(
  shortestCompletingWord("1s3 456", ["looks", "pest", "stew", "show"]),
);
