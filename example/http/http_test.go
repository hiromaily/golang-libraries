package http_test

import (
	"encoding/json"
	"os"
	"testing"

	. "github.com/hiromaily/golibs/example/http"
	lg "github.com/hiromaily/golibs/log"
	tu "github.com/hiromaily/golibs/testutil"
)

type MessagesJSON struct {
	ContentType uint8  `json:"contentType"`
	Text        string `json:"text"`
}

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[HTTP]")
}

func setup() {
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()

	code := m.Run()

	teardown()

	os.Exit(code)
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestGetRequest(t *testing.T) {
	//tu.SkipLog(t)
	status, _, err := GetRequestSimple("https://www.google.co.jp/")

	if err != nil {
		t.Fatalf("TestGetRequest[1]: %s", err)
	}
	if status != 200 {
		t.Errorf("TestGetRequest[2]: %d", status)
	}
	//if body != "hoge" {
	//	t.Errorf("TestGetRequest[3]: %s", body)
	//}

}

func TestGetRequest2(t *testing.T) {
	status, _, err := GetRequestWithData("http://www.yahoo.co.jp/")
	if err != nil {
		t.Fatalf("TestGetRequest2[1]: %s", err)
	}
	if status != 200 {
		t.Errorf("TestGetRequest2[2]: %d", status)
	}
	//lg.Debugf("body: %v", body)
}

func TestPostRequest(t *testing.T) {
	tu.SkipLog(t)

	url := "https://www.google.co.jp/"

	message := MessagesJSON{
		ContentType: 1,
		Text:        "something code",
	}

	byteBody, _ := json.Marshal(message)

	status, _, err := PostRequest(url, byteBody)
	if err != nil {
		t.Fatalf("TestPostRequest1: %s", err)
	}
	if status != 200 {
		t.Errorf("TestPostRequest2: %d", status)
	}
	lg.Debugf("byteBody: %v", byteBody)
}
