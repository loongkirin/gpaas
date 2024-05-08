package dataquery

type DbQueryWhere struct {
	QueryFilters []DbQueryFilter `json:"query_filters"`
	Connector    string          `json:"connector"`
}

func NewDbQueryWhere(filters []DbQueryFilter, connector string) DbQueryWhere {
	return DbQueryWhere{
		QueryFilters: filters,
		Connector:    connector,
	}
}
