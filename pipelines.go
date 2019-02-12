package webapi

import (
	"fmt"
)

type Pipeline struct {
	Name	string `json:"name"`
}

func (c *Client) PipelinesList() ([]Pipeline, error) {
	var arr []Pipeline

	url:= CF_URL+"pipelines/"

	body, err := c.DoGet(url)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Body:\n %s",body)
	return arr, nil
}