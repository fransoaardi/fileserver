package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	h := http.DefaultServeMux
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// only gets POST
		if r.Method != http.MethodPost {
			fmt.Println("only POST method is allowed")
			return
		}

		// get file from request
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		f, err := os.OpenFile("./"+header.Filename, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
	})

	server := http.Server{
		Addr:    "0.0.0.0:5000",
		Handler: h,
	}

	fmt.Println("server started running on 0.0.0.0:5000")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server terminated")
}
