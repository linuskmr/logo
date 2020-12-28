# log + go = logo

logo is a logger for go.

- Can log filename, line number, struct and function name at which the log method is called

- Can suppress log messages with low log level

- Logging without creating a logger struct first
- Colored output of log level via ANSI-Code
- String and JSON output
- Can write to any `io.Writer`

## Config Options

A new logger can be created via `logo.New(flags)` or `logo.Logger{Property: Value}`. All `bool` values have matching flags, which can be combined via the bitwise or `flagA | flagB`. The `logger.Config(flags)` function can reconfigure an existing logger.

The standard logger used by `logo.Print()` can be reconfigured by using `logo.Config(flags)` or replaced with `logo.Standard = logo.Logo{}`.

| Property   | Type         | Description                                                  | Default                                       |
| ---------- | ------------ | ------------------------------------------------------------ | --------------------------------------------- |
| Date       | `bool`       | Specifies whether the date should be displayed               | `true`                                        |
| Time       | `bool`       | Specifies whether the time should be displayed               | `true`                                        |
| Millis     | `bool`       | Specifies whether the milliseconds should be displayed       | `true`                                        |
| Filename   | `bool`       | Specifies whether the filename and line should be displayed that was logged | `true`                                        |
| Funcname   | `bool`       | Specifies whether the struct and function should be displayed where the logger was called | `true`                                        |
| Json       | `bool`       | Specifies whether the output should be in Json format        | `true`                                        |
| Output     | `io.Writer`  | The `io.Writer` the output of the logger shoud be written to | `os.Stdout`                                   |
| DateFormat | `string`     | The date format of the output (in go's time format)          | YYYY-MM-DD (in go's time format "2006-01-02") |
| TimeFormat | `string`     | The time format of the output (in go's time format)          | HH:MM:SS (in go's time format "15:04:05")     |
| Level      | `logo.Level` | Level is the log level of this logger. The logger only logs messages with a level greater or equal to the log level. By default AllLevels is selected, which logs everything. | `logo.AllLevels`                              |

## Log Levels

The logger logs all log messages with a log level greater or equal its `logger.Level`. The value of the log levels increases towards the bottom.

| Level      | Function | Color  |
| ---------- | -------- | ------ |
| DebugLevel | Debug()  | Green  |
| InfoLevel  | Info()   | Blue   |
| WarnLevel  | Warn()   | Yellow |
| ErrorLevel | Error()  | Red    |
| PrintLevel | Print()  | Normal |
| AllLevels  | -        | -      |

## Example

```go
import "logo"

func main() {
    logo.Print("Logging made easy")
    logo.Info("Different colors")
    logo.Warn("Uh yellow now");
    // PRINT: 2020-12-27 18:21:36.474 example.go:4 examples.main: Logging made easy
    // INFO : 2020-12-27 18:21:36.476 example.go:5 examples.main: Different colors
    // WARN : 2020-12-27 18:21:36.478 example.go:6 examples.main: Uh yellow now
    
    logo.Config(logo.Time | logo.Millis)
    logo.Error("Some config options")
    // WARN : 18:21:36.478: Some config options
    
    logger := logo.New(logo.Filename | logo.Date)
    logger.Debug("A logger object")
    // WARN : 2020-12-27 example.go:6: A logger object
    
    logger = logo.Logger{
        Date:  true,
        Level: WarnLevel,
    }
    logger.Debug("Due to the log level this message is not displayed")
    logger.Warn("But this message does")
    // WARN : 2020-12-27: But this message does
}
```