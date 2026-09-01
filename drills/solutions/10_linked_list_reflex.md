# Drill 10 — Linked Lists

## reverseList
- **Trigger:** walk a list backwards, or rebuild it reversed
- **Pattern:** `prev`/`cur`/`next` rewire; return `prev`, not `cur`
- **Bug:** losing `cur.Next` before reassigning it — save `next` first

## hasCycle
- **Trigger:** "is there a loop", with O(1) extra space demanded
- **Pattern:** Floyd — slow one step, fast two; they meet iff a cycle exists
- **Bug:** checking `fast != nil` only — `fast.Next` can be nil on the same step

## middleNode
- **Trigger:** middle of the list in one pass, length unknown
- **Pattern:** slow+fast; slow lands on the middle when fast runs off the end
- **Even length** returns the *second* middle with this loop shape

## mergeTwoLists
- **Trigger:** merge sorted inputs without materialising a slice
- **Pattern:** dummy head + tail pointer; append the smaller front each step
- **Bug:** forgetting to attach the non-empty remainder after the loop

## removeNthFromEnd
- **Trigger:** "nth from the end", one pass, no length available
- **Pattern:** lead pointer runs n ahead, then both advance; trail stops before the target
- **Dummy head** is what makes deleting the first node the same code as any other
- **Bug:** n larger than the list — decide (here: return the list unchanged)
