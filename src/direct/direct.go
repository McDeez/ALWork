package direct

// direct.go: A Simple http redirector
// Takes the TargetHost from the config and
// reroutes all reqs to that target host
// if and only if they pass the filters.

import (
	"config"
	"filter"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// We need a transport for our RoundTrip
var myTransport = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
}

func DirectIt(w http.ResponseWriter, r *http.Request) {
	// We need to unescape the queryString to really be
	// able to inspect it.
	unescapedURI, _ := url.QueryUnescape(r.RequestURI)

	// See if the queryString passes the filters
	if !filter.RunFilter(unescapedURI) {
		// queryString did not pass, send back generic "bad request"
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Attack detected in URI:", unescapedURI)
		return
	}

	// Create our redirect
	uri := config.MyConfig.TargetHost + r.RequestURI

	newReq, err := http.NewRequest(r.Method, uri, r.Body)
	if err != nil {
		log.Println("Error in new Req:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := myTransport.RoundTrip(newReq)
	if err != nil {
		log.Println("Err from RoundTrip", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	// Write all the stuff we got
	wHeader := w.Header()
	for key, valueArray := range resp.Header {
		for _, value := range valueArray {
			wHeader.Add(key, value)
		}
	}

	// Read/Writete the rest of the resp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error in new Req:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(body)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
