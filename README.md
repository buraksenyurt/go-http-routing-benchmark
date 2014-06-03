go-http-routing-benchmark
=========================

This benchmark suite aims to compare the performance of available HTTP request routers for Go by implementing the routing of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.


Included Routers:

 * [Gocraft Web](https://github.com/gocraft/web)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [net/http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [Martini](https://github.com/codegangsta/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)
 * [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter)

## Results

Benchmark System:
 * Intel Core i5 M 580 (4x 2.67GHz)
 * 2x 4 GiB DDR3-1066 RAM
 * go1.2.2 linux/amd64
 * Arch Linux amd64 (Linux Kernel 3.14.4)

```
#GithubAPI Routes: 203
#GPlusAPI Routes: 13
#ParseAPI Routes: 26
#Static Routes: 157

BenchmarkHttpServeMux_StaticAll             2000           1199162 ns/op             104 B/op          8 allocs/op
BenchmarkGocraftWeb_StaticAll              10000            213201 ns/op           49175 B/op        951 allocs/op
BenchmarkGorillaMux_StaticAll                500           5758968 ns/op           72304 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll             100000             25646 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll                   500           4235418 ns/op          145610 B/op       2521 allocs/op
BenchmarkPat_StaticAll                      1000           2140598 ns/op          554179 B/op      11249 allocs/op
BenchmarkTigerTonic_StaticAll              20000             77019 ns/op            7776 B/op        158 allocs/op
BenchmarkTraffic_StaticAll                   100          14632911 ns/op         3795334 B/op      27927 allocs/op
BenchmarkKocha_StaticAll                   50000             32559 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_Param                1000000              2099 ns/op             672 B/op          9 allocs/op
BenchmarkGorillaMux_Param                 500000              6028 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_Param                5000000               722 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Param                    200000              7895 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_Param                        500000              3457 ns/op            1061 B/op         17 allocs/op
BenchmarkTigerTonic_Param                 500000              3894 ns/op            1025 B/op         19 allocs/op
BenchmarkTraffic_Param                    500000              7673 ns/op            2026 B/op         23 allocs/op
BenchmarkKocha_Param                     5000000               459 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_ParamWrite           1000000              2244 ns/op             682 B/op         10 allocs/op
BenchmarkGorillaMux_ParamWrite            500000              6399 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite           2000000               789 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParamWrite               200000              9130 ns/op            1285 B/op         16 allocs/op
BenchmarkPat_ParamWrite                   500000              4028 ns/op            1128 B/op         19 allocs/op
BenchmarkTigerTonic_ParamWrite            500000              6177 ns/op            1483 B/op         25 allocs/op
BenchmarkTraffic_ParamWrite               200000              9259 ns/op            2459 B/op         27 allocs/op
BenchmarkKocha_ParamWrite                5000000               536 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_GithubStatic         1000000              1243 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             47884 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000              84.8 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21760 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14226 ns/op            3788 B/op         76 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               415 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             81759 ns/op           23359 B/op        172 allocs/op
BenchmarkKocha_GithubStatic             20000000               128 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GithubParam          1000000              2585 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_GithubParam            50000             33182 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          1000000              1017 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GithubParam              100000             28901 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10286 ns/op            2625 B/op         56 allocs/op
BenchmarkTigerTonic_GithubParam           500000              6638 ns/op            1483 B/op         28 allocs/op
BenchmarkTraffic_GithubParam               50000             37680 ns/op            7147 B/op         60 allocs/op
BenchmarkKocha_GithubParam               2000000               954 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_GithubAll               5000            501369 ns/op          136366 B/op       1914 allocs/op
BenchmarkGorillaMux_GithubAll                100          21039958 ns/op          153298 B/op       1419 allocs/op
BenchmarkHttpRouter_GithubAll              10000            173544 ns/op           57345 B/op        347 allocs/op
BenchmarkMartini_GithubAll                   100          12825707 ns/op          245373 B/op       2941 allocs/op
BenchmarkPat_GithubAll                       500           5834784 ns/op         1588280 B/op      32572 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1277511 ns/op          250731 B/op       6086 allocs/op
BenchmarkTraffic_GithubAll                   100          18116267 ns/op         3172976 B/op      24936 allocs/op
BenchmarkKocha_GithubAll                   10000            185067 ns/op           24051 B/op        847 allocs/op

BenchmarkGocraftWeb_GPlusStatic          1000000              1168 ns/op             297 B/op          6 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4134 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              46.6 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              6048 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               431 ns/op              99 B/op          2 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               237 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6197 ns/op            1510 B/op         19 allocs/op
BenchmarkKocha_GPlusStatic              20000000              91.1 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GPlusParam           1000000              2192 ns/op             673 B/op          9 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             10317 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam           2000000               818 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlusParam               200000             11102 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2683 ns/op             752 B/op         14 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              4421 ns/op            1102 B/op         19 allocs/op
BenchmarkTraffic_GPlusParam               200000             10169 ns/op            2040 B/op         23 allocs/op
BenchmarkKocha_GPlusParam                5000000               509 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_GPlus2Params         1000000              2618 ns/op             735 B/op         10 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             28645 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         2000000               930 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlus2Params              50000             36527 ns/op            1318 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000              8662 ns/op            2400 B/op         41 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              7117 ns/op            1584 B/op         28 allocs/op
BenchmarkTraffic_GPlus2Params              50000             31365 ns/op            3622 B/op         35 allocs/op
BenchmarkKocha_GPlus2Params              2000000               950 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_GPlusAll              100000             29630 ns/op            8341 B/op        117 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            180852 ns/op            9725 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll              200000             10130 ns/op            3774 B/op         22 allocs/op
BenchmarkMartini_GPlusAll                  10000            202275 ns/op           15526 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             67527 ns/op           17689 B/op        346 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             70610 ns/op           15473 B/op        322 allocs/op
BenchmarkTraffic_GPlusAll                  10000            216874 ns/op           42092 B/op        446 allocs/op
BenchmarkKocha_GPlusAll                   200000              8741 ns/op            1014 B/op         43 allocs/op

BenchmarkGocraftWeb_ParseStatic          1000000              1244 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8323 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              41.5 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              6821 ns/op             860 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1064 ns/op             249 B/op          5 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               352 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9793 ns/op            2391 B/op         25 allocs/op
BenchmarkKocha_ParseStatic              20000000               104 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_ParseParam           1000000              2246 ns/op             689 B/op          9 allocs/op
BenchmarkGorillaMux_ParseParam            200000              9823 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam           2000000               772 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParseParam               200000             10053 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3939 ns/op            1197 B/op         20 allocs/op
BenchmarkTigerTonic_ParseParam            500000              4245 ns/op            1083 B/op         19 allocs/op
BenchmarkTraffic_ParseParam               200000              9941 ns/op            2325 B/op         25 allocs/op
BenchmarkKocha_ParseParam                5000000               489 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_Parse2Params         1000000              2515 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10343 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params         2000000               866 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Parse2Params             200000             10115 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3902 ns/op             907 B/op         21 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              6702 ns/op            1487 B/op         28 allocs/op
BenchmarkTraffic_Parse2Params             200000             10218 ns/op            2131 B/op         25 allocs/op
BenchmarkKocha_Parse2Params              2000000               869 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_ParseAll               50000             51167 ns/op           14300 B/op        209 allocs/op
BenchmarkGorillaMux_ParseAll                5000            374276 ns/op           17258 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              200000             14866 ns/op            5490 B/op         33 allocs/op
BenchmarkMartini_ParseAll                  10000            259554 ns/op           27677 B/op        333 allocs/op
BenchmarkPat_ParseAll                      50000             74040 ns/op           18277 B/op        385 allocs/op
BenchmarkTigerTonic_ParseAll               20000             95040 ns/op           20848 B/op        420 allocs/op
BenchmarkTraffic_ParseAll                   5000            316818 ns/op           70572 B/op        763 allocs/op
BenchmarkKocha_ParseAll                   200000             12411 ns/op            1159 B/op         54 allocs/op
```
