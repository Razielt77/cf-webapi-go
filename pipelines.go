package webapi

import (
	"encoding/json"
	"fmt"
)

type Pipeline struct {
	Name	string `json:"name"`
}

type Option func (s string) string

func OptionID (s string) Option{
	return func(url string) string{
		return url+"id="+ s
	}
}

type PipelinePayload struct{
	Pipelines	[]PipelineRaw	`json:"docs"`
	Count		int				`json:"count"`
}

type PipelineRaw struct{

}

func (c *Client) PipelinesList(options ...Option) ([]Pipeline, error) {
	var arr []Pipeline

	url:= CF_URL+"pipelines/"

	fmt.Printf("url is: %s\noption length is %v\n", url, len(options))

	url = BuildURL(url, options)

	fmt.Printf("url is: %s\n", url)

	body, err := c.DoGet(url)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	pipelines := PipelinePayload{}

	err = json.Unmarshal(body, &pipelines)

	if err !=nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Arr size is: %v\n Count is: %v\n",len(pipelines.Pipelines), pipelines.Count)

	//fmt.Printf("Body:\n %s\n",body)
	return arr, nil
}

