package types

type Pagination struct {
	Offset         int
	Limit          int
	OrderBy        string
	OrderDirection string
}
