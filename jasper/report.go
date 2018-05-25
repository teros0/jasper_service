package jasper

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Report struct {
	CommonAttributes
	//QueryRef      `json:"query"`
	DataSourceRef `json:"dataSource"`
	//InputControls []InputControlRef `json:"inputControls"`
	JRXML `json:"jrxml"`
}

type JasperReport struct {
	SQL string `xml:"queryString"`
}

func NewReport(name, content, datasource string) *Report {
	report := &Report{}
	report.Label = name
	report.DataSourceRef.URI = datasource
	report.JRXML.Label = fmt.Sprintf("%sjrxml", name)
	report.JRXML.Type = "jrxml"
	report.JRXML.Content = content
	return report
}

func (r *Report) GetSQL() (string, error) {
	var jr JasperReport
	bts, err := base64.StdEncoding.DecodeString(r.Content)
	if err != nil {
		return "", fmt.Errorf("while decoding content -> %s", err)
	}
	if err := xml.Unmarshal(bts, &jr); err != nil {
		return "", fmt.Errorf("while unmarshalling xml -> %s", err)
	}
	jr.SQL = strings.TrimLeft(jr.SQL, "\n \t")
	return jr.SQL, nil
}

func (r *Report) createRequest(body io.Reader) (*http.Request, error) {
	method := http.MethodPost
	url := serverURL + resourcesURL + reportStorage
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Report.createRequest -> %s", err)
	}
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/repository.reportUnit+json")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (r *Report) PostToServer() error {
	client := http.Client{Timeout: 15 * time.Second}
	bs, err := json.Marshal(r)
	if err != nil {
		return fmt.Errorf("Report.PostToServer -> %s", err)
	}
	fmt.Println("Create report json", string(bs))
	reader := bytes.NewReader(bs)
	req, err := r.createRequest(reader)
	if err != nil {
		return fmt.Errorf("Report.createRequest -> %s", err)
	}
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("Query.PostToServer -> %s", err)
	}
	return nil
}

func exportFromServer(url string) ([]byte, error) {
	client := http.Client{Timeout: 15 * time.Second}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Basic amFzcGVyYWRtaW46amFzcGVyYWRtaW4=")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return bs, nil
}

func formatExportURL(name, format string, args []Argument) string {
	url := fmt.Sprintf("%s%s%s/%s.%s?", serverURL, exportReportURL, reportStorage, name, format)
	argStr := ""
	for i, v := range args {
		argStr += fmt.Sprintf("%s=%v", v.Name, v.Value)
		if i != len(args)-1 {
			argStr += "&"
		}
	}
	url += argStr
	fmt.Println(url)
	return url
}
