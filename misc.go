package webapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	UNAUTHORIZED = "Unauthorized Call"
	UNKNOWN_ERROR = "Unknown Error"
)

func (c *Client) DoGet (url string)([]byte, error){

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


	status := ExtractResponseCode(resp)

	//fmt.Printf("Status is: %v of type: %T\n",status,status)

	if status != http.StatusOK{
		switch status{
		case http.StatusUnauthorized:
			err = errors.New(UNAUTHORIZED)
		default:
			err = errors.New(UNKNOWN_ERROR)
		}
	}

	return body,err

}


func (c *Client) DoPost (url string, v interface{})([]byte, error){



	jsn, err := json.Marshal(v)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsn))

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	req.Header.Add("Authorization", string("Bearer " + c.token))

	resp, err := c.httpClient.Do(req)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()



	status := ExtractResponseCode(resp)


	if status != http.StatusOK{
		switch status{
		case http.StatusUnauthorized:
			err = errors.New(UNAUTHORIZED)
		default:
			err = errors.New(UNKNOWN_ERROR)
		}
		fmt.Println(err)
		body, _ := ioutil.ReadAll(resp.Body)
		return body, err
	}


	body, err := ioutil.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err)
		return nil, err
	}

	return body,err

}

type Option func (s string) string


func BuildURL(url string, options []Option) string{

	if len(options) > 0 {
		url = url + "?"
		for _, option := range options {
			url = option(url)
			url = url + "&"
		}
		url = strings.TrimRight(url,"&")
	}

	return url
}

func ExtractResponseCode(r *http.Response) int {

	re, _ := regexp.Compile(`\d\d\d`)
	match := re.FindString(r.Status)
	status, _ := strconv.Atoi(match)
	return status

}