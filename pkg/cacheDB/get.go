package cacheDB

func (t *cacheDB) Get(table, key string) (string, error) {
	val, err := t.client.HGet(table, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}
