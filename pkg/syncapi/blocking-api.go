package syncapi

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tomwerneruk/aws-kata-tool/pkg/cpuload"
)

type microblog struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

var messages = []microblog{
	{Username: "wewjewew", Message: "klqkwekwqekqelkw"},
	{Username: "lkwlewm", Message: "lwewk l2k3 l23k 2l3k232n3m"},
}

func getMessages(c *gin.Context) {
	// Create some fake API request delay
	cpuload.RunCPULoad(1000, 500)
	c.IndentedJSON(http.StatusOK, messages)
}

func getMimicResponse(c *gin.Context) {
	cpuload.RunCPULoad(1000, 250)
	//mimicsrv.MimicResponse()
	type Response struct {
		Metadata map[string]string
		Headers  http.Header
	}

	response_metadata_map := make(map[string]string)

	response_metadata_map["host"] = c.Request.Host
	response_metadata_map["method"] = c.Request.Method
	response_metadata_map["proto"] = c.Request.Proto
	response_metadata_map["remote_addr"] = c.Request.RemoteAddr
	response_metadata_map["request_uri"] = c.Request.RequestURI

	response_summary := Response{
		Metadata: response_metadata_map,
		Headers:  c.Request.Header,
	}

	mimic_format := c.Query("format")
	if mimic_format == "json" {
		c.JSON(http.StatusOK, response_summary)
	} else if mimic_format == "prettyJson" {
		c.IndentedJSON(http.StatusOK, response_summary)
	} else if mimic_format == "pretty" {
		c.HTML(http.StatusOK, "mimic-response.html", response_summary)
	} else {
		c.JSON(http.StatusOK, response_summary)
	}
}

func ServeMain(port int) {
	addr := fmt.Sprintf(":%d", port)
	fmt.Println("starting server on", addr)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/restapi/messages", getMessages)
	router.GET("/mimic", getMimicResponse)

	router.Run(addr)
}
