package src

import (
    "log"
    "github.com/gorilla/mux"
    "net/http"
    "encoding/json"
    "fmt"
    //"github.com/justinas/alice"
)

type ShortLink struct{
    Router *mux.Router
    Middlerwares *Middlerware
}

type shortReq struct{
    url string `json:"url" validate:"nonezero"`
    expire int64 `json:"expire" validate:"min=0"`
}

type shortResp struct{
    shortLink string `json:"shortlink"`
}

func (sl *ShortLink) Initialize(){
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    sl.Router = mux.NewRouter()
    sl.Middlerwares = &Middlerware{}
    sl.InitilizeRoutes()
}

func (sl *ShortLink) InitilizeRoutes(){
    sl.Router.HandleFunc("/create", sl.createLink).Methods("POST")
    sl.Router.HandleFunc("/get", sl.getInfo).Methods("GET")
    sl.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}}", sl.redirect).Methods("GET")
    /* md := alice.New(sl.Middlerwares.LoggingHandler, sl.Middlerwares.RecoverHandler)
    sl.Router.Handle("/create", md.ThenFunc(sl.createLink)).Methods("POST")
    sl.Router.Handle("/get", md.ThenFunc(sl.getInfo)).Methods("GET")
    sl.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,11}}", md.ThenFunc(sl.redirect)).Methods("GET")*/
}

func (sl *ShortLink) createLink(w http.ResponseWriter, r *http.Request){
    var req shortReq
    if err:= json.NewDecoder(r.Body).Decode(&req); err!=nil{
        responsewithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("parse para error", r.Body)})
        return
    }
    /*if err:= validator.Validate(r); err != nil{
        responsewithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("validate para error", r.Body)})
        return
    }*/
    defer r.Body.Close()
    fmt.Print("%v\n", req)
}

func (sl *ShortLink) getInfo(w http.ResponseWriter, r *http.Request){
    val := r.URL.Query()
    s := val.Get("shortlink")
    fmt.Print("%s\n" ,s)
}

func (sl *ShortLink) redirect(w http.ResponseWriter, r *http.Request){
    val := mux.Vars(r)
    fmt.Print("%s\n", val["shortlink"])
}

func (sl *ShortLink) Run(addr string){
    log.Fatal(http.ListenAndServe(addr, sl.Router))
}

func responsewithError(w http.ResponseWriter, err error){
    switch e := err.(type){
    case Error:
        responseWithJson(w, e.Status(), e.Error())
        log.Printf("Http %d - %s", e.Status(), e)
    default:
        responseWithJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
    }
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}){
    resp,_ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(resp)
}