package gorm_utils

import (
	"fmt"
	"log"
	"net/url"
	"os"
)

// MakeDsn returns DSN for [gorm.io/driver/postgres.Open]
//
// TBW
func MakeDsn(DB_NAME string) string {
	dsn := "FIXME"
	if os.Getenv("INSTANCE_UNIX_SOCKET") != "" {
		dsn = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", os.Getenv("INSTANCE_UNIX_SOCKET"), "5432", DB_NAME, os.Getenv("DB_USER"), os.Getenv("DB_PASS"))
		log.Printf("dsn: %s\n", dsn)
	} else {
		u, err := url.Parse(os.Getenv("DB_URL"))
		if err != nil {
			log.Fatalf("ERROR: %v\n", err)
		}
		u.User = url.UserPassword(os.Getenv("DB_USER"), os.Getenv("DB_PASS"))

		dsn = u.String() + "/" + DB_NAME
	}

	return dsn
}
