package writeline

/*
https://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html

- Position the Cursor:
  \033[<L>;<C>H
     Or
  \033[<L>;<C>f
  puts the cursor at line L and column C.
- Move the cursor up N lines:
  \033[<N>A
- Move the cursor down N lines:
  \033[<N>B
- Move the cursor forward N columns:
  \033[<N>C
- Move the cursor backward N columns:
  \033[<N>D

- Clear the screen, move to (0,0):
  \033[2J
- Erase to end of line:
  \033[K

- Save cursor position:
  \033[s
- Restore cursor position:
  \033[u
*/
const (
	putLineColumn         = "\033[%v;%vf"
	moveUpLines           = "\033[%vA"
	moveDownLines         = "\033[%vB"
	moveForwardColumns    = "\033[%vC"
	moveBackwardColumns   = "\033[%vD"
	clearScreen           = "\033[2J"
	eraseToEnd            = "\033[K"
	saveCursorPosition    = "\033[s"
	restoreCursorPosition = "\033[u"

	carriageReturn = "\r"
)
