Official Go Doc: [https://go.dev/doc/](https://go.dev/doc/)

## Go Rules

- double qutes ""
- := keyword : used when declaring a variable where the type should be inferred by Go
  - example using `:=` keyword:
    ```go
    helloWorld := "Hello World"
    ```
    example using `=` keyword:
    ```go
    var helloWorld string = "Hello World"
    ```
- func - defining functions

#### Packages

- package main : tells Go which one is main module
- fmt (the Format package): one of many go packages that holds many helpful functions such as Print
- math
- functions that are to be imported Must start with a captalized letter (e.g. Print or Scan)

#### Functions

- main() - tells Go where the execution starts

- fmt.Print() - console log, no line break
- fmt.Println() - console log with line break
- fmt.Printf() - console log formatted string
- fmt.Sprint() - returns string
- fmt.Springf() - returns formatted string
- fmt.Scan() - scans for user input on terminal
  - pointer (&) - points to a variable for storing user input (Limitations - Can't take multiple words input.)
- math.Pow() - power in other languages
