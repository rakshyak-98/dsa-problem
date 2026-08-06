# Drill 01 — Arrays

## reverseInPlace
- **Trigger:** in-place reverse
- **Pattern:** two pointers L/R swap
- **Bug:** forget to return same slice

## indexOfMax
- **Trigger:** scan for best index
- **Pattern:** single pass, track best
- **Bug:** return 0 on empty (decide: -1 or panic)

## rotateRight
- **Trigger:** rotate by k
- **Pattern:** k %= n; reverse whole, reverse [0:k), reverse [k:n)
- **Bug:** k not reduced modulo n

## runningSum
- **Trigger:** prefix output
- **Pattern:** out[i] = out[i-1] + arr[i]
- **Bug:** mutate input instead of new slice
