> ## Simple Go logging library
> ### This library is my own simple logging library that I use for my personal projects.
> ### Feel free to go through the examples to get started.

![Go version](https://img.shields.io/badge/go_version-as_long_as_it_go-blue)
#### How to install
```shell
go get github.com/DinnieJ/logger
```
Check the sample code in  *_example* folders

You can write your own logger by implementing the Logger interface

Example 
```go
package main

import (
    "github.com/DinnieJ/logger"
)
// This example will be updated, i'm too lazy rn, sorry
type MyLogger struct {
    // your logger implementation
}

func (l *MyLogger) Info(msg string) {
    // your logger implementation
}
```

## License

The project is licensed under the [MIT License](LICENSE).