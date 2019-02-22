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
		fmt.Println(err)
		return nil, err
	}


	err = json.Unmarshal(body, &info)

	if err !=nil {
		fmt.Println(err)
		return nil, err
	}

	for _, account := range info.Accounts {
		if account.Name == info.ActiveAccount {
			account.Token = c.token
		}
	}

	//fmt.Printf("Arr size is: %v\nCount is: %v\n",len(pipelines.Pipelines), pipelines.Count)

	return info, nil
}
