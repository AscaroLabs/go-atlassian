package main

import (
	"context"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		issueKeyOrID = "DESK-3"
		slaMetricID  = 1
	)

	sla, response, err := atlassian.ServiceManagement.Request.SLA.Get(context.Background(), issueKeyOrID, slaMetricID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", response.Bytes.String())
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.Code)
	log.Println("HTTP Endpoint Used", response.Endpoint)
	log.Println("SLA ID", sla.ID)
	log.Println("SLA Name", sla.Name)

}
