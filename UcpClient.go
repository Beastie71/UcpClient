package UcpClient

import (

)

type Client struct {
    BaseURL 		*url.URL
    Username		string
    Password		string
 
    httpClient *http.Client
}

type userOrg struct {
  
  "fullName": "string",
  "id": "string",
  "isActive": true,
  "isAdmin": true,
  "isImported": true,
  "isOrg": true,
  "membersCount": 0,
  "name": "string"

}

func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

func (s *Client) AddUserOrg(userOrg *UserOrgInst) error {
	url := fmt.Sprintf(baseURL+"/accounts/", s.Username)
	fmt.Println(url)
	j, err := json.Marshal(todo)
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
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}


