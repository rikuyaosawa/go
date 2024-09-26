### Error Handling in Go

- It works a bit differently than other languages.
- In Go, **functions are written such that errors don't crash the application**.

```go
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}
	return balance, nil
```

- panic("message") method: exit out of program
