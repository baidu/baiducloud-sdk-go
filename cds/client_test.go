package cds

import (
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/baidu/baiducloud-sdk-go/bce"
	"github.com/gorilla/mux"
)

var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

//var bceConfig = bce.NewConfig(credentials)
var bceConfig = &bce.Config{
	Credentials: credentials,
	Checksum:    true,
	Region:      os.Getenv("Region"),
}
var bccClient = NewClient(bceConfig)

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
	cdsClient = NewClient(bceConfig)
	r := mux.NewRouter()
	// loadbalancer
	r.HandleFunc("/v2/volume", handleGetVolumes).Methods("GET")
	r.HandleFunc("/v2/volume", handleCreateVolumes).Methods("POST")
	r.HandleFunc("/v2/volume/{volumeid}", handleDeleteVolume).Methods("DELETE")
	r.HandleFunc("/v2/volume/{volumeid}", handleAttachVolume).Methods("PUT")
	r.HandleFunc("/v2/volume/{volumeid}", handleDetachVolume).Methods("PUT")
	r.HandleFunc("/v2/volume/{volumeid}", handleDescribeVolume).Methods("GET")

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
