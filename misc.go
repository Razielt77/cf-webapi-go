package webapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) DoGet (url string)([]byte, error){

	req, err := http.NewRequest("GET", url, nil)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("Authorization", string("Bearer " + c.token))
	resp, err := c.httpClient.Do(req)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	return body,err

}

func BuildURL(url string, options ...Option) string{
	if len(options) > 0 {
		url = url + "?"
	}
	for _, option := range options {
		url = option(url)
		url = url + "&"
	}
	return
}