partialsum
==========

A Go library for succinct partial sums data structure

PartialSum stores non-negative integers V[0...n) and supports operations
in O(1) time using at most (s + n) bits where s is the sum of V[0...n).

PartialSum supports following operations in O(1) time

	- IncTail(ind, val) V[ind] += val. ind should be the last one, that is ind >= Num
	- Lookup(ind) returns V[ind]
	- Sum(ind) returns V[0]+V[1]+...+V[ind-1]
	- Find(val) returns ind satisfying Sum(ind) <= val < Sum(ind+1)
		and offset = val - Sum(ind). If there are multiple inds
		satisfiying this condition, return the minimum one.

PartialSum is sutable for storing small (but somethimes large) non-negative integers
such as lengths of string.

Note that if we store S[0...n) using interger array (e.g. []uint64),
where S[i] = Sum(i), then S requires O(log n) time for Find(), and needs 64n bits.

|       | PartialSum | []uint64 |
|-------|------------|-----------------------|
|Space(bits) | n + s | 64n |
|Lookup | O(1) | O(1) |
|Sum | O(1) | O(1) |
|Find | O(1) | O(log n) |

partialsum uses [rsdic](http://github.com/hillbig/rsdic/) (succinct rank/select dictionary)


Usage
=====
```
import "github.com/hillbig/partialsum"

ps := partialsum.New()

ps.IncTail(0, 5)
ps.IncTail(2, 4)
ps.IncTail(2, 2)
ps.IncTail(3, 3)

// ps stores [5, 0, 6, 3]
// S = [0, 5, 5, 13, 14]

ps.Num() // 4
ps.AllSum() // 14
ps.Lookup(2) // 6
ps.Sum(2) // 5
ps.Find(5) // 2, 0 because S[2] <= 5 < S[3], and 5 - S[2] = 0
```
