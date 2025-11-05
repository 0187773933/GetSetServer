package server

import (
	// "fmt"
	"encoding/json"
	bolt "github.com/boltdb/bolt"
	fiber "github.com/gofiber/fiber/v2"
	// encryption "github.com/0187773933/encryption/v1/encryption"
)

func ( s *Server ) Get( context *fiber.Ctx ) ( error ) {
	if !s.ValidateContext( context ) {
		return context.Status( fiber.StatusUnauthorized ).JSON( fiber.Map{
			"error": "invalid or missing API key" ,
		})
	}
	key := context.Params( "key" )
	var value interface{}
	s.DB.View( func( tx *bolt.Tx ) error {
		all_bucket := tx.Bucket( []byte( "j" ) )
		if all_bucket == nil { return nil }
		value_bytes := all_bucket.Get( []byte( key ) )
		if value_bytes == nil { return nil }
		decode_err := json.Unmarshal( value_bytes , &value )
		if decode_err != nil {
			value = string( value_bytes )
		}
		return nil
	})
	return context.JSON( fiber.Map{
		key: value ,
	})
}