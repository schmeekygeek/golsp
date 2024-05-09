package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// We'll specify params later
	// Params ...
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int    `json:"id,omitempty"`
	// Result
	// Error
}

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}
