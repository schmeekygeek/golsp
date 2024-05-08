package rpc_test

import (
	"lspexample/rpc"
	"testing"
)

type EncodingExample struct {
  Testing bool `json:"testing"`
}

func TestEncode(t *testing.T) {
  expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
  got := rpc.EncodeMessage(EncodingExample{Testing:true})

  if expected != got {
    t.Fatalf("Expected: %s, Got: %s", expected, got)
  }
}

func TestDecode(t *testing.T) {
  incomingMessage := "Content-Length: 18\r\n\r\n{\"Method\":\"hello\"}"
  method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
  contentLength := len(content)

  if err != nil {
    t.Fatal(err)
  }

  if contentLength != 18 {
    t.Fatalf("Expected: %d, Got: %d", 18, contentLength)
  }

  if method != "hello" {
    t.Fatalf("Expected: %s, Got, %s", "hello", method)
  }
}
