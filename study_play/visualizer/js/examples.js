/**
 * Example algorithms with step-by-step visualization.
 */
const EXAMPLES = {
  array: [
    {
      id: "bubble-sort",
      name: "Bubble Sort",
      input: "5, 2, 8, 1, 9, 3",
      code: `async function run(viz, data) {
  const arr = [...data];
  viz.setArray(arr);
  await viz.step({ message: "Start bubble sort", line: 2, variables: { n: arr.length } });

  for (let i = 0; i < arr.length; i++) {
    for (let j = 0; j < arr.length - i - 1; j++) {
      viz.highlight([j, j + 1], "compare");
      viz.setPointers({ left: j, right: j + 1 });
      await viz.step({
        message: \`Compare arr[\${j}]=\${arr[j]} and arr[\${j + 1}]=\${arr[j + 1]}\`,
        line: 7,
        variables: { i, j },
      });

      if (arr[j] > arr[j + 1]) {
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
        viz.setArray([...arr]);
        viz.highlight([j, j + 1], "swap");
        await viz.step({ message: "Swapped adjacent elements", line: 12, array: [...arr] });
      }
    }
    viz.highlight([arr.length - i - 1], "found");
    await viz.step({ message: \`Position \${arr.length - i - 1} is sorted\`, line: 16 });
  }

  viz.highlight(arr.map((_, i) => i), "found");
  await viz.step({ message: "Array fully sorted!", line: 20 });
  return arr;
}`,
    },
    {
      id: "two-sum",
      name: "Two Sum (hash map)",
      input: "2, 7, 11, 15",
      code: `async function run(viz, data) {
  const nums = [...data];
  const target = 9;
  const map = new Map();
  viz.setArray(nums);
  await viz.step({ message: \`Find two numbers that sum to \${target}\`, line: 4, variables: { target } });

  for (let i = 0; i < nums.length; i++) {
    const complement = target - nums[i];
    viz.highlight([i], "compare");
    viz.setPointers({ i });
    await viz.step({
      message: \`nums[\${i}]=\${nums[i]}, need complement \${complement}\`,
      line: 9,
      variables: { i, complement, map: Object.fromEntries(map) },
    });

    if (map.has(complement)) {
      const j = map.get(complement);
      viz.highlight([i, j], "found");
      viz.setPointers({ i, j });
      await viz.step({ message: \`Found pair: [\${j}, \${i}]\`, line: 14, variables: { result: [j, i] } });
      return [j, i];
    }

    map.set(nums[i], i);
    viz.highlight([i], "visited");
    await viz.step({ message: \`Store \${nums[i]} → index \${i} in hash map\`, line: 19, variables: { map: Object.fromEntries(map) } });
  }
  return [-1, -1];
}`,
    },
    {
      id: "binary-search",
      name: "Binary Search",
      input: "1, 3, 5, 7, 9, 11, 13",
      code: `async function run(viz, data) {
  const arr = [...data].sort((a, b) => a - b);
  const target = 7;
  let left = 0;
  let right = arr.length - 1;
  viz.setArray(arr);
  await viz.step({ message: \`Search for \${target} in sorted array\`, line: 4, variables: { target, left, right } });

  while (left <= right) {
    const mid = Math.floor((left + right) / 2);
    viz.setPointers({ left, mid, right });
    viz.highlight([left, mid, right], "compare");
    await viz.step({
      message: \`mid=\${mid}, arr[mid]=\${arr[mid]}\`,
      line: 10,
      variables: { left, mid, right },
    });

    if (arr[mid] === target) {
      viz.highlight([mid], "found");
      await viz.step({ message: \`Found \${target} at index \${mid}!\`, line: 15 });
      return mid;
    }
    if (arr[mid] < target) {
      left = mid + 1;
      await viz.step({ message: "Target is in right half", line: 19, variables: { left, right } });
    } else {
      right = mid - 1;
      await viz.step({ message: "Target is in left half", line: 22, variables: { left, right } });
    }
  }
  await viz.step({ message: "Not found", line: 26 });
  return -1;
}`,
    },
    {
      id: "two-pointers",
      name: "Two Pointers — Palindrome Check",
      input: "1, 2, 2, 1",
      code: `async function run(viz, data) {
  const arr = [...data];
  let left = 0;
  let right = arr.length - 1;
  viz.setArray(arr);
  await viz.step({ message: "Check if array is palindrome", line: 4, variables: { left, right } });

  while (left < right) {
    viz.setPointers({ left, right });
    viz.highlight([left, right], "compare");
    await viz.step({
      message: \`Compare arr[\${left}]=\${arr[left]} vs arr[\${right}]=\${arr[right]}\`,
      line: 9,
      variables: { left, right },
    });

    if (arr[left] !== arr[right]) {
      viz.highlight([left, right], "swap");
      await viz.step({ message: "Mismatch — not a palindrome", line: 14 });
      return false;
    }
    left++;
    right--;
  }

  viz.highlight(arr.map((_, i) => i), "found");
  await viz.step({ message: "Palindrome confirmed!", line: 22 });
  return true;
}`,
    },
  ],

  string: [
    {
      id: "palindrome",
      name: "Palindrome Check",
      input: "racecar",
      code: `async function run(viz, data) {
  const s = String(data);
  let left = 0;
  let right = s.length - 1;
  viz.setString(s);
  await viz.step({ message: "Check if string is a palindrome", line: 4, variables: { left, right } });

  while (left < right) {
    viz.setPointers({ left, right });
    viz.highlight([left, right], "compare");
    await viz.step({
      message: \`Compare s[\${left}]='\${s[left]}' vs s[\${right}]='\${s[right]}'\`,
      line: 9,
      variables: { left, right },
    });

    if (s[left] !== s[right]) {
      viz.highlight([left, right], "mismatch");
      await viz.step({ message: "Mismatch — not a palindrome", line: 14 });
      return false;
    }
    viz.highlight([left, right], "match");
    await viz.step({ message: "Characters match, move inward", line: 18 });
    left++;
    right--;
  }

  viz.highlight([...Array(s.length).keys()], "found");
  await viz.step({ message: "Palindrome confirmed!", line: 24 });
  return true;
}`,
    },
    {
      id: "valid-anagram",
      name: "Valid Anagram",
      input: "anagram",
      code: `async function run(viz, data) {
  const s = String(data);
  const t = "nagaram";
  viz.setString(s);
  await viz.step({ message: \`Check if "\${s}" is anagram of "\${t}"\`, line: 4, variables: { s, t } });

  if (s.length !== t.length) return false;

  const count = {};
  for (let i = 0; i < s.length; i++) {
    viz.highlight([i], "compare");
    viz.setPointers({ i });
    count[s[i]] = (count[s[i]] || 0) + 1;
    await viz.step({
      message: \`Count s[\${i}]='\${s[i]}' → \${count[s[i]]}\`,
      line: 12,
      variables: { i, char: s[i], count: { ...count } },
    });
  }

  for (let j = 0; j < t.length; j++) {
    const ch = t[j];
    viz.setString(t);
    viz.highlight([j], "compare");
    viz.setPointers({ j });
    await viz.step({
      message: \`Check t[\${j}]='\${ch}' in count map\`,
      line: 22,
      variables: { j, char: ch, count: { ...count } },
    });

    if (!count[ch]) {
      viz.highlight([j], "mismatch");
      await viz.step({ message: \`Extra char '\${ch}' — not an anagram\`, line: 27 });
      return false;
    }
    count[ch]--;
    viz.highlight([j], "match");
    await viz.step({ message: \`Decrement count for '\${ch}'\`, line: 31, variables: { count: { ...count } } });
  }

  viz.setString(s);
  viz.highlight([...Array(s.length).keys()], "found");
  await viz.step({ message: "Valid anagram!", line: 36 });
  return true;
}`,
    },
    {
      id: "reverse-string",
      name: "Reverse String (two pointers)",
      input: "hello",
      code: `async function run(viz, data) {
  const chars = [...String(data)];
  let left = 0;
  let right = chars.length - 1;
  viz.setString(chars.join(""));
  await viz.step({ message: "Reverse string in-place with two pointers", line: 5, variables: { left, right } });

  while (left < right) {
    viz.setPointers({ left, right });
    viz.highlight([left, right], "compare");
    await viz.step({
      message: \`Swap chars at \${left} ('\${chars[left]}') and \${right} ('\${chars[right]}')\`,
      line: 10,
      variables: { left, right },
    });

    [chars[left], chars[right]] = [chars[right], chars[left]];
    viz.setString(chars.join(""));
    viz.highlight([left, right], "swap");
    await viz.step({ message: "Swapped!", line: 15, string: chars.join("") });

    left++;
    right--;
  }

  viz.highlight([...Array(chars.length).keys()], "found");
  await viz.step({ message: \`Reversed: "\${chars.join("")}"\`, line: 22 });
  return chars.join("");
}`,
    },
    {
      id: "longest-substring",
      name: "Longest Substring (sliding window)",
      input: "abcabcbb",
      code: `async function run(viz, data) {
  const s = String(data);
  const seen = new Set();
  let left = 0;
  let maxLen = 0;
  let best = [0, 0];
  viz.setString(s);
  await viz.step({ message: "Sliding window — longest substring without repeats", line: 6 });

  for (let right = 0; right < s.length; right++) {
    while (seen.has(s[right])) {
      seen.delete(s[left]);
      viz.highlight([left], "visited");
      await viz.step({
        message: \`Duplicate '\${s[right]}' — shrink window, remove s[\${left}]\`,
        line: 12,
        variables: { left, right, seen: [...seen] },
      });
      left++;
    }

    seen.add(s[right]);
    viz.setPointers({ left, right });
    viz.highlightRange(left, right, "compare");
    const len = right - left + 1;
    if (len > maxLen) {
      maxLen = len;
      best = [left, right];
    }
    await viz.step({
      message: \`Window [\${left}..\${right}] = "\${s.slice(left, right + 1)}" len=\${len}\`,
      line: 24,
      variables: { left, right, maxLen, window: s.slice(left, right + 1) },
    });
  }

  viz.highlightRange(best[0], best[1], "found");
  await viz.step({
    message: \`Longest substring: "\${s.slice(best[0], best[1] + 1)}" (length \${maxLen})\`,
    line: 30,
    variables: { maxLen, substring: s.slice(best[0], best[1] + 1) },
  });
  return maxLen;
}`,
    },
    {
      id: "valid-parentheses",
      name: "Valid Parentheses",
      input: "({[]})",
      code: `async function run(viz, data) {
  const s = String(data);
  const stack = [];
  const pairs = { ")": "(", "]": "[", "}": "{" };
  viz.setString(s);
  await viz.step({ message: "Match brackets using a stack", line: 5, variables: { stack: [] } });

  for (let i = 0; i < s.length; i++) {
    const ch = s[i];
    viz.highlight([i], "compare");
    viz.setPointers({ i });
    await viz.step({
      message: \`Process s[\${i}]='\${ch}'\`,
      line: 10,
      variables: { i, ch, stack: [...stack] },
    });

    if (ch === "(" || ch === "[" || ch === "{") {
      stack.push(ch);
      viz.highlight([i], "visited");
      await viz.step({ message: \`Open bracket — push '\${ch}'\`, line: 15, variables: { stack: [...stack] } });
    } else {
      if (stack.length === 0 || stack[stack.length - 1] !== pairs[ch]) {
        viz.highlight([i], "mismatch");
        await viz.step({ message: \`No match for '\${ch}'\`, line: 20 });
        return false;
      }
      stack.pop();
      viz.highlight([i], "match");
      await viz.step({ message: \`Matched '\${ch}' — pop stack\`, line: 25, variables: { stack: [...stack] } });
    }
  }

  const valid = stack.length === 0;
  if (valid) viz.highlight([...Array(s.length).keys()], "found");
  await viz.step({ message: valid ? "All brackets matched!" : "Unclosed brackets remain", line: 31, variables: { valid, stack } });
  return valid;
}`,
    },
  ],

  tree: [
    {
      id: "inorder",
      name: "Inorder Traversal (iterative)",
      input: "1, 2, 3, 4, 5, null, null",
      code: `async function run(viz, data) {
  const root = buildTreeFromLevelOrder(data);
  viz.setTree(root);
  const result = [];
  const stack = [];
  let current = root;
  await viz.step({ message: "Start inorder traversal (Left → Root → Right)", line: 6 });

  while (current || stack.length) {
    while (current) {
      stack.push(current);
      viz.highlightNodes([current.id]);
      await viz.step({
        message: \`Go left from node \${current.val}, push to stack\`,
        line: 11,
        variables: { stack: stack.map((n) => n.val), result: [...result] },
      });
      current = current.left;
    }

    current = stack.pop();
    result.push(current.val);
    viz.highlightNodes([current.id]);
    await viz.step({
      message: \`Visit node \${current.val}\`,
      line: 19,
      variables: { stack: stack.map((n) => n.val), result: [...result] },
    });
    current = current.right;
  }

  await viz.step({ message: \`Inorder result: [\${result.join(", ")}]\`, line: 26, variables: { result } });
  return result;
}`,
    },
    {
      id: "bfs-level",
      name: "BFS Level Order",
      input: "3, 9, 20, null, null, 15, 7",
      code: `async function run(viz, data) {
  const root = buildTreeFromLevelOrder(data);
  if (!root) return [];
  viz.setTree(root);
  const result = [];
  const queue = [root];
  await viz.step({ message: "BFS level-order traversal", line: 6 });

  while (queue.length) {
    const levelSize = queue.length;
    const level = [];
    await viz.step({
      message: \`Process level with \${levelSize} node(s)\`,
      line: 11,
      variables: { queue: queue.map((n) => n.val), levelSize },
    });

    for (let i = 0; i < levelSize; i++) {
      const node = queue.shift();
      level.push(node.val);
      viz.highlightNodes([node.id]);
      await viz.step({
        message: \`Visit node \${node.val}\`,
        line: 17,
        variables: { level: [...level], queue: queue.map((n) => n.val) },
      });

      if (node.left) queue.push(node.left);
      if (node.right) queue.push(node.right);
    }
    result.push(level);
  }

  await viz.step({ message: "BFS complete", line: 27, variables: { result } });
  return result;
}`,
    },
    {
      id: "max-depth",
      name: "Max Depth (DFS)",
      input: "1, 2, 3, 4, 5, null, null",
      code: `async function run(viz, data) {
  const root = buildTreeFromLevelOrder(data);
  viz.setTree(root);

  async function depth(node, d) {
    if (!node) {
      await viz.step({ message: \`Null node, depth \${d}\`, line: 6, variables: { depth: d } });
      return d;
    }
    viz.highlightNodes([node.id]);
    await viz.step({
      message: \`At node \${node.val}, current depth \${d}\`,
      line: 10,
      variables: { node: node.val, depth: d },
    });

    const left = await depth(node.left, d + 1);
    const right = await depth(node.right, d + 1);
    const maxD = Math.max(left, right);
    await viz.step({
      message: \`Node \${node.val}: max(\${left}, \${right}) = \${maxD}\`,
      line: 16,
      variables: { left, right, maxDepth: maxD },
    });
    return maxD;
  }

  const ans = await depth(root, 0);
  await viz.step({ message: \`Max depth = \${ans}\`, line: 22, variables: { maxDepth: ans } });
  return ans;
}`,
    },
  ],

  graph: [
    {
      id: "bfs-grid",
      name: "BFS Shortest Path",
      input: `0, 0, 0, 1
0, 1, 0, 0
0, 0, 0, 0
1, 0, 0, 0`,
      code: `async function run(viz, data) {
  const grid = data.map((row) => row.map(Number));
  const rows = grid.length;
  const cols = grid[0].length;
  viz.setGraph(grid);
  viz.setCellTypes({ "0,0": "start", [\`\${rows - 1},\${cols - 1}\`]: "end" });
  await viz.step({ message: "BFS from top-left to bottom-right", line: 6 });

  const dirs = [[0, 1], [1, 0], [0, -1], [-1, 0]];
  const queue = [[0, 0, 1]];
  const visited = new Set(["0,0"]);
  viz.highlightCells([[0, 0]], "highlight");
  await viz.step({ message: "Enqueue start (0,0)", line: 12, variables: { queue: [[0, 0, 1]] } });

  while (queue.length) {
    const [r, c, dist] = queue.shift();
    viz.highlightCells([[r, c]], "highlight");
    await viz.step({
      message: \`Visit (\${r},\${c}), distance=\${dist}\`,
      line: 17,
      variables: { r, c, dist, queueLen: queue.length },
    });

    if (r === rows - 1 && c === cols - 1) {
      viz.highlightCells([[r, c]], "path");
      await viz.step({ message: \`Reached goal! Shortest path = \${dist}\`, line: 22 });
      return dist;
    }

    for (const [dr, dc] of dirs) {
      const nr = r + dr;
      const nc = c + dc;
      const key = \`\${nr},\${nc}\`;
      if (nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] === 0 && !visited.has(key)) {
        visited.add(key);
        queue.push([nr, nc, dist + 1]);
        viz.highlightCells([[nr, nc]], "path");
        await viz.step({
          message: \`Enqueue neighbor (\${nr},\${nc})\`,
          line: 31,
          variables: { neighbor: [nr, nc], dist: dist + 1 },
        });
      }
    }
  }

  await viz.step({ message: "No path found", line: 37 });
  return -1;
}`,
    },
    {
      id: "flood-fill",
      name: "Flood Fill",
      input: `1, 1, 1
1, 1, 0
1, 0, 1`,
      code: `async function run(viz, data) {
  const image = data.map((row) => row.map(Number));
  const sr = 1, sc = 1, newColor = 2;
  const orig = image[sr][sc];
  viz.setGraph(image);
  viz.highlightCells([[sr, sc]], "highlight");
  await viz.step({
    message: \`Flood fill from (\${sr},\${sc}), orig=\${orig} → newColor=\${newColor}\`,
    line: 6,
    variables: { sr, sc, orig, newColor },
  });

  function dfs(r, c) {
    if (r < 0 || c < 0 || r >= image.length || c >= image[0].length) return;
    if (image[r][c] !== orig) return;
    image[r][c] = newColor;
    viz.setGraph(image.map((row) => [...row]));
    viz.highlightCells([[r, c]], "highlight");
  }

  const stack = [[sr, sc]];
  while (stack.length) {
    const [r, c] = stack.pop();
    if (r < 0 || c < 0 || r >= image.length || c >= image[0].length) continue;
    if (image[r][c] !== orig) continue;

    image[r][c] = newColor;
    viz.setGraph(image.map((row) => [...row]));
    viz.highlightCells([[r, c]], "highlight");
    await viz.step({
      message: \`Paint cell (\${r},\${c}) with color \${newColor}\`,
      line: 22,
      variables: { r, c, stackLen: stack.length },
    });

    stack.push([r + 1, c], [r - 1, c], [r, c + 1], [r, c - 1]);
  }

  await viz.step({ message: "Flood fill complete", line: 28 });
  return image;
}`,
    },
    {
      id: "num-islands",
      name: "Count Islands (DFS)",
      input: `1, 1, 0
0, 1, 0
1, 0, 1`,
      code: `async function run(viz, data) {
  const grid = data.map((row) => row.map(String));
  const rows = grid.length;
  const cols = grid[0].length;
  let count = 0;
  viz.setGraph(grid);
  await viz.step({ message: "Count islands ('1' = land)", line: 6, variables: { count } });

  async function dfs(r, c) {
    if (r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] !== "1") return;
    grid[r][c] = "0";
    viz.setGraph(grid.map((row) => [...row]));
    viz.highlightCells([[r, c]], "highlight");
    await viz.step({ message: \`Mark land at (\${r},\${c}) visited\`, line: 13 });

    await dfs(r + 1, c);
    await dfs(r - 1, c);
    await dfs(r, c + 1);
    await dfs(r, c - 1);
  }

  for (let r = 0; r < rows; r++) {
    for (let c = 0; c < cols; c++) {
      if (grid[r][c] === "1") {
        count++;
        viz.highlightCells([[r, c]], "path");
        await viz.step({
          message: \`New island #\${count} starting at (\${r},\${c})\`,
          line: 24,
          variables: { count, start: [r, c] },
        });
        await dfs(r, c);
      }
    }
  }

  await viz.step({ message: \`Total islands: \${count}\`, line: 31, variables: { count } });
  return count;
}`,
    },
  ],
};

if (typeof module !== "undefined") {
  module.exports = { EXAMPLES };
}
