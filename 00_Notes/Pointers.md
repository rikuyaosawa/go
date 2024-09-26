### Pointers

What is a Pointer?

- variables that store value addresses instead of value

Why use a Pointer?

- Avoiding unnecessary values<br>
  By default, Go creates a copy when passing values to functions. <br>
  For large data, this might be too much. With **Pinters**, only one value is <br>
  stored in memory and the "address" is passed around.

- Directly mutate values<br>
  The function can directly edit the underlying value - no return value is required.<br>
  => might lead to a more unreadable code. Name your functions wisely!<br>
