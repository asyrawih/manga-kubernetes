package utils

func WithOrderBy(mainQuery string) func(column string, sort string) string {
	return func(column string, sort string) string {
		if sort == "" {
			sort = "DESC"
		}
		mainQuery += " ORDER BY " + column + " " + sort
		return mainQuery
	}
}

func WithLimit(mainQuery string) func(limit string, offset string) string {
	return func(limit string, offset string) string {
		mainQuery += " LIMIT " + limit + " OFFSET " + limit
		return mainQuery
	}
}
