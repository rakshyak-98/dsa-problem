# Drill 02 — Hashing

## twoSum
- **Trigger:** complement lookup
- **Bug:** store index after check, not before

## containsDuplicate
- **Trigger:** seen before?
- **Pattern:** set or map keys
- **Bug:** using slice scan O(n²)

## groupAnagrams
- **Trigger:** same multiset of letters
- **Pattern:** key = sorted string OR [26]int count
- **Bug:** using raw string as key when order differs

## firstUniqueChar
- **Trigger:** count then second pass
- **Pattern:** freq map → scan string for count==1
- **Bug:** return byte not string
