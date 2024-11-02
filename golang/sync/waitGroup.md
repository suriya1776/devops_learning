# waitGroup in golang

In Go, a WaitGroup is a synchronization construct that allows you to wait for a collection of goroutines to finish executing. It’s particularly useful when you have multiple tasks that need to run in parallel, and you want to ensure they all complete before moving forward.


![Waitgroup Image](./waitGroup.gif)


## Key Methods of WaitGroup
- Add(delta int): Increments the WaitGroup counter by the specified amount (delta). Typically, you call this with the number of goroutines you plan to launch.

- Done(): Decrements the WaitGroup counter by 1. Each goroutine should call this when it finishes its work.


- Wait(): Blocks the execution of the calling goroutine until the counter becomes zero, meaning all tasks have finished.



The WaitGroup counter starts at zero. Each time Add is called with a positive value, the counter increases. Every time Done is called (or Add is called with a negative value), the counter decreases. When the counter reaches zero, Wait unblocks.


## waitGroup in code

[wait group in code](./waitGroup.go)

```
var wg sync.WaitGroup
```

- This declares a global WaitGroup variable wg, which we’ll use to manage the synchronization of the goroutines. The WaitGroup keeps a counter that tracks how many goroutines are active.

```
wg.Add(3) 
```
- wg.Add(3): This increases the WaitGroup counter by 3. This tells the WaitGroup to expect four Done() calls, so it knows how many tasks (goroutines) will run.

```
// Launch goroutines
go performTask("Task 1")
go performTask("Task 2")
go performTask("Task 3")
```
- go performTask("Task 1"): Each go statement launches performTask as a separate goroutine, running concurrently. Each goroutine will decrement the WaitGroup counter by calling Done when it finishes.

```
wg.Wait() // Wait until all goroutines finish
```

- wg.Wait(): This blocks the main function from proceeding until the WaitGroup counter reaches zero. It ensures that the program waits for all launched goroutines to finish before continuing.


```
Task 2 finished
Task 3 finished
Task 1 finished
All tasks completed.
```

- Here as you see , each tasks executed parallely and task2 and task 3 completed before Task1 completed. 