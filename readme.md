
# go-comment-macros

  Go comment macros useful for injecting logging, tracing, debugging, or stats related code.

  This is a port of [node-comment-macros](https://github.com/visionmedia/node-comment-macros) to Go.

  Stop writing code like this:

```go
func (u *User) Save() {
	fmt.Println("start saving user")
	metrics.Start("saving user")
	Db.save(u)
	metrics.End("saving user")
	fmt.Println("end saving user")
}
```

## About

  I wouldn't recommend this at the library level, normally even at the application level I wouldn't recommend it, but some of our projects require a _lot_ of logging and metrics, so this helps cut the clutter.

## Installation

```
$ go get github.com/olivoil/go-comment-macros
```

## Example

 You can specify a `prefix` which defaults to ":",
 telling comment-macros what is and what is not a macro.

```go
func (u *User) Save() {
	//: start saving user
	user, err := Db.save(u)
	//: end saving user
}
```

 Then you can map them to new values. Note that if you
 return an empty string `""` then nothing will happen,
 so if you have no plugins these comments will simply
 be removed.

```go
import "github.com/olivoil/go-comment-macros"
import "strings"
import "fmt"

m := macros.NewMacro()

m.Register(func(label string) string {
  return "fmt.Println(\"" + label + "\")"
})

m.Register(func(label string) string {
  if 0 == strings.Index(label, "start ") {
    return "metrics.Start(\"" + strings.Replace(label, "start ", "", 1) + "\")"
  }

  if 0 == strings.Index(label, "end ") {
    return "metrics.End(\"" + strings.Replace(label, "end ", "", 1) + "\")"
  }

  return ""
})

s := m.Process(str)
fmt.Println(s)
```

  Yielding:

```go
func (u *User) Save() {
	// normal comment
	fmt.Println("start saving user")
	metrics.Start("saving user")
	user, err := Db.save(u)
	fmt.Println("end saving user")
	metrics.End("saving user")
}
```

# License

  MIT
