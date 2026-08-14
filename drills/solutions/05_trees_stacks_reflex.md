# Drill 05 — Trees & Stacks

## inorderTraversal (iterative)
- **Pattern:** stack simulates recursion; go left, process, go right

## preorderTraversal (iterative)
- **Pattern:** process node, then push right then left (stack LIFO gives left first)

## postorderTraversal (iterative)
- **Pattern:** reverse-preorder (process, push left then right, reverse result)

## levelOrderTraversal (BFS)
- **Pattern:** queue; dequeue node, enqueue children left-to-right

## maxDepth
- **Recurrence:** 1 + max(left, right); base nil → 0

## isValidParentheses
- **Pattern:** push openers; pop on closer with match check
- **Bug:** forget to check stack empty at end

## dailyTemperatures
- **Pattern:** monotonic decreasing stack of indices
- **Pop** when current temp warmer; set answer for popped index
