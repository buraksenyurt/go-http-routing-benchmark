// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package benchmark

import (
	"net/http"
	"testing"
)

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var gplusAPI = []route{
	// People
	{"GET", "/people/:userId"},
	{"GET", "/people"},
	{"GET", "/activities/:activityId/people/:collection"},
	{"GET", "/people/:userId/people/:collection"},
	{"GET", "/people/:userId/openIdConnect"},

	// Activities
	{"GET", "/people/:userId/activities/:collection"},
	{"GET", "/activities/:activityId"},
	{"GET", "/activities"},

	// Comments
	{"GET", "/activities/:activityId/comments"},
	{"GET", "/comments/:commentId"},

	// Moments
	{"POST", "/people/:userId/moments/:collection"},
	{"GET", "/people/:userId/moments/:collection"},
	{"DELETE", "/moments/:id"},
}

var (
	gplusGocraftWeb  http.Handler
	gplusGorillaMux  http.Handler
	gplusHttpRouter  http.Handler
	gplusHttpTreeMux http.Handler
	gplusMartini     http.Handler
	gplusPat         http.Handler
	gplusTigerTonic  http.Handler
	gplusTraffic     http.Handler
	gplusGoji        http.Handler
	gplusKocha       http.Handler
	gplusDenco       http.Handler
	gplusRevel       http.Handler
)

func init() {
	println("#GPlusAPI Routes:", len(gplusAPI))

	gplusGocraftWeb = loadGocraftWeb(gplusAPI)
	gplusGorillaMux = loadGorillaMux(gplusAPI)
	gplusHttpRouter = loadHttpRouter(gplusAPI)
	gplusHttpTreeMux = loadHttpTreeMux(gplusAPI)
	gplusMartini = loadMartini(gplusAPI)
	gplusPat = loadPat(gplusAPI)
	gplusTigerTonic = loadTigerTonic(gplusAPI)
	gplusTraffic = loadTraffic(gplusAPI)
	gplusGoji = loadGoji(gplusAPI)
	gplusKocha = loadKocha(gplusAPI)
	gplusDenco = loadDenco(gplusAPI)
	gplusRevel = loadRevel(gplusAPI)
}

// Static
func BenchmarkGocraftWeb_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGocraftWeb, req)
}
func BenchmarkGorillaMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	req.RequestURI = "/people"
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkMartini_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkTigerTonic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkGoji_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusGoji, req)
}
func BenchmarkKocha_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusKocha, req)
}
func BenchmarkDenco_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkRevel_GPlusStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people", nil)
	benchRequest(b, gplusRevel, req)
}

// One Param
func BenchmarkGocraftWeb_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGocraftWeb, req)
}
func BenchmarkGorillaMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	req.RequestURI = "/people/118051310819094153327"
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkMartini_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkTigerTonic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkGoji_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusGoji, req)
}
func BenchmarkKocha_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusKocha, req)
}
func BenchmarkDenco_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkRevel_GPlusParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327", nil)
	benchRequest(b, gplusRevel, req)
}

// Two Params
func BenchmarkGocraftWeb_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGocraftWeb, req)
}
func BenchmarkGorillaMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGorillaMux, req)
}
func BenchmarkHttpRouter_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusHttpRouter, req)
}
func BenchmarkHttpTreeMux_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	req.RequestURI = "/people/118051310819094153327/activities/123456789"
	benchRequest(b, gplusHttpTreeMux, req)
}
func BenchmarkMartini_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusMartini, req)
}
func BenchmarkPat_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusPat, req)
}
func BenchmarkTigerTonic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTigerTonic, req)
}
func BenchmarkTraffic_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusTraffic, req)
}
func BenchmarkGoji_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusGoji, req)
}
func BenchmarkKocha_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusKocha, req)
}
func BenchmarkDenco_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusDenco, req)
}
func BenchmarkRevel_GPlus2Params(b *testing.B) {
	req, _ := http.NewRequest("GET", "/people/118051310819094153327/activities/123456789", nil)
	benchRequest(b, gplusRevel, req)
}

// All Routes
func BenchmarkGocraftWeb_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGocraftWeb, gplusAPI)
}
func BenchmarkGorillaMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGorillaMux, gplusAPI)
}
func BenchmarkHttpRouter_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpRouter, gplusAPI)
}
func BenchmarkHttpTreeMux_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusHttpTreeMux, gplusAPI)
}
func BenchmarkMartini_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusMartini, gplusAPI)
}
func BenchmarkPat_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusPat, gplusAPI)
}
func BenchmarkTigerTonic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTigerTonic, gplusAPI)
}
func BenchmarkTraffic_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusTraffic, gplusAPI)
}
func BenchmarkGoji_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusGoji, gplusAPI)
}
func BenchmarkKocha_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusKocha, gplusAPI)
}
func BenchmarkDenco_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusDenco, gplusAPI)
}
func BenchmarkRevel_GPlusAll(b *testing.B) {
	benchRoutes(b, gplusRevel, gplusAPI)
}
