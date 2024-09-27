### Pointers

#### What is a Pointer?

- variables that store value addresses instead of value

#### Why use a Pointer?

- Avoiding unnecessary values<br>
  By default, Go creates a copy when passing values to functions. <br>
  For large data, this might be too much. With **Pinters**, only one value is <br>
  stored in memory and the "address" is passed around.

- Directly mutate values<br>
  The function can directly edit the underlying value - no return value is required.<br>
  => might lead to a more unreadable code. Name your functions wisely!<br>

#### How to use a pointer

- `&` indicate that it is a pointer.

For example...

```go
    age := 24
    agePointer := &age
```

If you hover over `agePointer` variable, you see that its type is `*int` which indicates that it is a pointer.<br>
So you could also do this...

```go
    age := 24
    var agePointer *int
    agePointer = &age
```

#### Dereferencing

In order to use the value behind a pointer, you have to do something called **dereferencing**<br>
by adding an `*` in front of the pointer variable.

```go
    fmt.Println("Age: ", *agePointer)
```

#### A Pointer's Null Value

All values in Go have a so-called "Null Value" - i.e., the value that's set as a default<br>
if no value is assigned to a variable.<br>

For example, the null value of an int variable is `0`. Of a float64, it would be `0.0`. Of a string, it's `""`.<br>
For a pointer, it's `nil` - a special value built-into Go.<br>

nil represents the absence of an address value - i.e., a pointer pointing at no address / no value in memory.
