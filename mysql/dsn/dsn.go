package main

import (
	"crypto/rsa"
	"crypto/tls"
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	dsn := "username:password@protocol(address)/dbname?param=value"
	parseDsn(dsn)
}

type Config struct {
	User             string            // Username
	Passwd           string            // Password (requires User)
	Net              string            // Network type
	Addr             string            // Network address (requires Net)
	DBName           string            // Database name
	Params           map[string]string // Connection parameters
	Collation        string            // Connection collation
	Loc              *time.Location    // Location for time.Time values
	MaxAllowedPacket int               // Max packet size allowed
	ServerPubKey     string            // Server public key name
	pubKey           *rsa.PublicKey    // Server public key
	TLSConfig        string            // TLS configuration name
	tls              *tls.Config       // TLS configuration
	Timeout          time.Duration     // Dial timeout
	ReadTimeout      time.Duration     // I/O read timeout
	WriteTimeout     time.Duration     // I/O write timeout

	AllowAllFiles           bool // Allow all files to be used with LOAD DATA LOCAL INFILE
	AllowCleartextPasswords bool // Allows the cleartext client side plugin
	AllowNativePasswords    bool // Allows the native password authentication method
	AllowOldPasswords       bool // Allows the old insecure password method
	CheckConnLiveness       bool // Check connections for liveness before using them
	ClientFoundRows         bool // Return number of matching rows instead of rows changed
	ColumnsWithAlias        bool // Prepend table alias to column names
	InterpolateParams       bool // Interpolate placeholders into query string
	MultiStatements         bool // Allow multiple statements in one query
	ParseTime               bool // Parse time values to time.Time
	RejectReadOnly          bool // Reject read-only connections
}

const preRegExp = `(\w+):(\w+)@(\w+)\((\w+)\)`

func parseDsn(dsn string) {
	// 	"username:password@protocol(address)/dbname?param=value"
	// find last /
	slashs := strings.Split(dsn, "/")
	if len(slashs) == 1 {
		return
	}
	re := regexp.MustCompile(preRegExp)

	result := re.FindAllSubmatch([]byte(dsn), -1)
	username := result[0][1]
	password := result[0][2]
	protocol := result[0][3]
	address := result[0][4]
	config := &Config{
		User:   string(username),
		Passwd: string(password),
		Addr:   string(address),
		Net:    string(protocol),
	}
	fmt.Printf("%+v", config)
	// username := strings.Split(dsn, ":")[0]

}
