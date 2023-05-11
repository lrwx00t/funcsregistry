# funcsregistry
Demo for registering funcs in Go

## Usage

Adding a new function to register is as simple as adding a new struct in the `registry` and implement the `Register` interface.

```bash
❯ go run main.go
demo1 registered.. {}
demo2 registered.. {}
demo3 registered.. {}
demo4 registered.. {}
```