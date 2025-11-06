package utils

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"net"
	"encoding/hex"
	"encoding/base64"
	"encoding/binary"
	types "github.com/0187773933/GetSetServer/v1/types"
	encryption "github.com/0187773933/encryption/v1/encryption"
)

func ParseConfig( file_path string ) ( result types.ConfigFile ) {
	file_data , _ := ioutil.ReadFile( file_path )
	err := json.Unmarshal( file_data , &result )
	if err != nil { fmt.Println( err ) }
	return
}

// https://stackoverflow.com/a/28862477
func GetLocalIPAddresses() ( ip_addresses []string ) {
	host , _ := os.Hostname()
	addrs , _ := net.LookupIP( host )
	encountered := make( map[ string ]bool )
	for _ , addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			ip := ipv4.String()
			if !encountered[ ip ] {
				encountered[ ip ] = true
				ip_addresses = append( ip_addresses , ip )
			}
		}
	}
	return
}

func ItoB( v uint64 ) []byte {
    b := make( []byte , 8 )
    binary.BigEndian.PutUint64( b , v )
    return b
}

func GenerateNewKeys() {
	admin_login_url := encryption.GenerateRandomString( 16 )
	admin_prefix := encryption.GenerateRandomString( 6 )
	login_url := encryption.GenerateRandomString( 16 )
	prefix := encryption.GenerateRandomString( 6 )
	// https://github.com/gofiber/fiber/blob/main/middleware/encryptcookie/utils.go#L91
	// https://github.com/0187773933/encryption/blob/master/v1/encryption/encryption.go#L46
	// cookie_secret := fiber_cookie.GenerateKey()
	cookie_secret_bytes := encryption.GenerateRandomBytes( 32 )
	cookie_secret := base64.StdEncoding.EncodeToString( cookie_secret_bytes )
	cookie_secret_message := encryption.GenerateRandomString( 16 )
	admin_cookie_secret_message := encryption.GenerateRandomString( 16 )
	admin_username := encryption.GenerateRandomString( 16 )
	admin_password := encryption.GenerateRandomString( 16 )
	api_key := encryption.GenerateRandomString( 16 )
	bolt_name := encryption.GenerateRandomString( 6 ) + ".db"
	bolt_prefix := encryption.GenerateRandomString( 6 )
	bolt_encryption_key := encryption.GenerateRandomString( 64 )
	redis_prefix := encryption.GenerateRandomString( 6 )
	log_name := encryption.GenerateRandomString( 6 ) + ".db"
	log_key := encryption.GenerateRandomString( 6 )
	log_encryption_key := encryption.GenerateRandomString( 64 )
	kyber_private , kyber_public := encryption.KyberGenerateKeyPair()
	kyber_private_string := hex.EncodeToString( kyber_private[ : ] )
	kyber_public_string := hex.EncodeToString( kyber_public[ : ] )
	fmt.Println( "Generated New Keys :" )
	fmt.Printf( "\tURL - Admin Login === %s\n" , admin_login_url )
	fmt.Printf( "\tURL - Admin Prefix === %s\n" , admin_prefix )
	fmt.Printf( "\tURL - Login === %s\n" , login_url )
	fmt.Printf( "\tURL - Prefix === %s\n" , prefix )
	fmt.Printf( "\tCOOKIE - Secret === %s\n" , cookie_secret )
	fmt.Printf( "\tCOOKIE - USER - Message === %s\n" , cookie_secret_message )
	fmt.Printf( "\tCOOKIE - ADMIN - Message === %s\n" , admin_cookie_secret_message )
	fmt.Printf( "\tCREDS - Admin Username === %s\n" , admin_username )
	fmt.Printf( "\tCREDS - Admin Password === %s\n" , admin_password )
	fmt.Printf( "\tCREDS - API Key === %s\n" , api_key )
	fmt.Printf( "\tAdmin Username === %s\n" , admin_username )
	fmt.Printf( "\tAdmin Password === %s\n" , admin_password )
	fmt.Printf( "\tLOG - Log Name === %s\n" , log_name )
	fmt.Printf( "\tLOG - Log Key === %s\n" , log_key )
	fmt.Printf( "\tLOG - Encryption Key === %s\n" , log_encryption_key )
	fmt.Printf( "\tBOLT - Name === %s\n" , bolt_name )
	fmt.Printf( "\tBOLT - Prefix === %s\n" , bolt_prefix )
	fmt.Printf( "\tBOLT - Encryption Key === %s\n" , bolt_encryption_key )
	fmt.Printf( "\tREDIS - Prefix === %s\n" , redis_prefix )
	fmt.Printf( "\tKYBER - Private Key === %s\n" , kyber_private_string )
	fmt.Printf( "\tKYBER - Public Key === %s\n" , kyber_public_string )
	panic( "Exiting" )
}