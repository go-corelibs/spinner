[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/spinner)
[![codecov](https://codecov.io/gh/go-corelibs/spinner/graph/badge.svg?token=3jksYHZI8p)](https://codecov.io/gh/go-corelibs/spinner)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/spinner)](https://goreportcard.com/report/github.com/go-corelibs/spinner)

# go-corelibs/spinner - unicode character spinner

spinner implements a means of rendering unicode characters that
give a sense that something is happening and that the user is waiting
during that process.

# Installation

``` shell
> go get github.com/go-corelibs/spinner@latest
```

# Examples

## Spinner

``` go
func main() {
    // make a new Spinner with all the defaults and no func
    s := spinner.New(nil, nil)
    symbol := s.Symbol()
    // symbol == spinner.DefaultSymbols[0]
    s.Increment()
    // symbol == spinner.DefaultSymbols[1]
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
