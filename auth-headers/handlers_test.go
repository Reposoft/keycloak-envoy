package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGatekeeperHeadersToEnvoyExtAuthHandler(t *testing.T) {
	handler := http.HandlerFunc(GatekeeperHeadersToEnvoyExtAuthHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	client := &http.Client{}

	req, err := http.NewRequest("GET", server.URL, nil)
	// https://www.keycloak.org/docs/latest/securing_apps/index.html#upstream-headers
	req.Header.Add("X-Auth-Subject", `29a2f562-8697-4e10-b78b-3b287d662bca`)
	req.Header.Add("X-Auth-Session-State", `391f1a7d-74dd-4a1c-83e0-6c902fb7101c`)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
	}

	headers := resp.Header

	if len(headers["X-Auth-Subject"]) != 1 || headers["X-Auth-Subject"][0] != "29a2f562-8697-4e10-b78b-3b287d662bca" {
		t.Error("Should copy the X-Auth-Subject header for forwarding")
	}
	if len(headers["X-Auth-Session-State"]) != 1 || headers["X-Auth-Session-State"][0] != "391f1a7d-74dd-4a1c-83e0-6c902fb7101c" {
		t.Error("Should copy the X-Auth-Session-State header for forwarding")
	}

	// we don't really care about body, but let's keep a test around to alert us on changes
	expected := "At path /, sub 29a2f562-8697-4e10-b78b-3b287d662bca session 391f1a7d-74dd-4a1c-83e0-6c902fb7101c"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if expected != string(body) {
		t.Errorf("Expected the message '%s', got '%s'\n", expected, string(body))
	}
}
