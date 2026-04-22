package entity

type Item struct {
	Value     string
	ExpiresAt *int64 //ponteiro pq tem valores que nao irao expirar
}

func (i *Item) IsExpired(now int64) bool {
	if i.ExpiresAt == nil {
		return false
	}

	return now < *i.ExpiresAt
}
