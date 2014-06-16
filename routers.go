// Copyright 2014 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"revel"
	"runtime"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/bmizerany/pat"
	"github.com/dimfeld/httptreemux"
	"github.com/go-martini/martini"
	"github.com/gocraft/web"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"github.com/naoina/denco"
	"github.com/naoina/kocha-urlrouter"
	_ "github.com/naoina/kocha-urlrouter/doublearray"
	"github.com/pilu/traffic"
	"github.com/rcrowley/go-tigertonic"
	goji "github.com/zenazn/goji/web"
)

type route struct {
	method string
	path   string
}

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

var nullLogger *log.Logger

func init() {
	// beego sets it to runtime.NumCPU()
	// Currently none of the contestors does concurrent routing
	runtime.GOMAXPROCS(1)

	// makes logging 'webscale' (ignores them)
	log.SetOutput(new(mockResponseWriter))
	nullLogger = log.New(new(mockResponseWriter), "", 0)

	beego.RunMode = "prod"
	martini.Env = martini.Prod
	traffic.SetVar("env", "bench")
}

// Common
func httpHandlerFunc(w http.ResponseWriter, r *http.Request) {}

// beego
func beegoHandler(ctx *context.Context) {}

func beegoHandlerWrite(ctx *context.Context) {
	ctx.WriteString(ctx.Input.Param(":name"))
}

func loadBeego(routes []route) http.Handler {
	re := regexp.MustCompile(":([^/]*)")
	app := beego.NewApp()
	for _, route := range routes {
		route.path = re.ReplaceAllString(route.path, ":$1!")

		switch route.method {
		case "GET":
			app.Get(route.path, beegoHandler)
		case "POST":
			app.Post(route.path, beegoHandler)
		case "PUT":
			app.Put(route.path, beegoHandler)
		case "PATCH":
			app.Patch(route.path, beegoHandler)
		case "DELETE":
			app.Delete(route.path, beegoHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return app.Handlers
}

func loadBeegoSingle(method, path string, handler beego.FilterFunc) http.Handler {
	app := beego.NewApp()
	switch method {
	case "GET":
		app.Get(path, handler)
	case "POST":
		app.Post(path, handler)
	case "PUT":
		app.Put(path, handler)
	case "PATCH":
		app.Patch(path, handler)
	case "DELETE":
		app.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return app.Handlers
}

// Denco
func dencoHandle(w http.ResponseWriter, r *http.Request, _ denco.Params) {}

func dencoHandlerWrite(w http.ResponseWriter, r *http.Request, params denco.Params) {
	io.WriteString(w, params.Get("name"))
}

func loadDenco(routes []route) http.Handler {
	mux := denco.NewMux()
	handlers := make([]denco.Handler, len(routes))
	for i, route := range routes {
		handlers[i] = mux.Handler(route.method, route.path, dencoHandle)
	}
	handler, err := mux.Build(handlers)
	if err != nil {
		panic(err)
	}
	return handler
}

func loadDencoSingle(method, path string, handle denco.HandlerFunc) http.Handler {
	mux := denco.NewMux()
	handler, err := mux.Build([]denco.Handler{{method, path, handle}})
	if err != nil {
		panic(err)
	}
	return handler
}

// gocraft/web
type gocraftWebContext struct{}

func gocraftWebHandler(w web.ResponseWriter, r *web.Request) {}

func gocraftWebHandlerWrite(w web.ResponseWriter, r *web.Request) {
	io.WriteString(w, r.PathParams["name"])
}

func loadGocraftWeb(routes []route) http.Handler {
	router := web.New(gocraftWebContext{})
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, gocraftWebHandler)
		case "POST":
			router.Post(route.path, gocraftWebHandler)
		case "PUT":
			router.Put(route.path, gocraftWebHandler)
		case "PATCH":
			router.Patch(route.path, gocraftWebHandler)
		case "DELETE":
			router.Delete(route.path, gocraftWebHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

func loadGocraftWebSingle(method, path string, handler interface{}) http.Handler {
	router := web.New(gocraftWebContext{})
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return router
}

// goji
func gojiFuncWrite(c goji.C, w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, c.URLParams["name"])
}

func loadGoji(routes []route) http.Handler {
	mux := goji.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			mux.Get(route.path, httpHandlerFunc)
		case "POST":
			mux.Post(route.path, httpHandlerFunc)
		case "PUT":
			mux.Put(route.path, httpHandlerFunc)
		case "PATCH":
			mux.Patch(route.path, httpHandlerFunc)
		case "DELETE":
			mux.Delete(route.path, httpHandlerFunc)
		default:
			panic("Unknown HTTP method: " + route.method)
		}
	}
	return mux
}

func loadGojiSingle(method, path string, handler interface{}) http.Handler {
	mux := goji.New()
	switch method {
	case "GET":
		mux.Get(path, handler)
	case "POST":
		mux.Post(path, handler)
	case "PUT":
		mux.Put(path, handler)
	case "PATCH":
		mux.Patch(path, handler)
	case "DELETE":
		mux.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return mux
}

// go-json-rest/rest
func goJsonRestHandler(w rest.ResponseWriter, req *rest.Request) {}

func goJsonRestHandlerWrite(w rest.ResponseWriter, req *rest.Request) {
	io.WriteString(w.(io.Writer), req.PathParam("name"))
}

func newGoJsonRestResourceHandler() *rest.ResourceHandler {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
		Logger:            nullLogger,
		ErrorLogger:       nullLogger,
		DisableXPoweredBy: true,
	}
	return &handler
}

func loadGoJsonRest(routes []route) http.Handler {
	handler := newGoJsonRestResourceHandler()
	restRoutes := make([]*rest.Route, 0, len(routes))
	for _, route := range routes {
		restRoutes = append(restRoutes,
			&rest.Route{route.method, route.path, goJsonRestHandler},
		)
	}
	handler.SetRoutes(restRoutes...)
	return handler
}

func loadGoJsonRestSingle(method, path string, hfunc rest.HandlerFunc) http.Handler {
	handler := newGoJsonRestResourceHandler()
	handler.SetRoutes(
		&rest.Route{method, path, hfunc},
	)
	return handler
}

// gorilla/mux
func gorillaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	io.WriteString(w, params["name"])
}

func loadGorillaMux(routes []route) http.Handler {
	re := regexp.MustCompile(":([^/]*)")
	m := mux.NewRouter()
	for _, route := range routes {
		m.HandleFunc(
			re.ReplaceAllString(route.path, "{$1}"),
			httpHandlerFunc,
		).Methods(route.method)
	}
	return m
}

func loadGorillaMuxSingle(method, path string, handler http.HandlerFunc) http.Handler {
	m := mux.NewRouter()
	m.HandleFunc(path, handler).Methods(method)
	return m
}

// HttpRouter
func httpRouterHandle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {}

func httpRouterHandleWrite(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	io.WriteString(w, ps.ByName("name"))
}

func loadHttpRouter(routes []route) http.Handler {
	router := httprouter.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, httpRouterHandle)
	}
	return router
}

func loadHttpRouterSingle(method, path string, handle httprouter.Handle) http.Handler {
	router := httprouter.New()
	router.Handle(method, path, handle)
	return router
}

// httpTreeMux
func httpTreeMuxHandler(w http.ResponseWriter, r *http.Request, vars map[string]string) {}

func httpTreeMuxHandlerWrite(w http.ResponseWriter, r *http.Request, vars map[string]string) {
	io.WriteString(w, vars["name"])
}

func loadHttpTreeMux(routes []route) http.Handler {
	router := httptreemux.New()
	for _, route := range routes {
		router.Handle(route.method, route.path, httpTreeMuxHandler)
	}
	return router
}

func loadHttpTreeMuxSingle(method, path string, handler httptreemux.HandlerFunc) http.Handler {
	router := httptreemux.New()
	router.Handle(method, path, handler)
	return router
}

// Kocha-urlrouter
type kochaHandler struct {
	routerMap map[string]urlrouter.URLRouter
	params    []urlrouter.Param
}

func (h *kochaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	meth, params := h.routerMap[r.Method].Lookup(r.URL.Path)
	h.params = params
	meth.(http.HandlerFunc).ServeHTTP(w, r)
}

func (h *kochaHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (h *kochaHandler) Post(w http.ResponseWriter, r *http.Request)   {}
func (h *kochaHandler) Put(w http.ResponseWriter, r *http.Request)    {}
func (h *kochaHandler) Patch(w http.ResponseWriter, r *http.Request)  {}
func (h *kochaHandler) Delete(w http.ResponseWriter, r *http.Request) {}
func (h *kochaHandler) kochaHandlerWrite(w http.ResponseWriter, r *http.Request) {
	var name string
	for _, param := range h.params {
		if param.Name == "name" {
			name = param.Value
			break
		}
	}
	io.WriteString(w, name)
}

func loadKocha(routes []route) http.Handler {
	handler := &kochaHandler{routerMap: map[string]urlrouter.URLRouter{
		"GET":    urlrouter.NewURLRouter("doublearray"),
		"POST":   urlrouter.NewURLRouter("doublearray"),
		"PUT":    urlrouter.NewURLRouter("doublearray"),
		"PATCH":  urlrouter.NewURLRouter("doublearray"),
		"DELETE": urlrouter.NewURLRouter("doublearray"),
	}}
	recordMap := make(map[string][]urlrouter.Record)
	for _, route := range routes {
		var f http.HandlerFunc
		switch route.method {
		case "GET":
			f = handler.Get
		case "POST":
			f = handler.Post
		case "PUT":
			f = handler.Put
		case "PATCH":
			f = handler.Patch
		case "DELETE":
			f = handler.Delete
		}
		recordMap[route.method] = append(
			recordMap[route.method],
			urlrouter.NewRecord(route.path, f),
		)
	}
	for method, records := range recordMap {
		if err := handler.routerMap[method].Build(records); err != nil {
			panic(err)
		}
	}
	return handler
}

func loadKochaSingle(method, path string, handler *kochaHandler, hfunc http.HandlerFunc) http.Handler {
	handler.routerMap = map[string]urlrouter.URLRouter{
		method: urlrouter.NewURLRouter("doublearray"),
	}

	if err := handler.routerMap[method].Build([]urlrouter.Record{
		urlrouter.NewRecord(path, hfunc),
	}); err != nil {
		panic(err)
	}
	return handler
}

// Martini
func martiniHandler() {}

func martiniHandlerWrite(params martini.Params) string {
	return params["name"]
}

func loadMartini(routes []route) http.Handler {
	router := martini.NewRouter()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, martiniHandler)
		case "POST":
			router.Post(route.path, martiniHandler)
		case "PUT":
			router.Put(route.path, martiniHandler)
		case "PATCH":
			router.Patch(route.path, martiniHandler)
		case "DELETE":
			router.Delete(route.path, martiniHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	martini := martini.New()
	martini.Action(router.Handle)
	return martini
}

func loadMartiniSingle(method, path string, handler interface{}) http.Handler {
	router := martini.NewRouter()
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}

	martini := martini.New()
	martini.Action(router.Handle)
	return martini
}

// pat
func patHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get(":name"))
}

func loadPat(routes []route) http.Handler {
	m := pat.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			m.Get(route.path, http.HandlerFunc(httpHandlerFunc))
		case "POST":
			m.Post(route.path, http.HandlerFunc(httpHandlerFunc))
		case "PUT":
			m.Put(route.path, http.HandlerFunc(httpHandlerFunc))
		case "DELETE":
			m.Del(route.path, http.HandlerFunc(httpHandlerFunc))
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return m
}

func loadPatSingle(method, path string, handler http.Handler) http.Handler {
	m := pat.New()
	switch method {
	case "GET":
		m.Get(path, handler)
	case "POST":
		m.Post(path, handler)
	case "PUT":
		m.Put(path, handler)
	case "DELETE":
		m.Del(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return m
}

// Revel
type revelHandler struct {
	router  *revel.Router
	params  map[string][]string
	methods map[string]http.HandlerFunc
}

func newRevelHandler(router *revel.Router) *revelHandler {
	return &revelHandler{
		router:  router,
		methods: make(map[string]http.HandlerFunc),
	}
}

func (h *revelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	match := h.router.Route(r)
	if match == nil {
		panic("Route not found: " + r.URL.Path)
	}
	h.params = match.Params
	h.methods[match.ControllerName+"."+match.MethodName].ServeHTTP(w, r)
}

func (h *revelHandler) Get(w http.ResponseWriter, r *http.Request)    {}
func (h *revelHandler) Post(w http.ResponseWriter, r *http.Request)   {}
func (h *revelHandler) Put(w http.ResponseWriter, r *http.Request)    {}
func (h *revelHandler) Patch(w http.ResponseWriter, r *http.Request)  {}
func (h *revelHandler) Delete(w http.ResponseWriter, r *http.Request) {}
func (h *revelHandler) revelHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, h.params["name"][0])
}

func loadRevel(routes []route) *revelHandler {
	router := revel.NewRouter("")
	handler := newRevelHandler(router)
	for _, r := range routes {
		var f http.HandlerFunc
		switch r.method {
		case "GET":
			f = handler.Get
		case "POST":
			f = handler.Post
		case "PUT":
			f = handler.Put
		case "PATCH":
			f = handler.Patch
		case "DELETE":
			f = handler.Delete
		}
		action := "revel." + r.method
		handler.methods[action] = f
		route := revel.NewRoute(r.method, r.path, action, "", "", 0)
		if err := router.Tree.Add(route.TreePath, route); err != nil {
			panic(err)
		}
	}
	return handler
}

func loadRevelSingle(method, path string, handler *revelHandler, hfunc http.HandlerFunc) http.Handler {
	handler.router = revel.NewRouter("")
	action := "revel." + method
	handler.methods = map[string]http.HandlerFunc{
		action: hfunc,
	}
	route := revel.NewRoute(method, path, action, "", "", 0)
	if err := handler.router.Tree.Add(route.TreePath, route); err != nil {
		panic(err)
	}
	return handler
}

// Tiger Tonic
func tigerTonicHandlerWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}

func loadTigerTonic(routes []route) http.Handler {
	re := regexp.MustCompile(":([^/]*)")
	mux := tigertonic.NewTrieServeMux()
	for _, route := range routes {
		mux.HandleFunc(route.method, re.ReplaceAllString(route.path, "{$1}"), httpHandlerFunc)
	}
	return mux
}

func loadTigerTonicSingle(method, path string, handler http.HandlerFunc) http.Handler {
	mux := tigertonic.NewTrieServeMux()
	mux.HandleFunc(method, path, handler)
	return mux
}

// Traffic
func trafficHandlerWrite(w traffic.ResponseWriter, r *traffic.Request) {
	io.WriteString(w, r.URL.Query().Get("name"))
}
func trafficHandler(w traffic.ResponseWriter, r *traffic.Request) {}

func loadTraffic(routes []route) http.Handler {
	router := traffic.New()
	for _, route := range routes {
		switch route.method {
		case "GET":
			router.Get(route.path, trafficHandler)
		case "POST":
			router.Post(route.path, trafficHandler)
		case "PUT":
			router.Put(route.path, trafficHandler)
		case "PATCH":
			router.Patch(route.path, trafficHandler)
		case "DELETE":
			router.Delete(route.path, trafficHandler)
		default:
			panic("Unknow HTTP method: " + route.method)
		}
	}
	return router
}

func loadTrafficSingle(method, path string, handler traffic.HttpHandleFunc) http.Handler {
	router := traffic.New()
	switch method {
	case "GET":
		router.Get(path, handler)
	case "POST":
		router.Post(path, handler)
	case "PUT":
		router.Put(path, handler)
	case "PATCH":
		router.Patch(path, handler)
	case "DELETE":
		router.Delete(path, handler)
	default:
		panic("Unknow HTTP method: " + method)
	}
	return router
}

// Usage notice
func main() {
	fmt.Println("Usage: go test -bench=.")
	os.Exit(1)
}
