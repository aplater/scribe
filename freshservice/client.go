package freshservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/payvision-development/scribe/httpclient"
)

// Freshservice struct
type Freshservice struct {
	URL    string
	APIKey string
}

// NewClient func
func NewClient(url string, apikey string) *Freshservice {
	fs := new(Freshservice)
	fs.URL = url
	fs.APIKey = apikey
	return fs
}

// CreateChange func
func (fs *Freshservice) CreateChange(c *Change) (*ItilChange, error) {
	url := fmt.Sprintf(fs.URL + "/itil/changes.json")

	change, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(change))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(fs.APIKey, "X")
	res, err := httpclient.DoRequest(req)

	var response ItilChange

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// UpdateChangeStatus func
func (fs *Freshservice) UpdateChangeStatus(change int64, status Status) (*ItilChange, error) {

	url := fmt.Sprintf(fs.URL + "/itil/changes/" + strconv.FormatInt(change, 10) + ".json")

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(`{"itil_change":{"status":"`+strconv.Itoa(int(uint8(status)))+`"}}`)))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(fs.APIKey, "X")
	res, err := httpclient.DoRequest(req)

	var response ItilChange

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// AddChangeNote func
func (fs *Freshservice) AddChangeNote(change int64, note string) (*Note, error) {
	url := fmt.Sprintf(fs.URL + "/itil/changes/" + strconv.FormatInt(change, 10) + "/notes.json")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(`{"itil_note": {"body":"`+note+`"}}`)))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(fs.APIKey, "X")
	res, err := httpclient.DoRequest(req)

	var response Note

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
