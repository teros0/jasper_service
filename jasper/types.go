package jasper

type InputControl struct {
	CommonAttributes
	Mandatory   bool `json:"mandatory"`
	ReadOnly    bool `json:"readOnly"`
	Visible     bool `json:"visible"`
	Type        int  `json:"type"`
	DataTypeRef `json:"dataType"`
}

type CommonAttributes struct {
	Label       string `json:"label"`
	Description string `json:"description"`
}

type Reference struct {
	URI string `json:"uri"`
}

type QueryRef struct {
	Reference `json:"queryReference"`
}

type DataSourceRef struct {
	Reference `json:"dataSourceReference"`
}

type InputControlRef struct {
	Reference `json:"inputControlReference"`
}

type DataTypeRef struct {
	Reference `json:"dataTypeReference"`
}

type JRXML struct {
	JRXMLFile `json:"jrxmlFile"`
}

type JRXMLFile struct {
	Label   string `json:"label"`
	Type    string `json:"type"`
	Content string `json:"content"`
}
