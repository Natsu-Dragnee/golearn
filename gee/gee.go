package gee

import (
	fmt "fmt"
	"net/http"
)

type handel func(http.ResponseWriter, *http.Request)


type Engine struct{
	router map[string]handel
}

func New() *Engine{
	return &Engine{router:make(map[string]handel)}
}

func (engine *Engine) addRoute(method string,pattern string,f handel){
	key := method + "-" + pattern
	engine.router[key] = f
}

func(engine * Engine) GET(pattern string,f handel){
	engine.addRoute("GET",pattern,f)
}

func(engine* Engine) POST(pattern string,f handel){
	engine.addRoute("POST",pattern,f)
}

func(engine* Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr,engine)
}

func(engine* Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handel,ok := engine.router[key]; ok {
		handel(w, req)
	}else{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}