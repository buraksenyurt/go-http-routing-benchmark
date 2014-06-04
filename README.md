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

BenchmarkHttpServeMux_StaticAll             2000           1191484 ns/op             104 B/op          8 allocs/op
BenchmarkGocraftWeb_StaticAll              10000            211628 ns/op           49151 B/op        951 allocs/op
BenchmarkGorillaMux_StaticAll                500           5740597 ns/op           72281 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll             100000             25341 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll                   500           4306435 ns/op          145552 B/op       2521 allocs/op
BenchmarkPat_StaticAll                      1000           2087807 ns/op          554182 B/op      11249 allocs/op
BenchmarkTigerTonic_StaticAll              20000             75880 ns/op            7775 B/op        158 allocs/op
BenchmarkTraffic_StaticAll                   100          14411410 ns/op         3795145 B/op      27926 allocs/op
BenchmarkKocha_StaticAll                  100000             26720 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_Param                1000000              2073 ns/op             672 B/op          9 allocs/op
BenchmarkGorillaMux_Param                 500000              6103 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_Param                5000000               721 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Param                    200000              7870 ns/op            1185 B/op         13 allocs/op
BenchmarkPat_Param                        500000              3413 ns/op            1061 B/op         17 allocs/op
BenchmarkTigerTonic_Param                1000000              3109 ns/op             830 B/op         16 allocs/op
BenchmarkTraffic_Param                    500000              7618 ns/op            2024 B/op         23 allocs/op
BenchmarkKocha_Param                     5000000               450 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_ParamWrite           1000000              2199 ns/op             681 B/op         10 allocs/op
BenchmarkGorillaMux_ParamWrite            500000              6601 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite           2000000               779 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParamWrite               200000              9219 ns/op            1285 B/op         16 allocs/op
BenchmarkPat_ParamWrite                   500000              3935 ns/op            1128 B/op         19 allocs/op
BenchmarkTigerTonic_ParamWrite            500000              5368 ns/op            1288 B/op         22 allocs/op
BenchmarkTraffic_ParamWrite               200000              9143 ns/op            2457 B/op         27 allocs/op
BenchmarkKocha_ParamWrite                5000000               527 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_GithubStatic         1000000              1233 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             47688 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000              87.6 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21492 ns/op             859 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14071 ns/op            3787 B/op         76 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               411 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             81192 ns/op           23357 B/op        172 allocs/op
BenchmarkKocha_GithubStatic             20000000               104 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GithubParam          1000000              2571 ns/op             735 B/op         10 allocs/op
BenchmarkGorillaMux_GithubParam            50000             32533 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          2000000              1001 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GithubParam              100000             29295 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10139 ns/op            2624 B/op         56 allocs/op
BenchmarkTigerTonic_GithubParam           500000              5848 ns/op            1289 B/op         25 allocs/op
BenchmarkTraffic_GithubParam               50000             37494 ns/op            7146 B/op         60 allocs/op
BenchmarkKocha_GithubParam               2000000               943 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_GithubAll               5000            497417 ns/op          136285 B/op       1914 allocs/op
BenchmarkGorillaMux_GithubAll                100          20991494 ns/op          153242 B/op       1419 allocs/op
BenchmarkHttpRouter_GithubAll              10000            171746 ns/op           57342 B/op        347 allocs/op
BenchmarkMartini_GithubAll                   100          12759051 ns/op          245260 B/op       2940 allocs/op
BenchmarkPat_GithubAll                       500           5779354 ns/op         1587845 B/op      32570 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1127252 ns/op          218104 B/op       5581 allocs/op
BenchmarkTraffic_GithubAll                   100          18132660 ns/op         3172910 B/op      24936 allocs/op
BenchmarkKocha_GithubAll                   10000            186238 ns/op           24036 B/op        847 allocs/op

BenchmarkGocraftWeb_GPlusStatic          1000000              1156 ns/op             297 B/op          6 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4322 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              50.0 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              6085 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               433 ns/op              99 B/op          2 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               234 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6188 ns/op            1509 B/op         19 allocs/op
BenchmarkKocha_GPlusStatic              50000000              70.7 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GPlusParam           1000000              2170 ns/op             672 B/op          9 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             10396 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam           2000000               818 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlusParam               200000             11105 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2668 ns/op             752 B/op         14 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              3606 ns/op             907 B/op         16 allocs/op
BenchmarkTraffic_GPlusParam               200000             10146 ns/op            2039 B/op         23 allocs/op
BenchmarkKocha_GPlusParam                5000000               528 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_GPlus2Params         1000000              2601 ns/op             735 B/op         10 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             28356 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         2000000               927 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlus2Params              50000             35777 ns/op            1317 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000              8592 ns/op            2399 B/op         41 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              6215 ns/op            1391 B/op         25 allocs/op
BenchmarkTraffic_GPlus2Params              50000             31391 ns/op            3620 B/op         35 allocs/op
BenchmarkKocha_GPlus2Params              2000000               956 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_GPlusAll              100000             29203 ns/op            8336 B/op        117 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            180427 ns/op            9723 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll              200000             10127 ns/op            3774 B/op         22 allocs/op
BenchmarkMartini_GPlusAll                  10000            202288 ns/op           15519 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             67234 ns/op           17682 B/op        346 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             61545 ns/op           13326 B/op        289 allocs/op
BenchmarkTraffic_GPlusAll                  10000            214557 ns/op           42066 B/op        446 allocs/op
BenchmarkKocha_GPlusAll                   200000              8857 ns/op            1013 B/op         43 allocs/op

BenchmarkGocraftWeb_ParseStatic          1000000              1223 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8546 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              41.0 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              6940 ns/op             860 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1045 ns/op             249 B/op          5 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               363 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9766 ns/op            2390 B/op         25 allocs/op
BenchmarkKocha_ParseStatic              20000000              76.4 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_ParseParam           1000000              2207 ns/op             688 B/op          9 allocs/op
BenchmarkGorillaMux_ParseParam            200000              9798 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam           2000000               767 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParseParam               200000             10043 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3910 ns/op            1196 B/op         20 allocs/op
BenchmarkTigerTonic_ParseParam            500000              3456 ns/op             888 B/op         16 allocs/op
BenchmarkTraffic_ParseParam               200000              9729 ns/op            2324 B/op         25 allocs/op
BenchmarkKocha_ParseParam                5000000               485 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_Parse2Params         1000000              2485 ns/op             735 B/op         10 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10454 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params         2000000               851 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Parse2Params             200000             10760 ns/op            1218 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3825 ns/op             907 B/op         21 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              5859 ns/op            1293 B/op         25 allocs/op
BenchmarkTraffic_Parse2Params             200000             10037 ns/op            2129 B/op         25 allocs/op
BenchmarkKocha_Parse2Params              2000000               859 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_ParseAll               50000             50881 ns/op           14293 B/op        209 allocs/op
BenchmarkGorillaMux_ParseAll                5000            379142 ns/op           17252 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              200000             14797 ns/op            5490 B/op         33 allocs/op
BenchmarkMartini_ParseAll                  10000            260559 ns/op           27668 B/op        333 allocs/op
BenchmarkPat_ParseAll                      50000             74220 ns/op           18271 B/op        385 allocs/op
BenchmarkTigerTonic_ParseAll               20000             81828 ns/op           17724 B/op        372 allocs/op
BenchmarkTraffic_ParseAll                   5000            315006 ns/op           70532 B/op        763 allocs/op
BenchmarkKocha_ParseAll                   200000             12037 ns/op            1159 B/op         54 allocs/op
```
