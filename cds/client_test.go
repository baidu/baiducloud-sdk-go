package cds

import (
	"net/http"
	"os"
	"github.com/baidu/baiducloud-sdk-go/bce"
	"net/http/httptest"
	"github.com/gorilla/mux"
)

var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

//var bceConfig = bce.NewConfig(credentials)
var bceConfig = &bce.Config{
	Credentials: credentials,
	Checksum:    true,
	Region:      os.Getenv("BOS_REGION"),
}
var bccConfig = NewConfig(bceConfig)
var bccClient = NewClient(bccConfig)

var (
	testHTTPServer *httptest.Server
	cdsClient      *Client
)

func init() {
	var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

	//var bceConfig = bce.NewConfig(credentials)
	var bceConfig = &bce.Config{
		Credentials: credentials,
		Checksum:    true,
		Region:      os.Getenv("CDS_REGION"),
	}
	var bccConfig = NewConfig(bceConfig)
	cdsClient = NewClient(bccConfig)
	r := mux.NewRouter()
	// loadbalancer
	r.HandleFunc("/v1/blb", handleGetVolumes).Methods("GET")
	r.HandleFunc("/v1/blb", handleCreateVolumes).Methods("POST")
	r.HandleFunc("/v1/blb/{blbid}", handleDeleteVolume).Methods("DELETE")
	r.HandleFunc("/v1/blb/{blbid}", handleAttachVolume).Methods("PUT")
	r.HandleFunc("/v1/blb", handleDetachVolume).Methods("POST")
	r.HandleFunc("/v1/blb", handleDescribeVolume).Methods("POST")

	// start
	testHTTPServer = httptest.NewServer(r)
	cdsClient.Endpoint = testHTTPServer.URL
}

//Todo: add mock functions
func handleGetVolumes(w http.ResponseWriter, r *http.Request) {
}

func handleCreateVolumes(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteVolume(w http.ResponseWriter, r *http.Request) {
}

func handleAttachVolume(w http.ResponseWriter, r *http.Request) {
}

func handleDetachVolume(w http.ResponseWriter, r *http.Request) {
}

func handleDescribeVolume(w http.ResponseWriter, r *http.Request) {
}
