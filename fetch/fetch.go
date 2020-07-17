package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch() {
	for _, url := range os.Args[1:] {

		resp, errResp := http.Get(url)
		if errResp != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", errResp)
			os.Exit(1)
		}

		body, errRead := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if errRead != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s: %v\n", url, errRead)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
