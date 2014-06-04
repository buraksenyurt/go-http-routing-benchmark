go-http-routing-benchmark
=========================

This benchmark suite aims to compare the performance of available HTTP request routers for Go by implementing the routing of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.


Included Routers:

 * [Gocraft Web](https://github.com/gocraft/web)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [net/http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [HttpTreeMux](https://github.com/dimfeld/httptreemux)
 * [Martini](https://github.com/codegangsta/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)
 * [Goji](https://github.com/zenazn/goji)
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

BenchmarkGocraftWeb_Param                1000000              2100 ns/op             673 B/op          9 allocs/op
BenchmarkGorillaMux_Param                 500000              6179 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_Param                5000000               721 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Param               2000000               787 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Param                    200000              7766 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_Param                       1000000              3407 ns/op            1061 B/op         17 allocs/op
BenchmarkTigerTonic_Param                1000000              3147 ns/op             830 B/op         16 allocs/op
BenchmarkTraffic_Param                    500000              7482 ns/op            2026 B/op         23 allocs/op
BenchmarkGoji_Param                      1000000              1057 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param                     5000000               457 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_Param20               200000             10058 ns/op            3863 B/op         17 allocs/op
BenchmarkGorillaMux_Param20               100000             21346 ns/op            3311 B/op         10 allocs/op
BenchmarkHttpRouter_Param20               500000              5972 ns/op            2219 B/op          4 allocs/op
BenchmarkMartini_Param20                   50000             68462 ns/op            3713 B/op         16 allocs/op
BenchmarkPat_Param20                      100000             27855 ns/op            5731 B/op        154 allocs/op
BenchmarkTigerTonic_Param20                50000             45176 ns/op           11067 B/op        176 allocs/op
BenchmarkTraffic_Param20                   50000             34626 ns/op            8239 B/op         68 allocs/op
BenchmarkGoji_Param20                     500000              4475 ns/op            1260 B/op          2 allocs/op
BenchmarkKocha_Param20                    500000              5957 ns/op            1839 B/op         27 allocs/op

BenchmarkGocraftWeb_ParamWrite           1000000              2252 ns/op             682 B/op         10 allocs/op
BenchmarkGorillaMux_ParamWrite            500000              6473 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite           2000000               794 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParamWrite          2000000               870 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParamWrite               200000              9031 ns/op            1286 B/op         16 allocs/op
BenchmarkPat_ParamWrite                   500000              3969 ns/op            1128 B/op         19 allocs/op
BenchmarkTigerTonic_ParamWrite            500000              5397 ns/op            1288 B/op         22 allocs/op
BenchmarkTraffic_ParamWrite               200000              9702 ns/op            2459 B/op         27 allocs/op
BenchmarkGoji_ParamWrite                 1000000              1126 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParamWrite                5000000               534 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_GithubStatic         1000000              1251 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             47508 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000              84.9 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       20000000              82.9 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21481 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14232 ns/op            3788 B/op         76 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               401 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             82340 ns/op           23360 B/op        172 allocs/op
BenchmarkGoji_GithubStatic               5000000               477 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic             20000000               106 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GithubParam          1000000              2580 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_GithubParam            50000             33129 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          1000000              1019 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GithubParam         1000000              1141 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GithubParam              100000             28458 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10255 ns/op            2625 B/op         56 allocs/op
BenchmarkTigerTonic_GithubParam           500000              5924 ns/op            1289 B/op         25 allocs/op
BenchmarkTraffic_GithubParam               50000             38092 ns/op            7148 B/op         60 allocs/op
BenchmarkGoji_GithubParam                1000000              1888 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GithubParam               2000000               965 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_GithubAll               5000            503894 ns/op          136391 B/op       1914 allocs/op
BenchmarkGorillaMux_GithubAll                100          21007993 ns/op          153321 B/op       1419 allocs/op
BenchmarkHttpRouter_GithubAll              10000            175054 ns/op           57349 B/op        347 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            193433 ns/op           57347 B/op        347 allocs/op
BenchmarkMartini_GithubAll                   100          12895280 ns/op          245389 B/op       2941 allocs/op
BenchmarkPat_GithubAll                       500           5801080 ns/op         1588320 B/op      32572 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1141948 ns/op          218258 B/op       5582 allocs/op
BenchmarkTraffic_GithubAll                   100          18166704 ns/op         3173637 B/op      24941 allocs/op
BenchmarkGoji_GithubAll                     2000            859482 ns/op           57308 B/op        347 allocs/op
BenchmarkKocha_GithubAll                   10000            186609 ns/op           24054 B/op        847 allocs/op

BenchmarkGocraftWeb_GPlusStatic          1000000              1177 ns/op             297 B/op          6 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4356 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              46.6 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        50000000              45.7 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              6089 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               428 ns/op              99 B/op          2 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               243 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6232 ns/op            1510 B/op         19 allocs/op
BenchmarkGoji_GPlusStatic                5000000               331 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic              50000000              71.4 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GPlusParam           1000000              2170 ns/op             673 B/op          9 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             10441 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam           2000000               821 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2000000               892 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlusParam               200000             10946 ns/op            1187 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2679 ns/op             752 B/op         14 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              3670 ns/op             907 B/op         16 allocs/op
BenchmarkTraffic_GPlusParam               200000              9979 ns/op            2040 B/op         23 allocs/op
BenchmarkGoji_GPlusParam                 1000000              1211 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlusParam                5000000               537 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_GPlus2Params         1000000              2645 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             28403 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         2000000               932 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        1000000              1068 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlus2Params              50000             36387 ns/op            1318 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000              8600 ns/op            2400 B/op         41 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              6289 ns/op            1391 B/op         25 allocs/op
BenchmarkTraffic_GPlus2Params              50000             31226 ns/op            3622 B/op         35 allocs/op
BenchmarkGoji_GPlus2Params               1000000              1801 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlus2Params              2000000               982 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_GPlusAll              100000             29596 ns/op            8343 B/op        117 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            181566 ns/op            9726 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll              200000             10275 ns/op            3774 B/op         22 allocs/op
BenchmarkHttpTreeMux_GPlusAll             200000             12752 ns/op            3774 B/op         22 allocs/op
BenchmarkMartini_GPlusAll                  10000            203829 ns/op           15530 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             66973 ns/op           17692 B/op        346 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             62226 ns/op           13333 B/op        289 allocs/op
BenchmarkTraffic_GPlusAll                  10000            218756 ns/op           42101 B/op        447 allocs/op
BenchmarkGoji_GPlusAll                    100000             16992 ns/op            3774 B/op         22 allocs/op
BenchmarkKocha_GPlusAll                   200000              8990 ns/op            1014 B/op         43 allocs/op

BenchmarkGocraftWeb_ParseStatic          1000000              1243 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8668 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              41.1 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        20000000              81.6 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              6841 ns/op             860 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1063 ns/op             249 B/op          5 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               352 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9952 ns/op            2392 B/op         25 allocs/op
BenchmarkGoji_ParseStatic                5000000               426 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic              20000000              76.5 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_ParseParam           1000000              2230 ns/op             689 B/op          9 allocs/op
BenchmarkGorillaMux_ParseParam            200000             10016 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam           2000000               773 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParseParam          2000000               870 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParseParam               200000              9924 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3897 ns/op            1197 B/op         20 allocs/op
BenchmarkTigerTonic_ParseParam            500000              3511 ns/op             888 B/op         16 allocs/op
BenchmarkTraffic_ParseParam               200000              9774 ns/op            2326 B/op         25 allocs/op
BenchmarkGoji_ParseParam                 1000000              1359 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParseParam                5000000               494 ns/op              58 B/op          3 allocs/op

BenchmarkGocraftWeb_Parse2Params         1000000              2499 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10740 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params         2000000               856 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Parse2Params        2000000               994 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Parse2Params             200000             10185 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3881 ns/op             908 B/op         21 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              5915 ns/op            1294 B/op         25 allocs/op
BenchmarkTraffic_Parse2Params             200000             10022 ns/op            2131 B/op         25 allocs/op
BenchmarkGoji_Parse2Params               1000000              1366 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Parse2Params              2000000               861 ns/op             132 B/op          5 allocs/op

BenchmarkGocraftWeb_ParseAll               50000             51994 ns/op           14302 B/op        209 allocs/op
BenchmarkGorillaMux_ParseAll                5000            384687 ns/op           17260 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              200000             15004 ns/op            5491 B/op         33 allocs/op
BenchmarkHttpTreeMux_ParseAll             100000             17476 ns/op            5490 B/op         33 allocs/op
BenchmarkMartini_ParseAll                  10000            259029 ns/op           27686 B/op        333 allocs/op
BenchmarkPat_ParseAll                      50000             74085 ns/op           18280 B/op        385 allocs/op
BenchmarkTigerTonic_ParseAll               20000             82102 ns/op           17735 B/op        372 allocs/op
BenchmarkTraffic_ParseAll                   5000            317812 ns/op           70586 B/op        763 allocs/op
BenchmarkGoji_ParseAll                    100000             27721 ns/op            5491 B/op         33 allocs/op
BenchmarkKocha_ParseAll                   200000             12085 ns/op            1159 B/op         54 allocs/op

BenchmarkHttpServeMux_StaticAll             2000           1200105 ns/op             104 B/op          8 allocs/op
BenchmarkGocraftWeb_StaticAll              10000            214494 ns/op           49182 B/op        951 allocs/op
BenchmarkGorillaMux_StaticAll                500           5835945 ns/op           72345 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll             100000             26135 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll            100000             26158 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll                   500           4253387 ns/op          145612 B/op       2521 allocs/op
BenchmarkPat_StaticAll                      1000           2106109 ns/op          554252 B/op      11250 allocs/op
BenchmarkTigerTonic_StaticAll              20000             76379 ns/op            7777 B/op        158 allocs/op
BenchmarkTraffic_StaticAll                   100          14679153 ns/op         3796022 B/op      27933 allocs/op
BenchmarkGoji_StaticAll                    10000            101306 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_StaticAll                  100000             26923 ns/op               0 B/op          0 allocs/op
```
