# Drill 08 — Heaps (bonus)

## kthLargest
- **Pattern:** min-heap size k OR quickselect
- **Heap:** push all, pop until k left; top is answer

## lastStoneWeight
- **Pattern:** max-heap; pop two largest, push diff if non-zero

## mergeKSorted
- **Pattern:** min-heap of (value, listIndex, elemIndex)
- **Push** next element from list when one is popped
