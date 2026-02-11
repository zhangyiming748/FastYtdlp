package FastYtdlp

import (
	"os"
	"testing"
)

// go test -v -timeout 0 -run ^TestYtdlp$
func TestYtdlp(t *testing.T) {
	root := "/data"
	//proxy := "http://192.168.5.115:8889"
	proxy:=os.Getenv("PROXY")
	Download(root, proxy,"/data/pornhub.cookie")
}
