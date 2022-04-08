package network

import "testing"

var baseUrl = "http://localhost:1500"

func TestGetAndPost(t *testing.T) {
	Get(baseUrl + "/admin/index")
}
