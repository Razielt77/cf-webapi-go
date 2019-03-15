package webapi

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct{
	Name 			string `json:"userName"`
	ActiveAccount	string `json:"activeAccountName"`
	Accounts 		[]AccountInfo `json:"account"`
	UserData 		UserDataInfo `json:"user_data"`
}

type AccountInfo struct{
	Name 			string `json:"name"`
	Token 			string `json:"token"`
	UserName		string `json:"user_name"`
}

type UserDataInfo struct{
	Image 			string `json:"image"`
}


func (c *Client) UserInfo(options ...Option) (info *UserInfo, err error) {


	info = &UserInfo{}

	url:= CF_URL+"user/"


	url = BuildURL(url, options)



	body, err := c.DoGet(url)

	if err != nil{
		return nil, err
	}


	err = json.Unmarshal(body, &info)

	if err !=nil {
		fmt.Println(err)
		return nil, err
	}

	for  i , _ := range info.Accounts {
		if info.Accounts[i].Name == info.ActiveAccount {
			info.Accounts[i].Token = c.token
			info.Accounts[i].UserName = info.Name
		}
	}

	//fmt.Printf("Arr size is: %v\nCount is: %v\n",len(pipelines.Pipelines), pipelines.Count)

	return info, nil
}
