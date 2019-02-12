package webapi

import (
	"fmt"
)

type Pipeline struct {
	Name	string `json:"name"`
}

type Option func(s string) string {
	retrun string
}

func OptionID (s string) Option{
	return func(url string) string{
		return url+"id:"+ s
	}
}

func (c *Client) PipelinesList(options ...Option) ([]Pipeline, error) {
	var arr []Pipeline

	url:= CF_URL+"pipelines/"

	url = BuildURL(url)

	body, err := c.DoGet(url)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Body:\n %s",body)
	return arr, nil
}

