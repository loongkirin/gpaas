package dataquery

type DbQueryFilter struct {
	FieldName       string        `json:"field_name"`
	FilterValues    []interface{} `json:"filter_values"`
	FilterOperation string        `json:"filter_operation"`
	Connector       string        `json:"connector"`
	FieldType       string        `json:"field_type"`
}

func NewDbQueryFilter(field string, values []interface{}, op string, ft string) DbQueryFilter {
	return DbQueryFilter{
		FieldName:       field,
		FilterValues:    values,
		FilterOperation: op,
		Connector:       "AND",
		FieldType:       ft,
	}
}
