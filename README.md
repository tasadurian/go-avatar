# go-avatar [![Go Report Card](https://goreportcard.com/badge/github.com/TheTommyTwitch/go-avatar)](https://goreportcard.com/report/github.com/TheTommyTwitch/go-avatar)
Go wrapper for the [avatar api](http://www.avatarapi.com/).

## Usage

```go get github.com/TheTommyTwitch/go-avatar```

```
func main() {
  c := new(Client)
  c.Username = "username"
  c.Password = "password"

  response, err := c.GetResponse("peter.smith@gmail.com")
  if err != nil {
    panic(err)
  }
  fmt.Printf("%+v", response)
}
```

Return structure:
```
type Profile struct {
	Name  string // name
	Image string // url to profile image
	Valid string // valid: true or false
}
```
