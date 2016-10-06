package avatar

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client initializes a new avatar api client
type Client struct {
	Username string
	Password string
}

// Profile ...
type Profile struct {
	Name  string `xml:"Name"`
	Image string `xml:"Image"`
	Valid string `xml:"Valid"`
}

// GetResponse takes in an email and returns an
// avatar api response structure
func (c *Client) GetResponse(email string) (*Profile, error) {
	if c.Username == "" || c.Password == "" {
		return nil, errors.New("Must have valid username and password")
	}
	res := Profile{}

	// "http://www.avatarapi.com/avatar.asmx/GetProfile?email=peter.smith@gmail.com&username=xxxxx&password=xxxxx"

	url := fmt.Sprintf("http://www.avatarapi.com/avatar.asmx/GetProfile?email="+email+"&username=%+v&password=%+v", c.Username, c.Password)
	fmt.Println("URL: ", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Status Code was: %+v", resp.StatusCode)
	}

	fmt.Println("STATUS CODE: ", resp.StatusCode)
	fmt.Println("BODY: ", resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
