# Callback Writer

Simple writer, which evaluates specified callback on each incoming write. Close
callback is also available.

Main use case is to use it with [go-lineflushwriter](https://github.com/reconquest/go-lineflushwriter)
to pass entire lines into callback:

```go
writer := lineflushwriter.New(
    callbackwriter.New(
        file,
        func (line []byte) {
            log.Println(string(line))
        },
        nil,
    ),
    &sync.Mutex{},
    nil,
)

// will write `hello` into file and duplicate it into on-screen log
io.WriteString(writer, `hello\n`)
```
