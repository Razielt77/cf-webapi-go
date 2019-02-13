package webapi

import (
	"encoding/json"
	"fmt"
)

func OptionPipelineID (s string) Option{
	return func(url string) string{
		return url+"pipeline="+ s
	}
}

func OptionLimit (s string) Option{
	return func(url string) string{
		return url+"limit="+ s
	}
}

// The json structure returned from Codefresh Web API
type WorkflowServerResponse struct {
	ServerTimestamp 	string `json:"serverTimestamp"`
	Workflows 			WorkflowsPayload `json:"workflows"`
}

// The array of workflows as returned from Codefresh Web API
type WorkflowsPayload struct {
	WorkflowsArr 	[]Workflow `json:"docs"`
	Total 		int	`json:"total"`
}

type Workflow struct {
	Status string `json:"status"`
	CreatedTS string `json:"created"`
	Committer 	string `json:"userName"`
	CommitMsg 	string `json:"commitMessage"`
	CommitUrl 	string `json:"commitURL"`
	Avatar 		string `json:"avatar"`
}



func (c *Client) WorkflowList (options ...Option) ([]Workflow, error){

	var arr []Workflow

	url:= CF_URL+"workflow/"

	url = BuildURL(url, options)


	body, err := c.DoGet(url)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	workflowsreponse := WorkflowServerResponse{}

	err = json.Unmarshal(body, &workflowsreponse)

	if err !=nil {
		fmt.Println(err)
		return nil, err
	}

	if len(workflowsreponse.Workflows.WorkflowsArr) > 0 {
		//fmt.Printf("Arr size is: %v\n",len(workflowsreponse.Workflows.WorkflowsArr))
		arr = workflowsreponse.Workflows.WorkflowsArr
	}

	return arr,nil
}
