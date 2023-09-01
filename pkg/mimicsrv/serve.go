package mimicsrv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getMimicRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request\n")

	type Response struct {
		Metadata map[string]string
		Headers  http.Header
	}

	response_metadata_map := make(map[string]string)
	// required_response_fields := ["Host"]

	// for key, value := range required_response_fields {
	// 	response_metadata_map[value] =
	// }
	response_metadata_map["host"] = r.Host
	response_metadata_map["method"] = r.Method
	response_metadata_map["proto"] = r.Proto
	response_metadata_map["remote_addr"] = r.RemoteAddr
	response_metadata_map["request_uri"] = r.RequestURI

	response_summary := Response{
		Metadata: response_metadata_map,
		Headers:  r.Header,
	}

	if r.URL.Query().Has("format") && r.URL.Query().Get("format") == "json" {
		jsonStr, err := json.Marshal(response_summary)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			io.WriteString(w, string(jsonStr))
		}
	} else if r.URL.Query().Has("format") && r.URL.Query().Get("format") == "pretty" {
		io.WriteString(w, "<html><head></head><body><table border=1><tr><th>Header Key</th><th>Header Value</th></tr>")
		for key, value := range r.Header {
			io.WriteString(w, fmt.Sprintf("<tr><td>%s</td><td>%s</td></tr>\n", key, value))
		}
		io.WriteString(w, "</table></body></html>")
	} else {
		io.WriteString(w, "Metadata\n")
		for key, value := range response_summary.Metadata {
			io.WriteString(w, fmt.Sprintf("%s : %s\n", key, value))
		}
		io.WriteString(w, "\nHeaders\n")
		for key, value := range response_summary.Headers {
			io.WriteString(w, fmt.Sprintf("%s : %s\n", key, value))
		}
	}

}

func ServeMimic(port int) {
	http.HandleFunc("/mimic", getMimicRoot)

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("starting server on", addr)
	err := http.ListenAndServe(addr, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server Closed")
	} else if err != nil {
		fmt.Println("Error starting server", err)
		os.Exit(1)
	}
}
