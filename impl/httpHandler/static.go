package httpHandler

import (
	"fmt"
	"io/ioutil"
	stdHttp "net/http"
	"strings"
)

func StaticHandler(w stdHttp.ResponseWriter, r *stdHttp.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "js") {
		w.Header().Set("Content-Type", "text/javascript")
	} else {
		w.Header().Set("Content-Type", "text/css")
	}
	data, err := ioutil.ReadFile(path[1:])
	if err != nil {
		fmt.Print(err)
	}
	_, err = w.Write(data)
	if err != nil {
		fmt.Print(err)
	}
}
