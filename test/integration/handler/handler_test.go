// +build integration

package integration

import (
	"testing"

	"github.com/larwef/cert-monitor/pkg/cert"
	"github.com/larwef/cert-monitor/pkg/handler"
)

// Runs a live test with a request to Buypass test
func TestHandler(t *testing.T) {
	req := cert.BuypassTestRequest("993884871")

	_, err := handler.Handler(nil, *req)
	if err != nil {
		t.Errorf("Handler returned error: %v", err)
	}
}
