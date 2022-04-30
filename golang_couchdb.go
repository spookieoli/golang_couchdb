package golangcouchdb

// Type for Connection to Couchdb
type CouchDBAPI struct {
	Url               string
	Username          string
	Passwort          string
	clientMaxWaitTime int64
}
