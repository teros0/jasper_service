package jasper

const (
	datasource      = "/datasources/testdb"
	authorization   = "Basic amFzcGVyYWRtaW46amFzcGVyYWRtaW4="
	serverURL       = "http://192.168.57.3:8080/jasperserver/rest_v2"
	resourcesURL    = "/resources"
	exportReportURL = "/reports"
	queryStorage    = "/datasources/queries"
	reportStorage   = "/reports"
)

type Argument struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func PostReport(name, content string) error {
	report := NewReport(name, content, datasource)
	/*sql, err := report.GetSQL()
	if err != nil {
		return fmt.Errorf("while parsing for sql statement -> %s", err)
	}
	query := NewQuery(name, sql, datasource)
	if err = query.PostToServer(); err != nil {
		return fmt.Errorf("PostReport -> while posting query -> %s", err)
	}*/
	if err := report.PostToServer(); err != nil {
		return err
	}
	return nil
}

func ExportReport(name, format string, args []Argument) ([]byte, error) {
	//fmt.Printf("%+v", args)
	url := formatExportURL(name, format, args)
	doc, err := exportFromServer(url)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
