# Line
![line logo](line.png)

Line is a simple CLI utility for reading individual and ranges of lines from
a given file

[![Go Reference](https://pkg.go.dev/badge/github.com/asciifaceman/line.svg)](https://pkg.go.dev/github.com/asciifaceman/line) ![build workflow](https://github.com/asciifaceman/line/actions/workflows/release-build.yml/badge.svg) ![latest PR workflow](https://github.com/asciifaceman/line/actions/workflows/pr-build.yml/badge.svg) ![repo license](https://img.shields.io/github/license/asciifaceman/line.svg) ![downloads](https://img.shields.io/github/release/asciifaceman/line.svg)


# Installation

## Go

```
go install github.com/asciifaceman/line@latest
```

## Non-go

Binaries are provided in https://github.com/asciifaceman/line/releases for download

Throw them in your `$PATH`

# Usage

```
line [-l n -l n-N] filename
```

The primary functionality of line is surfaced via the `-l` flag which defines 
either a singleton of range of lines (such as 5-7). This flag can be repeated
to gather various singletons or ranges. If EOF is reached in the middle of a range line will WARN on stderr but return what it gathered, if EOF is reached before a requested line or range is reached it will error out.

```
$ line -l 2 -l 7-10 main.go
Copyright © 2023 Charles <asciifaceman> Corbett
func main() {
        cmd.Execute()
}
```

```
$ line -l 2 -l 7-15 main.go
Copyright © 2023 Charles <asciifaceman> Corbett
[WARN] EOF reached but did capture lines

func main() {
        cmd.Execute()
}
```

```
$ line -l 20-25  main.go
error: failed to read line range [20 25] from file [main.go]: EOF
```

# TODO

* Implement an os.File and os.Stdin split reader to support piping in
* Work on go [packaging](https://pkg.go.dev/about#adding-a-package) QOL for go and non-go users

# Authors

* [Charles Corbett](https://github.com/asciifaceman)

### Honors

* Naftuli Kay
    * Many thanks and inspiration for the project
* Karl Q
    * Helped me with the github workflows template that I'm reusing here
