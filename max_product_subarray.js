nums = [-3, -1, -1];
let max = 0;
let current = 1;

for (let i = 0; i < nums.length; i++) {
  current *= nums[i];
  max = Math.max(max, current);
  if (max < 0) max = 0;
}

console.log(max);
