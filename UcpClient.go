package UcpClient

import (
	"fmt"
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
)

type Client struct {
    BaseURL 		string
    Username		string
    Password		string
 
//    httpClient *http.Client
}

type userOrg struct {
  
  fullName string
  id string
  isActive bool
  isAdmin bool
  isImported bool
  isOrg bool
  membersCount int
  name string

}

func NewBasicAuthClient(baseurl, username, password string) *Client {
	return &Client{
		BaseURL:  baseurl,
		Username: username,
		Password: password,
	}
}

func (s *Client) AddUserOrg(UserOrgInst userOrg) error {
	url := fmt.Sprintf(s.BaseURL+"/accounts/", s.Username)
	fmt.Println(url)
	j, err := json.Marshal(UserOrgInst)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("ERROR %s", body)
	}
	return body, nil
}

func (s *Client) GetUserOrg(id int, whom string) (*userOrg, error) {
	url := fmt.Sprintf(s.BaseURL+"/accounts/"+whom, s.Username, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data userOrg
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
