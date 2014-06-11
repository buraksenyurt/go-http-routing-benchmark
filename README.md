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
 * [Revel](https://github.com/revel/revel)

## Results

Benchmark System:
 * Intel Core i5 M 580 (4x 2.67GHz)
 * 2x 4 GiB DDR3-1066 RAM
 * go1.2.2 linux/amd64
 * Arch Linux amd64 (Linux Kernel 3.14.6)

```
#GithubAPI Routes: 203
#GPlusAPI Routes: 13
#ParseAPI Routes: 26
#Static Routes: 157

BenchmarkGocraftWeb_Param                1000000              2059 ns/op             673 B/op          9 allocs/op
BenchmarkGorillaMux_Param                 500000              6182 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_Param                5000000               725 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Param               2000000               795 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Param                    500000              7684 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_Param                       1000000              3421 ns/op            1061 B/op         17 allocs/op
BenchmarkTigerTonic_Param                1000000              3104 ns/op             830 B/op         16 allocs/op
BenchmarkTraffic_Param                    200000              7664 ns/op            2026 B/op         23 allocs/op
BenchmarkGoji_Param                      1000000              1072 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param                     5000000               453 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_Param                     5000000               379 ns/op              50 B/op          2 allocs/op
BenchmarkRevel_Param                     1000000              1830 ns/op             638 B/op          8 allocs/op

BenchmarkGocraftWeb_Param20               200000             10068 ns/op            3863 B/op         17 allocs/op
BenchmarkGorillaMux_Param20               100000             21411 ns/op            3311 B/op         10 allocs/op
BenchmarkHttpRouter_Param20               500000              5888 ns/op            2219 B/op          4 allocs/op
BenchmarkHttpTreeMux_Param20              500000              6917 ns/op            2219 B/op          4 allocs/op
BenchmarkMartini_Param20                   50000             68439 ns/op            3711 B/op         16 allocs/op
BenchmarkPat_Param20                      100000             27870 ns/op            5730 B/op        154 allocs/op
BenchmarkTigerTonic_Param20                50000             46162 ns/op           11064 B/op        176 allocs/op
BenchmarkTraffic_Param20                   50000             35076 ns/op            8239 B/op         68 allocs/op
BenchmarkGoji_Param20                     500000              4462 ns/op            1260 B/op          2 allocs/op
BenchmarkKocha_Param20                    500000              5963 ns/op            1839 B/op         27 allocs/op
BenchmarkDenco_Param20                    500000              4210 ns/op            1679 B/op          7 allocs/op
BenchmarkRevel_Param20                    200000             11882 ns/op            4560 B/op         35 allocs/op

BenchmarkGocraftWeb_ParamWrite           1000000              2184 ns/op             682 B/op         10 allocs/op
BenchmarkGorillaMux_ParamWrite            500000              6531 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite           2000000               788 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParamWrite          2000000               860 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParamWrite               200000              8811 ns/op            1285 B/op         16 allocs/op
BenchmarkPat_ParamWrite                   500000              3880 ns/op            1128 B/op         19 allocs/op
BenchmarkTigerTonic_ParamWrite            500000              5439 ns/op            1288 B/op         22 allocs/op
BenchmarkTraffic_ParamWrite               200000              9094 ns/op            2459 B/op         27 allocs/op
BenchmarkGoji_ParamWrite                 1000000              1122 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParamWrite                5000000               530 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_ParamWrite                5000000               456 ns/op              50 B/op          2 allocs/op
BenchmarkRevel_ParamWrite                1000000              1893 ns/op             638 B/op          8 allocs/op

BenchmarkGocraftWeb_GithubStatic         1000000              1235 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             48213 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000              84.9 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       20000000              82.9 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21400 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14379 ns/op            3788 B/op         76 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               416 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             81332 ns/op           23360 B/op        172 allocs/op
BenchmarkGoji_GithubStatic               5000000               483 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic             20000000               107 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_GithubStatic             50000000              70.3 ns/op               0 B/op          0 allocs/op
BenchmarkRevel_GithubStatic              2000000               822 ns/op             180 B/op          4 allocs/op

BenchmarkGocraftWeb_GithubParam          1000000              2537 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_GithubParam            50000             33002 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          1000000              1011 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GithubParam         1000000              1131 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GithubParam              100000             28306 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10406 ns/op            2625 B/op         56 allocs/op
BenchmarkTigerTonic_GithubParam           500000              5929 ns/op            1289 B/op         25 allocs/op
BenchmarkTraffic_GithubParam               50000             38240 ns/op            7148 B/op         60 allocs/op
BenchmarkGoji_GithubParam                1000000              1909 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GithubParam               2000000               954 ns/op             132 B/op          5 allocs/op
BenchmarkDenco_GithubParam               2000000               801 ns/op             116 B/op          3 allocs/op
BenchmarkRevel_GithubParam               1000000              2449 ns/op             737 B/op         10 allocs/op

BenchmarkGocraftWeb_GithubAll               5000            499586 ns/op          136386 B/op       1914 allocs/op
BenchmarkGorillaMux_GithubAll                100          21086844 ns/op          153297 B/op       1419 allocs/op
BenchmarkHttpRouter_GithubAll              10000            173307 ns/op           57349 B/op        347 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            189550 ns/op           57347 B/op        347 allocs/op
BenchmarkMartini_GithubAll                   100          12835245 ns/op          245307 B/op       2940 allocs/op
BenchmarkPat_GithubAll                       500           5873958 ns/op         1588321 B/op      32572 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1141631 ns/op          218212 B/op       5582 allocs/op
BenchmarkTraffic_GithubAll                   100          18024869 ns/op         3173886 B/op      24944 allocs/op
BenchmarkGoji_GithubAll                     2000            857860 ns/op           57307 B/op        347 allocs/op
BenchmarkKocha_GithubAll                   10000            187317 ns/op           24045 B/op        847 allocs/op
BenchmarkDenco_GithubAll                   10000            148541 ns/op           21340 B/op        508 allocs/op
BenchmarkRevel_GithubAll                    5000            484703 ns/op          130990 B/op       1845 allocs/op

BenchmarkGocraftWeb_GPlusStatic          1000000              1152 ns/op             297 B/op          6 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4332 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              53.1 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        50000000              55.4 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              5876 ns/op             860 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               432 ns/op              99 B/op          2 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               235 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6185 ns/op            1510 B/op         19 allocs/op
BenchmarkGoji_GPlusStatic                5000000               422 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic              50000000              81.0 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_GPlusStatic              50000000              56.6 ns/op               0 B/op          0 allocs/op
BenchmarkRevel_GPlusStatic               2000000               711 ns/op             164 B/op          4 allocs/op

BenchmarkGocraftWeb_GPlusParam           1000000              2123 ns/op             673 B/op          9 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             10570 ns/op             785 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam           2000000               809 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2000000               873 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlusParam               200000             10978 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2652 ns/op             752 B/op         14 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              3616 ns/op             907 B/op         16 allocs/op
BenchmarkTraffic_GPlusParam               200000             10193 ns/op            2040 B/op         23 allocs/op
BenchmarkGoji_GPlusParam                 1000000              1194 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlusParam                5000000               525 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_GPlusParam                5000000               441 ns/op              50 B/op          2 allocs/op
BenchmarkRevel_GPlusParam                1000000              1947 ns/op             656 B/op          8 allocs/op

BenchmarkGocraftWeb_GPlus2Params         1000000              2597 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             28554 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         2000000               928 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        1000000              1074 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_GPlus2Params              50000             36111 ns/op            1318 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000              8676 ns/op            2400 B/op         41 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              6252 ns/op            1391 B/op         25 allocs/op
BenchmarkTraffic_GPlus2Params              50000             32095 ns/op            3622 B/op         35 allocs/op
BenchmarkGoji_GPlus2Params               1000000              1750 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlus2Params              2000000               972 ns/op             132 B/op          5 allocs/op
BenchmarkDenco_GPlus2Params              2000000               796 ns/op             116 B/op          3 allocs/op
BenchmarkRevel_GPlus2Params              1000000              2560 ns/op             751 B/op         10 allocs/op

BenchmarkGocraftWeb_GPlusAll              100000             29228 ns/op            8342 B/op        117 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            181973 ns/op            9725 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll              200000             10089 ns/op            3774 B/op         22 allocs/op
BenchmarkHttpTreeMux_GPlusAll             200000             10912 ns/op            3774 B/op         22 allocs/op
BenchmarkMartini_GPlusAll                  10000            199395 ns/op           15525 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             67253 ns/op           17688 B/op        346 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             61358 ns/op           13331 B/op        289 allocs/op
BenchmarkTraffic_GPlusAll                  10000            219449 ns/op           42095 B/op        447 allocs/op
BenchmarkGoji_GPlusAll                    100000             16714 ns/op            3774 B/op         22 allocs/op
BenchmarkKocha_GPlusAll                   200000              8965 ns/op            1014 B/op         43 allocs/op
BenchmarkDenco_GPlusAll                   500000              7332 ns/op             886 B/op         27 allocs/op
BenchmarkRevel_GPlusAll                   100000             26935 ns/op            7914 B/op        107 allocs/op

BenchmarkGocraftWeb_ParseStatic          1000000              1231 ns/op             313 B/op          6 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8773 ns/op             459 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              41.0 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        20000000              89.5 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              6748 ns/op             860 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1056 ns/op             249 B/op          5 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               348 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9874 ns/op            2391 B/op         25 allocs/op
BenchmarkGoji_ParseStatic                5000000               418 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic              20000000              75.4 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_ParseStatic              50000000              48.9 ns/op               0 B/op          0 allocs/op
BenchmarkRevel_ParseStatic               2000000               788 ns/op             180 B/op          4 allocs/op

BenchmarkGocraftWeb_ParseParam           1000000              2170 ns/op             688 B/op          9 allocs/op
BenchmarkGorillaMux_ParseParam            200000             10147 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam           2000000               773 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_ParseParam          2000000               885 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_ParseParam               200000              9861 ns/op            1186 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3876 ns/op            1196 B/op         20 allocs/op
BenchmarkTigerTonic_ParseParam            500000              3433 ns/op             888 B/op         16 allocs/op
BenchmarkTraffic_ParseParam               200000              9849 ns/op            2326 B/op         25 allocs/op
BenchmarkGoji_ParseParam                 1000000              1352 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParseParam                5000000               482 ns/op              58 B/op          3 allocs/op
BenchmarkDenco_ParseParam                5000000               408 ns/op              50 B/op          2 allocs/op
BenchmarkRevel_ParseParam                1000000              1934 ns/op             654 B/op          8 allocs/op

BenchmarkGocraftWeb_Parse2Params         1000000              2436 ns/op             736 B/op         10 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10406 ns/op             818 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params         2000000               848 ns/op             343 B/op          2 allocs/op
BenchmarkHttpTreeMux_Parse2Params        2000000              1003 ns/op             343 B/op          2 allocs/op
BenchmarkMartini_Parse2Params             200000              9851 ns/op            1219 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3870 ns/op             907 B/op         21 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              5811 ns/op            1294 B/op         25 allocs/op
BenchmarkTraffic_Parse2Params             200000              9925 ns/op            2131 B/op         25 allocs/op
BenchmarkGoji_Parse2Params               1000000              1318 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Parse2Params              2000000               843 ns/op             132 B/op          5 allocs/op
BenchmarkDenco_Parse2Params              5000000               691 ns/op             116 B/op          3 allocs/op
BenchmarkRevel_Parse2Params              1000000              2296 ns/op             721 B/op         10 allocs/op

BenchmarkGocraftWeb_ParseAll               50000             50505 ns/op           14299 B/op        209 allocs/op
BenchmarkGorillaMux_ParseAll                5000            380400 ns/op           17254 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              200000             14626 ns/op            5490 B/op         33 allocs/op
BenchmarkHttpTreeMux_ParseAll             100000             16865 ns/op            5490 B/op         33 allocs/op
BenchmarkMartini_ParseAll                  10000            253485 ns/op           27676 B/op        333 allocs/op
BenchmarkPat_ParseAll                      50000             72984 ns/op           18273 B/op        385 allocs/op
BenchmarkTigerTonic_ParseAll               20000             79999 ns/op           17728 B/op        372 allocs/op
BenchmarkTraffic_ParseAll                  10000            312515 ns/op           70562 B/op        763 allocs/op
BenchmarkGoji_ParseAll                    100000             26797 ns/op            5490 B/op         33 allocs/op
BenchmarkKocha_ParseAll                   200000             11981 ns/op            1159 B/op         54 allocs/op
BenchmarkDenco_ParseAll                   200000              9670 ns/op            1007 B/op         35 allocs/op
BenchmarkRevel_ParseAll                    50000             44246 ns/op           12626 B/op        176 allocs/op

BenchmarkHttpServeMux_StaticAll             2000           1280020 ns/op             104 B/op          8 allocs/op
BenchmarkGocraftWeb_StaticAll              10000            207285 ns/op           49172 B/op        951 allocs/op
BenchmarkGorillaMux_StaticAll                500           5577871 ns/op           72315 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll             100000             25577 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll            100000             25604 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll                   500           4159603 ns/op          145594 B/op       2521 allocs/op
BenchmarkPat_StaticAll                      1000           2091708 ns/op          554219 B/op      11249 allocs/op
BenchmarkTigerTonic_StaticAll              50000             74979 ns/op            7776 B/op        158 allocs/op
BenchmarkTraffic_StaticAll                   100          13998951 ns/op         3796343 B/op      27937 allocs/op
BenchmarkGoji_StaticAll                    10000            101862 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_StaticAll                  100000             26770 ns/op               0 B/op          0 allocs/op
BenchmarkDenco_StaticAll                  100000             15818 ns/op               0 B/op          0 allocs/op
BenchmarkRevel_StaticAll                   10000            144219 ns/op           30735 B/op        633 allocs/op
```
