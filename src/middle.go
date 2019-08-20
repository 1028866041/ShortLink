package src

import (
    "net/http"
    "time"
    "log"
)

type Middlerware struct {
}

func (m Middlerware) LoggingHandler(next http.Handler) http.Handler {

    fnc := func(w http.ResponseWriter, r *http.Request){
        t1 := time.Now()
        next.ServeHTTP(w, r)
        t2 := time.Now()
        log.Printf("[%s] %q %v", r.Method, r.URL.String(), t2.Sub(t1))
    }
    return http.HandlerFunc(fnc)
}

func (m Middlerware)RecoverHandler(next http.Handler) http.Handler {
    fnc := func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if  err:= recover(); err!=nil{
                log.Printf("recover from %+v", err)
            }
            http.Error(w, http.StatusText(500), 500)
        }()
        next.ServeHTTP(w, r)
    }
    return http.HandlerFunc(fnc)
}