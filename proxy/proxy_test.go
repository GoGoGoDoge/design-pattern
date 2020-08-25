package proxy_test

import (
	"proxy"
	"testing"
)

func TestProxy(t *testing.T) {
	myapp := &proxy.App{}
	myProxyApp := proxy.NewAppProxy(myapp)
	if myProxyApp.Do() != "Before:App is running#After" {
		t.Fatal("Incorrect proxy implementation")
	}
}
