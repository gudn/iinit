# iinit
Some initialization system

Sometimes we want initialize some non-connected subsystems. For example, we want
create logging and configuration subsystem. And each system wants to define own
CLI flag. But we can call `flag.Parse()` only once.

With `iinit` you can write this code:

```go
// logger.go
func init() {
  pretty := flag.Bool("pretty", false, "enable pretty-printing")
  iinit.SequentialS(
    flag.Parse,
    func() {
      // some work here
    }
  )
}
```

``` go
// config.go
func init() {
  configPath := flag.String("config", "config.json", "path to config file")
  iinit.SequentialS(
    flag.Parse,
    func() {
      // some work here
    }
  )
}
```

It computes graph of operators and execute it with maximum parallelism.
