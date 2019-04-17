package phpfpm

import (
	"bytes"
	"github.com/tomasen/fcgi_client"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func connectFpm(fpmSock string) *fcgiclient.FCGIClient {
	fcgi, err := fcgiclient.Dial("unix", fpmSock)
	if err != nil {
		log.Fatal("err@sock:", err)
	}
	return fcgi
}

func Get(w http.ResponseWriter, r *http.Request, webRoot string, fpmSock string) {
	fcgi := connectFpm(fpmSock)
	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = webRoot + "/" + r.URL.Path
	env["SERVER_SOFTWARE"] = "go / fcgiclient "
	env["REMOTE_ADDR"] = "127.0.0.1"
	env["QUERY_STRING"] = r.URL.RawQuery
	resp, err := fcgi.Get(env)
	if err != nil {
		log.Fatal("err@resp:", err)
	}
	sendResponse(w, resp)
}

func Post(w http.ResponseWriter, r *http.Request, webRoot string, fpmSock string) {
	fcgi := connectFpm(fpmSock)
	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = webRoot + "/" + r.URL.Path
	r.ParseForm()
	resp, err := fcgi.PostForm(env, r.Form)
	if err != nil {
		log.Fatal("err@resp:", err)
	}
	sendResponse(w, resp)
}

func Delete(w http.ResponseWriter, r *http.Request, webRoot string, fpmSock string) {
	fcgi := connectFpm(fpmSock)
	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = webRoot + "/" + r.URL.Path
	env["SERVER_SOFTWARE"] = "go / fcgiclient "
	env["REMOTE_ADDR"] = "127.0.0.1"
	env["QUERY_STRING"] = r.URL.RawQuery
	env["REQUEST_METHOD"] = "DELETE"
	env["CONTENT_LENGTH"] = "0"
	resp, err := fcgi.Request(env, nil)
	if err != nil {
		log.Fatal("err@resp:", err)
	}
	sendResponse(w, resp)
}

func Put(w http.ResponseWriter, r *http.Request, webRoot string, fpmSock string) {
	fcgi := connectFpm(fpmSock)
	r.ParseForm()
	body := bytes.NewReader([]byte(r.Form.Encode()))
	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = webRoot + "/" + r.URL.Path
	env["REQUEST_METHOD"] = "PUT"
	env["CONTENT_LENGTH"] = strconv.Itoa(body.Len())
	env["CONTENT_TYPE"] = "application/x-www-form-urlencoded"
	resp, err := fcgi.Request(env, body)
	if err != nil {
		log.Fatal("err@resp:", err)
	}
	sendResponse(w, resp)
}

func OtherHttpMethod(w http.ResponseWriter) {
	resp := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("This http method is not supported.")),
	}
	w.WriteHeader(501)
	sendResponse(w, resp)
}

func sendResponse(w http.ResponseWriter, resp *http.Response) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("err@body:", err)
	}
	if resp.Header.Get("Status") != "" {
		statusCode, _ := strconv.Atoi(resp.Header.Get("Status")[:3])
		w.WriteHeader(statusCode)
	}
	w.Write([]byte(content))
}
