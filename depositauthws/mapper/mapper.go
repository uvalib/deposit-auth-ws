package mapper

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/uvalib/deposit-auth-ws/depositauthws/dao"
	"github.com/uvalib/deposit-auth-ws/depositauthws/logger"
)

// create the cache of mapped values
var c = cache.New(cache.NoExpiration, cache.NoExpiration)

//
// LoadMappingCache -- load the mapping cache
//
func LoadMappingCache() error {

	mapper, err := dao.Store.GetFieldMapperList()
	if err != nil {
		return err
	}

	for _, m := range mapper {
		key := m.FieldClass + ":" + m.FieldSource
		value := m.FieldMapped
		//log.Printf( "Adding: %s -> %s", key, value )
		c.Set(key, value, cache.NoExpiration)
	}

	logger.Log(fmt.Sprintf("Added %d mappings to cache", len(mapper)))
	return nil
}

//
// MapField -- do the field mapping
//
func MapField(fieldClass string, sourceValue string) (string, bool) {

	// lookup the token in the cache
	token := fieldClass + ":" + sourceValue
	hit, found := c.Get(token)
	if found {
		return hit.(string), true
	}

	logger.Log(fmt.Sprintf("WARNING: mapped field not found: %s", token))
	return sourceValue, false
}

//
// end of file
//