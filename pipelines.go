package webapi

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

type Pipeline struct {
	Name	string `json:"name"`
}

func (c *Client) PipelinesList() ([]Pipeline, error) {
	var arr []Pipeline


	url:= CF_URL+"pipelines/"

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

	fmt.Printf("Body:\n %s",body)
	return arr, nil
}