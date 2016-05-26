package mapper

import (
    "log"
    "github.com/patrickmn/go-cache"
    "depositauthws/dao"
)

// create the cache of mapped values
var c = cache.New( cache.NoExpiration, cache.NoExpiration )

// whet we put in the cache
type Mapping struct {
    SourceValue   string
    MappedValue   string
}

// load the cache
func LoadMappingCache( ) error {

    mapper, err := dao.Database.GetFieldMapperList( )
    if err != nil {
        return err
    }

    for _, m := range mapper {
        key := m.FieldClass + ":" + m.FieldSource
        value := m.FieldMapped
        //log.Printf( "Adding: %s -> %s", key, value )
        c.Set( key, value, cache.NoExpiration )
    }

    log.Printf( "Added %d mappings to cache", len( mapper ) )
    return nil
}

// do the field mapping
func MapField( field_class string, source_value string ) ( string, bool ) {

    // lookup the token in the cache
    token := field_class + ":" + source_value
    hit, found := c.Get( token )
    if found {
        return hit.(string), true
    }

    log.Printf( "WARNING: mapped field not found: %s", token )
    return source_value, false
}