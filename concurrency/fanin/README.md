# Fan-In Pattern

__Fan-in__ is a concurrency pattern that allows merging input of several goroutines into one.

For example, we have two __goroutines__ Alice and Bob.

Each goroutine sends messages and we want to read it in one function.

Let's implement it in 2 steps.