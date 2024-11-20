# Exercise: Fibonacci closure
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

## Fibonacci Function
This Go program implements a function `fibonacci()` that returns a closure, which generates the Fibonacci sequence. Each time the closure is called, it returns the next number in the Fibonacci sequence.

### How It Works
1. **Fibonacci Function:**
    - The `fibonacci()` function defines two variables `a` and `b`, initialized to 0 and 1, which represent the first two numbers in the Fibonacci sequence.
    - It then returns a closure function that:
      - Updates the values of `a` and `b` by setting `a` to `b` and `b` to the sum of `a` and `b`.
      - Returns the current value of `a`, which is the next number in the Fibonacci sequence.


2. **Using the Closure:**
    - In the `main()` function, we call `fibonacci()` to obtain the closure and assign it to the variable `f`.
    - Then, a loop runs 10 times, calling `f()` on each iteration to print the next Fibonacci number.

### Example:
```go
package main

import "fmt"

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())  // Prints the first 10 Fibonacci numbers
    }
}
```
**Output:**
```
0
1
1
2
3
5
8
13
21
34
```