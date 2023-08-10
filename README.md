# Line

Line is a simple CLI utility for reading individual and ranges of lines from
a given file

# Usage

```
line [-l n -l n-N] filename
```

The primary functionality of line is surfaced via the `-l` flag which defines 
either a singleton of range of lines (such as 5-7). This flag can be repeated
to gather various singletons or ranges. If EOF is reached in the middle of a range line will WARN but return what it gathered, if EOF is reached before a requested line or range is reached it will error out.

```
$ line -l 2 -l 9-11 main.go
Copyright © 2023 Charles <asciifaceman> Corbett
func main() {
        cmd.Execute()
}
```

# TODO

* Implement an os.File and os.Stdin split reader to support piping in
* Packaging, builds

# Authors

* Charles <asciifaceman> Corbett

### Honors

* Naftuli Kay
    * Many thanks and inspiration for the project
