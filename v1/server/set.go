package server

import (
	"fmt"
	"net/url"
	bolt "github.com/boltdb/bolt"
	fiber "github.com/gofiber/fiber/v2"
	// encryption "github.com/0187773933/encryption/v1/encryption"
)

func ( s *Server ) Set( context *fiber.Ctx ) ( error ) {
	if !s.ValidateContext( context ) {
		return context.Status( fiber.StatusUnauthorized ).JSON( fiber.Map{
			"error": "invalid or missing API key" ,
		})
	}
	key := context.Params( "key" )
	value := context.Params( "value" )
	if key != "" && value != "" {
		value_escaped , _ := url.QueryUnescape( value )
		err := s.DB.Update( func( tx *bolt.Tx ) error {
			b , err := tx.CreateBucketIfNotExists( []byte( "j" ) )
			if err != nil { return err }
			fmt.Println( "SET KEY VALUE" , key , value_escaped )
			return b.Put( []byte( key ) , []byte( value_escaped ) )
		})
		if err != nil {
			return context.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
				"error": err.Error() ,
			})
		}
		return context.JSON( fiber.Map{
			"result": true ,
		})
	}
	body := context.Body()
	if len( body ) == 0 {
		params := context.Queries()
		if len( params ) == 0 {
			return context.Status( fiber.StatusBadRequest ).JSON( fiber.Map{
				"error": "no body and no query parameters" ,
			})
		}
		err := s.DB.Update( func( tx *bolt.Tx ) error {
			b , err := tx.CreateBucketIfNotExists( []byte( "j" ) )
			if err != nil { return err }
			for k , v := range params {
				if err := b.Put( []byte( k ) , []byte( v ) ); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return context.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
				"error": err.Error() ,
			})
		}
		return context.JSON( fiber.Map{
			"result": true ,
		})
	}
	err := s.DB.Update( func( tx *bolt.Tx ) error {
		b , err := tx.CreateBucketIfNotExists( []byte( "j" ) )
		if err != nil { return err }
		return b.Put( []byte( key ) , body )
	})
	if err != nil {
		return context.Status( fiber.StatusInternalServerError ).JSON( fiber.Map{
			"error": err.Error() ,
		})
	}
	return context.JSON( fiber.Map{
		"result": true ,
	})
}