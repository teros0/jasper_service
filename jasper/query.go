package jasper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Query struct {
	CommonAttributes
	Value         string `json:"value"`
	Language      string `json:"language"`
	DataSourceRef `json:"dataSource"`
}

func NewQuery(name, value, datasource string) *Query {
	query := &Query{}
	query.Label = fmt.Sprintf("%sQuery", name)
	query.Value = value
	query.Language = "sql"
	query.DataSourceRef.URI = datasource
	return query
}

func (q *Query) createRequest(body io.Reader) (*http.Request, error) {
	method := http.MethodPost
	url := serverURL + resourcesURL + queryStorage
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Query.createRequest -> %s", err)
	}
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/repository.query+json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (q *Query) PostToServer() error {
	client := &http.Client{Timeout: 15}
	jq, err := json.Marshal(q)
	if err != nil {
		return fmt.Errorf("Query.PostToServer -> %s", err)
	}
	rq := bytes.NewReader(jq)
	req, err := q.createRequest(rq)
	if err != nil {
		return fmt.Errorf("Query.PostToServer -> %s", err)
	}
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("Query.PostToServer -> %s", err)
	}
	/* implement some logic here
	switch resp.StatusCode {

	}*/
	return nil
}
