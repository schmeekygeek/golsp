package main

import (
	"bufio"
	"log"
	"lspexample/rpc"
	"os"
)

func main() {
  logger := getLogger("/home/abdul_samad/development/lsptest/log.txt")
  logger.Println("Started")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(rpc.Split)

  for scanner.Scan() {
    msg := scanner.Text()
    handleMessage(logger, msg)
  }
}

func handleMessage(logger *log.Logger, msg any) {
  logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
  logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
  if err != nil {
    panic("not a good file")
  }

  return log.New(logfile, "[golsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
