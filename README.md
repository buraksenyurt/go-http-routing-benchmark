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
 * [Denco](https://github.com/naoina/denco)

## Results

Benchmark System:
 * Intel Core i5 M 580 (4x 2.67GHz)
 * 2x 4 GiB DDR3-1066 RAM
 * go1.2.2 linux/amd64
 * Arch Linux amd64 (Linux Kernel 3.14.5)

```
#GithubAPI Routes: 203
#GPlusAPI Routes: 13
#ParseAPI Routes: 26
#Static Routes: 157

BenchmarkGocraftWeb_Param                1000000              2070 ns/op             673 B/op          9 allocs/op
BenchmarkGorillaMux_Param                 500000              6111 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_Param                5000000               724 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Param               2000000               799 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Param                    200000              7746 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_Param                       1000000              3403 ns/op            1061 B/op         17 allocs/op
BenchmarkTigerTonic_Param                1000000              3109 ns/op             831 B/op         16 allocs/op
BenchmarkTraffic_Param                    500000              7641 ns/op            2026 B/op         23 allocs/op
BenchmarkGoji_Param                      1000000              1052 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param                     5000000               454 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_Param                     5000000               377 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_Param20               200000             10102 ns/op            3864 B/op         17 allocs/op
BenchmarkGorillaMux_Param20               100000             21180 ns/op            3313 B/op         10 allocs/op
BenchmarkHttpRouter_Param20               500000              5973 ns/op            2219 B/op          4 allocs/op
BenchmarkHttpTreeMux_Param20              500000              6960 ns/op            2219 B/op          4 allocs/op
BenchmarkMartini_Param20                   50000             69963 ns/op            3711 B/op         16 allocs/op
BenchmarkPat_Param20                      100000             28017 ns/op            5733 B/op        154 allocs/op
BenchmarkTigerTonic_Param20                50000             45496 ns/op           11067 B/op        176 allocs/op
BenchmarkTraffic_Param20                   50000             34834 ns/op            8240 B/op         68 allocs/op
BenchmarkGoji_Param20                     500000              4442 ns/op            1260 B/op          2 allocs/op
BenchmarkKocha_Param20                    500000              5921 ns/op            1839 B/op         27 allocs/op
BenchmarkDenco_Param20                    500000              4235 ns/op            1679 B/op          7 allocs/op

BenchmarkGocraftWeb_ParamWrite           1000000              2205 ns/op             682 B/op         10 allocs/op
BenchmarkGorillaMux_ParamWrite            500000              6534 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite           2000000               787 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParamWrite          2000000               851 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParamWrite               200000              8934 ns/op            1286 B/op         16 allocs/op
BenchmarkPat_ParamWrite                   500000              3893 ns/op            1128 B/op         19 allocs/op
BenchmarkTigerTonic_ParamWrite            500000              5523 ns/op            1289 B/op         22 allocs/op
BenchmarkTraffic_ParamWrite               200000              9145 ns/op            2460 B/op         27 allocs/op
BenchmarkGoji_ParamWrite                 1000000              1104 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParamWrite                5000000               529 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_ParamWrite                5000000               457 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_GithubStatic         1000000              1226 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             47496 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000              83.6 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       20000000              82.7 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21498 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14183 ns/op            3788 B/op         76 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               408 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             81891 ns/op           23361 B/op        172 allocs/op
BenchmarkGoji_GithubStatic               5000000               473 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic             20000000               104 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_GithubStatic             50000000              68.4 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GithubParam          1000000              2538 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_GithubParam            50000             32959 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          1000000              1008 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GithubParam         1000000              1126 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GithubParam              100000             28862 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10366 ns/op            2625 B/op         56 allocs/op
BenchmarkTigerTonic_GithubParam           500000              5889 ns/op            1290 B/op         25 allocs/op
BenchmarkTraffic_GithubParam               50000             38382 ns/op            7148 B/op         60 allocs/op
BenchmarkGoji_GithubParam                1000000              1875 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GithubParam               2000000               979 ns/op             132 B/op          5 allocs/op
BenchmarkDenco_GithubParam               2000000               827 ns/op             116 B/op          3 allocs/op

BenchmarkGocraftWeb_GithubAll               5000            500594 ns/op          136407 B/op       1915 allocs/op
BenchmarkGorillaMux_GithubAll                100          20974751 ns/op          153332 B/op       1419 allocs/op
BenchmarkHttpRouter_GithubAll              10000            173594 ns/op           57351 B/op        347 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            190157 ns/op           57349 B/op        347 allocs/op
BenchmarkMartini_GithubAll                   100          12885120 ns/op          245356 B/op       2941 allocs/op
BenchmarkPat_GithubAll                       500           5813359 ns/op         1588664 B/op      32573 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1138578 ns/op          218336 B/op       5583 allocs/op
BenchmarkTraffic_GithubAll                   100          18290208 ns/op         3173802 B/op      24943 allocs/op
BenchmarkGoji_GithubAll                     2000            853669 ns/op           57307 B/op        347 allocs/op
BenchmarkKocha_GithubAll                   10000            185275 ns/op           24060 B/op        847 allocs/op
BenchmarkDenco_GithubAll                   10000            153058 ns/op           21353 B/op        508 allocs/op

BenchmarkGocraftWeb_GPlusStatic          1000000              1156 ns/op             297 B/op          6 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4253 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              47.4 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        50000000              45.5 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              5976 ns/op             861 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               446 ns/op              99 B/op          2 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               239 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6144 ns/op            1511 B/op         19 allocs/op
BenchmarkGoji_GPlusStatic                5000000               327 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic              50000000              70.9 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_GPlusStatic              50000000              44.8 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_GPlusParam           1000000              2140 ns/op             673 B/op          9 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             10404 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam           2000000               813 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2000000               882 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlusParam               200000             11012 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2646 ns/op             752 B/op         14 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              3623 ns/op             908 B/op         16 allocs/op
BenchmarkTraffic_GPlusParam               200000             10139 ns/op            2041 B/op         23 allocs/op
BenchmarkGoji_GPlusParam                 1000000              1163 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlusParam                5000000               525 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_GPlusParam                5000000               446 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_GPlus2Params         1000000              2628 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             28433 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         2000000               927 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        1000000              1066 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlus2Params              50000             35702 ns/op            1318 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000              8588 ns/op            2401 B/op         41 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              6255 ns/op            1391 B/op         25 allocs/op
BenchmarkTraffic_GPlus2Params              50000             30983 ns/op            3623 B/op         35 allocs/op
BenchmarkGoji_GPlus2Params               1000000              1756 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlus2Params              2000000               966 ns/op             132 B/op          5 allocs/op
BenchmarkDenco_GPlus2Params              2000000               818 ns/op             116 B/op          3 allocs/op

BenchmarkGocraftWeb_GPlusAll              100000             29642 ns/op            8343 B/op        117 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            183334 ns/op            9727 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll              200000             10103 ns/op            3775 B/op         22 allocs/op
BenchmarkHttpTreeMux_GPlusAll             200000             10950 ns/op            3775 B/op         22 allocs/op
BenchmarkMartini_GPlusAll                  10000            200166 ns/op           15531 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             66865 ns/op           17695 B/op        346 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             61863 ns/op           13335 B/op        289 allocs/op
BenchmarkTraffic_GPlusAll                  10000            218343 ns/op           42106 B/op        447 allocs/op
BenchmarkGoji_GPlusAll                    100000             16932 ns/op            3775 B/op         22 allocs/op
BenchmarkKocha_GPlusAll                   200000              8943 ns/op            1014 B/op         43 allocs/op
BenchmarkDenco_GPlusAll                   500000              7544 ns/op             886 B/op         27 allocs/op

BenchmarkGocraftWeb_ParseStatic          1000000              1231 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8479 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              41.3 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        20000000              82.7 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              6807 ns/op             860 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1054 ns/op             249 B/op          5 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               350 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9792 ns/op            2392 B/op         25 allocs/op
BenchmarkGoji_ParseStatic                5000000               424 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic              20000000              77.2 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_ParseStatic              50000000              49.0 ns/op               0 B/op          0 allocs/op

BenchmarkGocraftWeb_ParseParam           1000000              2195 ns/op             689 B/op          9 allocs/op
BenchmarkGorillaMux_ParseParam            200000              9896 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam           2000000               775 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParseParam          2000000               862 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParseParam               200000              9826 ns/op            1187 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3864 ns/op            1197 B/op         20 allocs/op
BenchmarkTigerTonic_ParseParam            500000              3447 ns/op             889 B/op         16 allocs/op
BenchmarkTraffic_ParseParam               200000              9865 ns/op            2326 B/op         25 allocs/op
BenchmarkGoji_ParseParam                 1000000              1335 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParseParam                5000000               491 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_ParseParam                5000000               410 ns/op              50 B/op          2 allocs/op

BenchmarkGocraftWeb_Parse2Params         1000000              2477 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10477 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params         2000000               871 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Parse2Params        2000000               996 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Parse2Params             200000             10146 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3809 ns/op             908 B/op         21 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              5852 ns/op            1294 B/op         25 allocs/op
BenchmarkTraffic_Parse2Params             200000              9986 ns/op            2131 B/op         25 allocs/op
BenchmarkGoji_Parse2Params               1000000              1313 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Parse2Params              2000000               864 ns/op             132 B/op          5 allocs/op
BenchmarkDenco_Parse2Params              5000000               694 ns/op             116 B/op          3 allocs/op

BenchmarkGocraftWeb_ParseAll               50000             51125 ns/op           14303 B/op        209 allocs/op
BenchmarkGorillaMux_ParseAll                5000            380992 ns/op           17263 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              200000             14825 ns/op            5491 B/op         33 allocs/op
BenchmarkHttpTreeMux_ParseAll             100000             16998 ns/op            5491 B/op         33 allocs/op
BenchmarkMartini_ParseAll                  10000            256962 ns/op           27684 B/op        333 allocs/op
BenchmarkPat_ParseAll                      50000             73114 ns/op           18283 B/op        385 allocs/op
BenchmarkTigerTonic_ParseAll               20000             81033 ns/op           17738 B/op        372 allocs/op
BenchmarkTraffic_ParseAll                   5000            314149 ns/op           70596 B/op        763 allocs/op
BenchmarkGoji_ParseAll                    100000             27308 ns/op            5491 B/op         33 allocs/op
BenchmarkKocha_ParseAll                   200000             12027 ns/op            1159 B/op         54 allocs/op
BenchmarkDenco_ParseAll                   200000             10017 ns/op            1008 B/op         35 allocs/op

BenchmarkHttpServeMux_StaticAll             2000           1211103 ns/op             104 B/op          8 allocs/op
BenchmarkGocraftWeb_StaticAll              10000            210604 ns/op           49187 B/op        951 allocs/op
BenchmarkGorillaMux_StaticAll                500           5834407 ns/op           72353 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll             100000             26101 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll            100000             26077 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll                   500           4230504 ns/op          145619 B/op       2521 allocs/op
BenchmarkPat_StaticAll                      1000           2095654 ns/op          554265 B/op      11250 allocs/op
BenchmarkTigerTonic_StaticAll              20000             76682 ns/op            7777 B/op        158 allocs/op
BenchmarkTraffic_StaticAll                   100          14647884 ns/op         3796423 B/op      27937 allocs/op
BenchmarkGoji_StaticAll                    10000            101863 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_StaticAll                  100000             26820 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_StaticAll                  100000             15473 ns/op               0 B/op          0 allocs/op
```
