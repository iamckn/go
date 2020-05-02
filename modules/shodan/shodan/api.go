package shodan

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err != nil {
		log.Panicln(err)
	}

	var ret APIInfo
	err = json.NewDecoder(res.Body).Decode(&ret)

	if err != nil {
		return nil, err
	}

	return &ret, nil
}
