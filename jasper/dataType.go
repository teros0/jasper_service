package jasper

type DataType struct {
	CommonAttributes
	Type      int    `json:"type"`
	MaxValue  string `json:"maxValue"`
	StrictMax bool   `json:"strictMax"`
	MinValue  string `json:"minValue"`
	StrictMin bool   `json:"strictMin"`
}
