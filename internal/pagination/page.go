package pagination

type Page[T any] struct {
	Items  []T `json:"items"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

func Build[T any](items []T, offset, limit int) Page[T] {
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 50
	}
	total := len(items)
	if offset > total {
		offset = total
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return Page[T]{Items: items[offset:end], Offset: offset, Limit: limit, Total: total}
}
