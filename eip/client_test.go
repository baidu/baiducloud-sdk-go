package eip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/baidu/baiducloud-sdk-go/bce"
	"github.com/gorilla/mux"
)

var (
	testHTTPServer *httptest.Server
	eipClient      *Client
)

func init() {
	var credentials, _ = bce.NewCredentialsFromFile("../aksk-test.json")

	//var bceConfig = bce.NewConfig(credentials)
	var bceConfig = &bce.Config{
		Credentials: credentials,
		Checksum:    true,
		Timeout:     5 * time.Second,
		Region:      os.Getenv("BOS_REGION"),
	}
	eipClient = NewEIPClient(bceConfig)
	// eipClient.SetDebug(true)
	r := mux.NewRouter()
	r.HandleFunc("/v1/eip", handleGetEips).Methods("GET")
	r.HandleFunc("/v1/eip/{ip}", handleDeleteEip).Methods("DELETE")
	r.HandleFunc("/v1/eip/{ip}", handleUnbindEip).Methods("PUT")
	r.HandleFunc("/v1/eip", handleCreateEip).Methods("POST")
	testHTTPServer = httptest.NewServer(r)
	eipClient.Endpoint = testHTTPServer.URL
}

func handleCreateEip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	args := &CreateEipArgs{}
	json.Unmarshal(body, args)
	if args.Billing.BillingMethod != "ByTraffic" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := `{
    "eip":"180.181.3.133"
}`
	fmt.Fprint(w, response)
}

func handleGetEips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := `{
    "eipList": [
        {
            "name":"eip-xrllt5M-1",
            "eip": "180.181.3.133",
            "status":"binded",
            "instanceType": "BCC",
            "instanceId": "i-IyWRtII7",
            "shareGroupId": "eg-0c31c93a",
            "eipInstanceType": "shared",
            "bandwidthInMbps": 5,
            "paymentTiming":"Prepaid",
            "billingMethod":null,
            "createTime":"2016-03-08T08:13:09Z",
            "expireTime":"2016-04-08T08:13:09Z"
        },
        {
            "name":"eip-scewa1M-1",
            "eip": "180.181.3.134",
            "status":"binded",
            "instanceType": "BCC",
            "instanceId": "i-KjdgweC4",
            "shareGroupId": null,
            "eipInstanceType": "normal",
            "bandwidthInMbps": 1,
            "paymentTiming":"Postpaid",
            "billingMethod":"ByTraffic",
            "createTime":"2016-03-08T08:13:09Z",
            "expireTime":null
        }
    ],
    "marker":"eip-DCB50385",
    "isTruncated": true,
    "nextMarker": "eip-DCB50387",
    "maxKeys": 2
}`
	fmt.Fprint(w, response)
}

func handleDeleteEip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ip := vars["ip"]
	if ip == "180.76.154.83" {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(400)
	}

}

func handleUnbindEip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	act := query["bind"]
	if len(act) > 0 {
		vars := mux.Vars(r)
		ip := vars["ip"]
		if ip == "180.76.247.62" {
			w.WriteHeader(200)
			return
		}
	}
	act = query["unbind"]
	if len(act) > 0 {
		vars := mux.Vars(r)
		ip := vars["ip"]
		if ip == expectUnbindEip.Ip {
			w.WriteHeader(200)
			return
		}
	}

	act = query["resize"]
	if len(act) > 0 {
		vars := mux.Vars(r)
		ip := vars["ip"]
		if ip == expectResizeEip.Ip {
			w.WriteHeader(200)
			return
		}
	}
	w.WriteHeader(400)
}
