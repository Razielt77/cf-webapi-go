package webapi

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetBuild (id string) (*Workflow, error){


	url:= CF_URL+"builds/"+id

	body, err := c.DoGet(url)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	workflow := &Workflow{}

	err = json.Unmarshal(body, workflow)

	if err !=nil {
		fmt.Println(err)
		return nil, err
	}


	return workflow,nil
}

