package domain

const (
	Asc  = "ASC"
	Desc = "DESC"
)

type OrderBy string

type QueryArgs struct {
	Limit   string
	Offset  string
	OrderBy OrderBy
}
