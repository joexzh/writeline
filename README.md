# writeline

Simple Go lib for overwrite stdout lines. 
It maintains new lines and overwrites to them.

## methods:


```go
import "github.com/joexzh/writeline"

func main() {
    // new writeline creates and maintains 10 new lines.
    // It holds a default flush timer with 200 milliseconds duration   
    lw, err := writeline.New(10)
    
    // reset the flush timer duration
    lw.SetFlushDuration(100 * time.Millisecond)
    
    // Flush all changes immediately
    lw.Flush()
    
    // overwrite line 0
    lw.WriteLine(0, writeline.Style(writeline.Bold+writeline.Green, "hi"))
    
    // overwrite line 9
    lw.WriteLine(9, "Hello world")
    lw.Flush()
    
    // overwrite the last line
    lineNum, err := lw.WriteLastLine("hi")
    
    // create a new line at the bottom and write to it
    newLineNum, err := lw.WriteNewLine("foo")
    
    // total maintained lines
    lw.Lines() // 11
    
    // close writeline
    lw.Close()
}

```

## Caution

Don't put any position control to the string.
Such as `\n`, `\033[K`...

From `New()` through `Close()`, for all goroutines, don't use other write functions that effect the result of `Stdout`.
Such as `fmt.Print()`.