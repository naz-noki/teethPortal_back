package cacheDB

func (t *cacheDB) GetAll(table string) ([]string, error) {
	values, errLRange := t.client.HGetAll(table).Result()

	if errLRange != nil {
		return nil, errLRange
	}
	i := 0
	result := make([]string, len(values))

	for _, value := range values {
		result[i] = value
		i++
	}

	return result, nil
}
