package httputils

import (
 "net/http"
 "io/ioutil"
 "strings"
 )


func MakeRequest(url string) (error, string){
 // Need to strip \n from passed in string to
 // prevent illegal character error
 resp, err := http.Get(strings.TrimSuffix(url, "\n"))

 if err != nil {
   return err, ""
 }
 defer resp.Body.Close()

 body, err := ioutil.ReadAll(resp.Body)

 return nil, string(body)
}
