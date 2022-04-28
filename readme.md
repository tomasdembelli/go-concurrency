# Go-Concurrency
Notes from [Concurrency in Go](https://katherine.cox-buday.com/concurrency-in-go/) by Katherine Cox-Buday.


Contents
| dir name | subject |
|----------|---------|
| character-counter | sync mutex |
| context-switch | bespoke benchmark |
| recover | how (not) to panic |
| cond | conditional wait `sync.Cond` |


## Atomicity:
Something is considered `atomic`, or to have the `property of atomicity`, if it is `indivisible` or `uninterruptible` winthin the context that it is operating.

The `context` is very important. An operation in one context might not be atomic in another! (process vs operating system)

For example, `i++` has three atomic oprations (read `i`, increase `i`, write `i`), but it may not be atomic itself unless the necessary techniques are employed for a particular scenario. 
If the entire operation is handled within a single goroutine and this goroutine does not expose the `i`, then this operation is atomic, and there is no need to anything special. 
But, if `i` is a shared variable across multiple goruitens, then we need to protect this operation against `race condition`s where more than one (concurrent) gourutine `race` to manipulate `i`.

_See page 8 for solving this issue with `sync.Mutex`._


## Memory Access Synchronization
Synchronize access to the memory between the critical sections. `Critical section` is the part of your program that needs exclusive access to a shared resource. It is very important to define the size of critial sections.

## Deadlock
All concurrent processes are waiting on one another. The program will never recover without outside intervention. _See page 10 for an example_

## Livelock
Two or more concurrent processes try to prevent a deadlock without a coordination. It might not appear as an issue, because the program keeps running. But in reality, the state of the program never advances.
_Example page 14_

## Starvation
Because of poor memory synchronization, one or more processes can hold on the shared memory more than others, and let them starve. Starvation can happen due to outside players, like CPU, file handles, DB connections. There might be other programs (outside of your application) consuming the shared resources.

Broadining the critical section would allow a gorotuine do more jobs. However, it also makes other goroutines starve. So, the balance should be found between `coarse-grained synchronization for performance` and ` fine-grained synchronization for fairness`.






