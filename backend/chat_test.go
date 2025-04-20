package main

// import (
// 	"net/http"
// 	"net/url"
// 	"strings"
// 	"testing"

// 	"github.com/gorilla/websocket"
// )

// func mustParseURL(raw string) *url.URL {
// 	u, err := url.Parse(raw)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return u
// }

// func TestWebSocket(t *testing.T) {
// 	// Log in and capture cookies
// 	t.Run("Loging For Posts Test", ts.login)

// 	// get login cookies
// 	cookies := ts.Client().Jar.Cookies(mustParseURL(ts.URL))

// 	// build request header with cookies
// 	header := http.Header{}
// 	for _, c := range cookies {
// 		header.Add("Cookie", c.Name+"="+c.Value)
// 	}
// 	header.Set("Origin", "http://localhost:3000")

// 	u := "ws" + strings.TrimPrefix(ts.URL, "http") + "/api/ws"
// 	dialer := websocket.Dialer{}
// 	conn, _, err := dialer.Dial(u, header)
// 	if err != nil {
// 		t.Fatalf("WebSocket connect failed: %v", err)
// 	}
// 	defer conn.Close()
// }
