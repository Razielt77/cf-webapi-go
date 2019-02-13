package webapi

import (
	"encoding/json"
	"fmt"
)

type Pipeline struct {
	Name	string `json:"name"`
	IsPublic 	bool `json:"is_public"`
	ID 			 string `json:"id"`
	Tags 		[]string `json:"tags"`
	LastWorkflow Workflow `json:"last_workflow"`
}


func OptionID (s string) Option{
	return func(url string) string{
		return url+"id="+ s
	}
}

func OptionTag (s string) Option {
	return func(url string) string{
		return url+"labels[tags]="+ s
	}
}

type PipelinePayload struct{
	Pipelines	[]PipelineRaw	`json:"docs"`
	Count		int				`json:"count"`
}

type PipelineRaw struct{
	Metadata 	PipelineMetaData `json:"metadata"`
}

type PipelineMetaData struct{
	Name 		string `json:"name"`
	IsPublic 	bool `json:"isPublic"`
	ID			string `json:"id"`
	Labels 		Labels `json:"labels"`
}

type Labels struct{
	Tags 		[]string `json:"tags"`
}

func (c *Client) PipelinesList(options ...Option) ([]Pipeline, error) {


	var arr []Pipeline

	url:= CF_URL+"pipelines/"


	url = BuildURL(url, options)



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

	//fmt.Printf("Arr size is: %v\nCount is: %v\n",len(pipelines.Pipelines), pipelines.Count)

	for _, pipeline := range pipelines.Pipelines {
		var p = Pipeline{Name:pipeline.Metadata.Name,IsPublic:pipeline.Metadata.IsPublic,ID:pipeline.Metadata.ID,Tags:pipeline.Metadata.Labels.Tags}
		wfarr, err := c.WorkflowList(OptionPipelineID(p.ID),OptionLimit("1"))
		if err != nil{
			fmt.Println(err)
			return nil, err
		}
		if len(wfarr) > 0 {
			p.LastWorkflow = wfarr[0]
		}else{
			p.LastWorkflow = Workflow{Status: "N\\A", CreatedTS: "N\\A", Committer: "N\\A", CommitMsg: "N\\A", CommitUrl: "N\\A", Avatar: "N\\A"}
		}
		arr = append(arr,p)
		//fmt.Printf("Pipeline Name is: %s IsPublic value %v\n",pipeline.Metadata.Name,pipeline.Metadata.IsPublic)
	}

	return arr, nil
}

