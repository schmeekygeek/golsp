package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"lspexample/analysis"
	"lspexample/lsp"
	"lspexample/rpc"
	"os"
)

func main() {
	logger := getLogger("/home/abdul_samad/development/lsptest/log.txt")
	logger.Println("Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

  state := analysis.NewState()
  writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
	logger.Printf("Received message with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Can't parse: %s", err)
		}

		logger.Printf(
			"Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)

    msg := lsp.NewInitializeResponse(request.ID)
    writeResponse(writer, msg)

    logger.Println("Sent reply")
  case "textDocument/didOpen":
    var request lsp.DidOpenTextDocumentNotification
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("textDocument/didOpen: %s", err)
    }
    logger.Printf(
      "Opened: %s",
      request.Params.TextDocument.URI,
    )
    state.OpenDocument(
      request.Params.TextDocument.URI,
      request.Params.TextDocument.Text,
    )
  case "textDocument/didChange":
    var request lsp.TextDocumentDidChangeNotification
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("textDocument/didChange: %s", err)
      return
    }
    logger.Printf(
      "Changed: %s",
      request.Params.TextDocument.URI,
    )
    for _, change := range request.Params.ContentChanges {
      state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
    }
  case "textDocument/hover":
    var request lsp.HoverRequest
    if err := json.Unmarshal(contents, &request); err != nil {
      logger.Printf("Can't parse: %s", err)
    }

    response := state.Hover(
      request.ID,
      request.Params.TextDocument.URI,
      request.Params.Position,
    )
    writeResponse(writer, response)
    logger.Println("Sent reply to textDocument/hover")
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(
    filename,
    os.O_CREATE|os.O_TRUNC|os.O_WRONLY,
    0666,
  )
	if err != nil {
		panic("not a good file")
	}

	return log.New(logfile, "[golsp]", log.Ldate|log.Ltime|log.Lshortfile)
}

func writeResponse(writer io.Writer, msg any) {
  reply := rpc.EncodeMessage(msg)
  writer.Write([]byte(reply))
}
