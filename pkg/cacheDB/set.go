package cacheDB

import "encoding/json"

func (t *cacheDB) Set(
	table, key string,
	value interface{},
) error {
	val, errMarshal := json.Marshal(value)

	if errMarshal != nil {
		return errMarshal
	}

	return t.client.HSet(table, key, string(val)).Err()
}
