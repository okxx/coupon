package handle

import "net/http"

// welcome
func Welcome(writer http.ResponseWriter, request *http.Request) {

	_,_=writer.Write([]byte("welcome coupon"))
}