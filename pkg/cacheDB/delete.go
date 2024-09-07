package cacheDB

func (t *cacheDB) Delete(table string) error {
	return t.client.Del(table).Err()
}
