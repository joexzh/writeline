# writeline

Simple Go lib for overwrite stdout lines. 
It maintains new lines and overwrites to them.

## methods:
```go
// new writeline creates and maintains 10 new lines.
// It holds a default flush timer with 200 milliseconds duration   
wl, err := NewWithStdout(10)

// reset the flush timer duration
wl.SetFlushDuration(100 * time.Millisecond)

// Flush all changes immediately
wl.Flush()

// overwrite line 0
wl.WriteLine(0, "Hello world")

// overwrite line 9 and flush all changes immediately
wl.WriteLine(9, "Hello world").Flush()

// overwrite the last line
wl.WriteLastLine("hi")

// create a new line at the bottom and write to it
wl.NewLine("foo")

// overwrite the line just created
wl.WriteLastLine(10, "bar")

// total maintained lines
wl.Lines() // 11

// close writeline
wl.Close()
```

## Caution

Don't put any position control to the string.
Such as `\n`, `\033[K`...

From `New()` through `Close()`, for all goroutines, don't use other write functions that effect the result of `Stdout`.
Such as `fmt.Print()`.