package webapi

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct{
	Name 			string `json:"userName"`
	DefaultAccount	int `json:"defaultAccount"`
	Accounts 		[]AccountInfo `json:"account"`
}

type AccountInfo struct{
	Name 			string `json:"name"`
}


func (c *Client) UserInfo(options ...Option) (info *UserInfo, err error) {


	info = &UserInfo{}

	url:= CF_URL+"user/"


	url = BuildURL(url, options)



	body, err := c.DoGet(url)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}


	err = json.Unmarshal(body, &info)

	if err !=nil {
		fmt.Println(err)
		return nil, err
	}

	//fmt.Printf("Arr size is: %v\nCount is: %v\n",len(pipelines.Pipelines), pipelines.Count)

	return info, nil
}
