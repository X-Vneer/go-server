package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/x-vneer/go-server/internal/handlers"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Server running port: ")
	PORT, readerError := reader.ReadString('\n')
	PORT = strings.TrimSpace(PORT)

	if readerError != nil {
		fmt.Println("Could not read")
		return
	}

	log.SetReportCaller(true)

	r := chi.NewRouter()

	handlers.Handler(r)

	fmt.Printf("Starting GO API service ... at port: %v\n", PORT)

	err := http.ListenAndServe("localhost:"+PORT, r)

	if err != nil {
		log.Error(err)
	}
}
