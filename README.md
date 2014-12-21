Go HTTP Router Benchmark
========================

This benchmark suite aims to compare the performance of HTTP request routers for [Go](https://golang.org) by implementing the routing structure of some real world APIs.
Some of the APIs are slightly adapted, since they can not be implemented 1:1 in some of the routers.

Of course the tested routers can be used for any kind of HTTP request → handler function routing, not only (REST) APIs.


#### Tested routers & frameworks:

 * [Beego](http://beego.me/)
 * [go-json-rest](https://github.com/ant0ine/go-json-rest)
 * [Denco](https://github.com/naoina/denco)
 * [Gocraft Web](https://github.com/gocraft/web)
 * [Goji](https://github.com/zenazn/goji/)
 * [Gorilla Mux](http://www.gorillatoolkit.org/pkg/mux)
 * [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux)
 * [HttpRouter](https://github.com/julienschmidt/httprouter)
 * [HttpTreeMux](https://github.com/dimfeld/httptreemux)
 * [Kocha-urlrouter](https://github.com/naoina/kocha-urlrouter)
 * [Martini](https://github.com/go-martini/martini)
 * [Pat](https://github.com/bmizerany/pat)
 * [Revel](https://github.com/revel/revel)
 * [TigerTonic](https://github.com/rcrowley/go-tigertonic)
 * [Traffic](https://github.com/pilu/traffic)


## Motivation

Go is a great language for web applications. Since the [default *request multiplexer*](http://golang.org/pkg/net/http/#ServeMux) of Go's net/http package is very simple and limited, an accordingly high number of HTTP request routers exist.

Unfortunately, most of the (early) routers use pretty bad routing algorithms. Moreover, many of them are very wasteful with memory allocations, which can become a problem in a language with Garbage Collection like Go, since every (heap) allocation results in more work for the Garbage Collector.

Lately more and more bloated frameworks pop up, outdoing one another in the number of features. This benchmark tries to measure their overhead.

Beware that we are comparing apples to oranges here, we compare feature-rich frameworks to packages with simple routing functionality only. But since we are only interested in decent request routing, I think this is not entirely unfair. The frameworks are configured to do as little additional work as possible.

If you care about performance, this benchmark can maybe help you find the right router, which scales with your application.

Personally, I prefer slim and optimized software, which is why I implemented [HttpRouter](https://github.com/julienschmidt/httprouter), which is also tested here. In fact, this benchmark suite started as part of the packages tests, but was then extended to a generic benchmark suite.
So keep in mind, that I am not completely unbiased :relieved:


## Results

Benchmark System:
 * Intel Core i5 M 580 (4x 2.67GHz)
 * 2x 4 GiB DDR3-1066 RAM
 * go1.2.2 linux/amd64
 * Arch Linux amd64 (Linux Kernel 3.14.6)


### Memory Consumption

Besides the micro-benchmarks, there are 3 sets of benchmarks where we play around with clones of some real-world APIs, and one benchmark with static routes only, to allow a comparison with [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).
The following table shows the memory required only for loading the routing structure for the respective API.
The best 3 values for each test are bold. I'm pretty sure you can detect a pattern :wink:

| Router       | Static    | GitHub     | Google+   | Parse     |
|:-------------|----------:|-----------:|----------:|----------:|
| HttpServeMux |  19280 B  |         -  |        -  |        -  |
| Beego        |  80536 B  |  502552 B  |  25048 B  |  39944 B  |
| Denco        |__11208 B__| __46424 B__|   8840 B  | __5576 B__|
| Gocraft Web  |  55784 B  |   97880 B  |   8952 B  |  12352 B  |
| Goji         |  35088 B  |   58840 B  | __3456 B__|   5952 B  |
| Go-Json-Rest | 145400 B  |  142064 B  |  12104 B  |  16832 B  |
| Gorilla Mux  | 333864 B  |  742488 B  |  34464 B  |  60176 B  |
| HttpRouter   |__24632 B__| __43736 B__| __3128 B__| __5768 B__|
| HttpTreeMux  |  78184 B  |   76672 B  |   7160 B  |   7664 B  |
| Kocha        | 117968 B  |  794864 B  | 129888 B  | 185184 B  |
| Martini      | 317248 B  |  527088 B  |  24848 B  |  46776 B  |
| Pat          |__18168 B__| __16488 B__| __1448 B__| __2072 B__|
| Revel        |  64840 B  |  104112 B  |   9264 B  |  16416 B  |
| TigerTonic   |  84120 B  |  103760 B  |   9960 B  |  10520 B  |
| Traffic      | 299336 B  |  501320 B  |  22208 B  |  43624 B  |

The first place goes to [Pat](https://github.com/bmizerany/pat), followed by [HttpRouter](https://github.com/julienschmidt/httprouter) and [Goji](https://github.com/zenazn/goji/). Now, before everyone starts reading the documentation of Pat, `[SPOILER]` this low memory consumption comes at the price of relatively bad routing performance. The routing structure of Pat is simple - probably too simple. `[/SPOILER]`.

Moreover main memory is cheap and usually not a scarce resource. As long as the router doesn't require Megabytes of memory, it should be no deal breaker. But it gives us a first hint how efficient or wasteful a router works.


### Static Routes

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

The only intention of this benchmark is to allow a comparison with the default router of Go's net/http package, [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is limited to static routes and does not support parameters in the route pattern.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.8% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.

```
BenchmarkHttpServeMux_StaticAll     2000           1219984 ns/op             104 B/op          8 allocs/op

BenchmarkBeego_StaticAll            1000           2068455 ns/op          521215 B/op      15181 allocs/op
BenchmarkDenco_StaticAll          100000             15893 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_StaticAll      10000            210764 ns/op           49245 B/op        951 allocs/op
BenchmarkGoji_StaticAll            10000            101666 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_StaticAll       2000           1252887 ns/op          183212 B/op       4130 allocs/op
BenchmarkGorillaMux_StaticAll        500           5634590 ns/op           72414 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll     100000             27286 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll     50000             27069 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_StaticAll          100000             26707 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll           500           4238591 ns/op          145816 B/op       2522 allocs/op
BenchmarkPat_StaticAll              1000           2374322 ns/op          554345 B/op      11250 allocs/op
BenchmarkRevel_StaticAll           10000            170411 ns/op           30754 B/op        633 allocs/op
BenchmarkTigerTonic_StaticAll      20000             75609 ns/op            7781 B/op        158 allocs/op
BenchmarkTraffic_StaticAll           100          14243435 ns/op         3798803 B/op      27958 allocs/op
```

### Micro Benchmarks

The following benchmarks measure the cost of some very basic operations.

In the first benchmark, only a single route, containing a parameter, is loaded into the routers. Then a request for a URL matching this pattern is made and the router has to call the respective registered handler function. End.
```
BenchmarkBeego_Param              500000              7608 ns/op            1213 B/op         21 allocs/op
BenchmarkDenco_Param            10000000               265 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_Param        1000000              2046 ns/op             674 B/op          9 allocs/op
BenchmarkGoji_Param              1000000              1054 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Param         200000              8899 ns/op            1804 B/op         30 allocs/op
BenchmarkGorillaMux_Param         500000              5999 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_Param       10000000               187 ns/op              33 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param       2000000               794 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param             5000000               455 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_Param            500000              7795 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_Param               1000000              2425 ns/op             687 B/op         14 allocs/op
BenchmarkRevel_Param             1000000              1820 ns/op             639 B/op          8 allocs/op
BenchmarkTigerTonic_Param         500000              3847 ns/op            1027 B/op         19 allocs/op
BenchmarkTraffic_Param            500000              7608 ns/op            2030 B/op         23 allocs/op
```

Same as before, but now with multiple parameters, all in the same single route. The intention is to see how the routers scale with the number of parameters. The values of the parameters must be passed to the handler function somehow, which requires allocations. Let's see how clever the routers solve this task with a route containing 5 and 20 parameters:
```
BenchmarkBeego_Param5             100000             25083 ns/op            1343 B/op         21 allocs/op
BenchmarkDenco_Param5            1000000              1422 ns/op             490 B/op          4 allocs/op
BenchmarkGocraftWeb_Param5       1000000              3210 ns/op             948 B/op         12 allocs/op
BenchmarkGoji_Param5             1000000              1551 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Param5        200000             13900 ns/op            3291 B/op         41 allocs/op
BenchmarkGorillaMux_Param5        200000             11399 ns/op             916 B/op          7 allocs/op
BenchmarkHttpRouter_Param5       5000000               505 ns/op             163 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param5      1000000              1223 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param5            1000000              1851 ns/op             449 B/op         10 allocs/op
BenchmarkMartini_Param5           100000             16133 ns/op            1317 B/op         13 allocs/op
BenchmarkPat_Param5               500000              4941 ns/op            1494 B/op         25 allocs/op
BenchmarkRevel_Param5            1000000              3231 ns/op             981 B/op         15 allocs/op
BenchmarkTigerTonic_Param5        200000             12315 ns/op            2640 B/op         53 allocs/op
BenchmarkTraffic_Param5           200000             12702 ns/op            2357 B/op         31 allocs/op

BenchmarkBeego_Param20             10000            140964 ns/op            3740 B/op         24 allocs/op
BenchmarkDenco_Param20            500000              4688 ns/op            2054 B/op          6 allocs/op
BenchmarkGocraftWeb_Param20       200000             10244 ns/op            3865 B/op         17 allocs/op
BenchmarkGoji_Param20             500000              4535 ns/op            1260 B/op          2 allocs/op
BenchmarkGoJsonRest_Param20        50000             35934 ns/op           10671 B/op         77 allocs/op
BenchmarkGorillaMux_Param20       100000             21193 ns/op            3313 B/op         10 allocs/op
BenchmarkHttpRouter_Param20      1000000              1531 ns/op             653 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param20      500000              7033 ns/op            2220 B/op          4 allocs/op
BenchmarkKocha_Param20            500000              6429 ns/op            1839 B/op         27 allocs/op
BenchmarkMartini_Param20           50000             70080 ns/op            3713 B/op         16 allocs/op
BenchmarkPat_Param20              500000              4967 ns/op            1494 B/op         25 allocs/op
BenchmarkRevel_Param20            200000             12056 ns/op            4564 B/op         35 allocs/op
BenchmarkTigerTonic_Param20        50000             48100 ns/op           11270 B/op        179 allocs/op
BenchmarkTraffic_Param20           50000             35198 ns/op            8251 B/op         68 allocs/op
```

Now let's see how expensive it is to access a parameter. The handler function reads the value (by the name of the parameter, e.g. with a map lookup; depends on the router) and writes it to our [web scale storage](https://www.youtube.com/watch?v=b2F-DItXtZs) (`/dev/null`).
```
BenchmarkBeego_ParamWrite         200000              9135 ns/op            1654 B/op         26 allocs/op
BenchmarkDenco_ParamWrite        5000000               343 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_ParamWrite   1000000              2202 ns/op             683 B/op         10 allocs/op
BenchmarkGoji_ParamWrite         1000000              1133 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_ParamWrite    200000             10767 ns/op            2285 B/op         35 allocs/op
BenchmarkGorillaMux_ParamWrite    500000              6305 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite  10000000               282 ns/op              33 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParamWrite  2000000               875 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParamWrite        5000000               530 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_ParamWrite       200000              9066 ns/op            1287 B/op         16 allocs/op
BenchmarkPat_ParamWrite           500000              4061 ns/op            1129 B/op         19 allocs/op
BenchmarkRevel_ParamWrite        1000000              1910 ns/op             639 B/op          8 allocs/op
BenchmarkTigerTonic_ParamWrite    500000              6150 ns/op            1485 B/op         25 allocs/op
BenchmarkTraffic_ParamWrite       200000              9141 ns/op            2464 B/op         27 allocs/op
```

### [Parse.com](https://parse.com/docs/rest#summary)

Enough of the micro benchmark stuff. Let's play a bit with real APIs. In the first set of benchmarks, we use a clone of the structure of [Parse](https://parse.com)'s decent medium-sized REST API, consisting of 26 routes.

The tasks are 1.) routing a static URL (no parameters), 2.) routing a URL containing 1 parameter, 3.) same with 2 parameters, 4.) route all of the routes once (like the StaticAll benchmark, but the routes now contain parameters).

Worth noting is, that the requested route might be a good case for some routing algorithms, while it is a bad case for another algorithm. The values might vary slightly depending on the selected route.

```
BenchmarkBeego_ParseStatic                500000              5042 ns/op            1295 B/op         22 allocs/op
BenchmarkDenco_ParseStatic              50000000              50.5 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParseStatic          1000000              1250 ns/op             314 B/op          6 allocs/op
BenchmarkGoji_ParseStatic                5000000               425 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_ParseStatic           500000              7258 ns/op            1147 B/op         26 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8353 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              49.3 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        20000000              83.8 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic              20000000              79.7 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              6832 ns/op             862 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1069 ns/op             249 B/op          5 allocs/op
BenchmarkRevel_ParseStatic               2000000               804 ns/op             180 B/op          4 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               351 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9666 ns/op            2396 B/op         25 allocs/op

BenchmarkBeego_ParseParam                 200000             11223 ns/op            1827 B/op         35 allocs/op
BenchmarkDenco_ParseParam               10000000               310 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_ParseParam           1000000              2231 ns/op             690 B/op          9 allocs/op
BenchmarkGoji_ParseParam                 1000000              1361 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_ParseParam            200000              9321 ns/op            1809 B/op         30 allocs/op
BenchmarkGorillaMux_ParseParam            200000              9629 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam          10000000               250 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParseParam          2000000               890 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParseParam                5000000               489 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_ParseParam               200000              9956 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3997 ns/op            1197 B/op         20 allocs/op
BenchmarkRevel_ParseParam                1000000              1971 ns/op             654 B/op          8 allocs/op
BenchmarkTigerTonic_ParseParam            500000              4162 ns/op            1084 B/op         19 allocs/op
BenchmarkTraffic_ParseParam               200000              9741 ns/op            2330 B/op         25 allocs/op

BenchmarkBeego_Parse2Params               200000             13036 ns/op            1988 B/op         35 allocs/op
BenchmarkDenco_Parse2Params              5000000               584 ns/op              99 B/op          2 allocs/op
BenchmarkGocraftWeb_Parse2Params         1000000              2483 ns/op             737 B/op         10 allocs/op
BenchmarkGoji_Parse2Params               1000000              1319 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Parse2Params          200000             10570 ns/op            2168 B/op         33 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10248 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params        10000000               285 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_Parse2Params        2000000              1001 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Parse2Params              2000000               865 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_Parse2Params             200000             10282 ns/op            1220 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3935 ns/op             908 B/op         21 allocs/op
BenchmarkRevel_Parse2Params              1000000              2352 ns/op             722 B/op         10 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              6908 ns/op            1488 B/op         28 allocs/op
BenchmarkTraffic_Parse2Params             200000             10809 ns/op            2135 B/op         25 allocs/op

BenchmarkBeego_ParseAll                    10000            272314 ns/op           40500 B/op        777 allocs/op
BenchmarkDenco_ParseAll                   500000              7582 ns/op             733 B/op         19 allocs/op
BenchmarkGocraftWeb_ParseAll               50000             51064 ns/op           14320 B/op        210 allocs/op
BenchmarkGoji_ParseAll                    100000             26931 ns/op            5491 B/op         33 allocs/op
BenchmarkGoJsonRest_ParseAll               10000            229236 ns/op           41659 B/op        759 allocs/op
BenchmarkGorillaMux_ParseAll                5000            369136 ns/op           17276 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              500000              5391 ns/op             665 B/op         16 allocs/op
BenchmarkHttpTreeMux_ParseAll             100000             17171 ns/op            5491 B/op         33 allocs/op
BenchmarkKocha_ParseAll                   200000             11928 ns/op            1160 B/op         54 allocs/op
BenchmarkMartini_ParseAll                  10000            257492 ns/op           27716 B/op        333 allocs/op
BenchmarkPat_ParseAll                      50000             74900 ns/op           18296 B/op        385 allocs/op
BenchmarkRevel_ParseAll                    50000             44436 ns/op           12647 B/op        176 allocs/op
BenchmarkTigerTonic_ParseAll               20000             92706 ns/op           20873 B/op        420 allocs/op
BenchmarkTraffic_ParseAll                  10000            308038 ns/op           70702 B/op        763 allocs/op
```


### [GitHub](http://developer.github.com/v3/)

The GitHub API is rather large, consisting of 203 routes. The tasks are basically the same as in the benchmarks before.

```
BenchmarkBeego_GithubStatic               500000              5462 ns/op            1197 B/op         38 allocs/op
BenchmarkDenco_GithubStatic             50000000              70.3 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GithubStatic         1000000              1245 ns/op             314 B/op          6 allocs/op
BenchmarkGoji_GithubStatic               5000000               481 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_GithubStatic          500000              7471 ns/op            1163 B/op         26 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             48127 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000               103 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       20000000              84.5 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic             20000000               104 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21472 ns/op             862 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14441 ns/op            3789 B/op         76 allocs/op
BenchmarkRevel_GithubStatic              2000000               814 ns/op             180 B/op          4 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               422 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             82883 ns/op           23368 B/op        172 allocs/op

BenchmarkBeego_GithubParam                 50000             63657 ns/op            3038 B/op         57 allocs/op
BenchmarkDenco_GithubParam               5000000               657 ns/op              99 B/op          2 allocs/op
BenchmarkGocraftWeb_GithubParam          1000000              2613 ns/op             737 B/op         10 allocs/op
BenchmarkGoji_GithubParam                1000000              1899 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GithubParam           200000             10858 ns/op            2193 B/op         33 allocs/op
BenchmarkGorillaMux_GithubParam            50000             32591 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          5000000               480 ns/op              98 B/op          1 allocs/op
BenchmarkHttpTreeMux_GithubParam         1000000              1155 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GithubParam               2000000               972 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_GithubParam              100000             28348 ns/op            1221 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10474 ns/op            2628 B/op         56 allocs/op
BenchmarkRevel_GithubParam               1000000              2474 ns/op             739 B/op         10 allocs/op
BenchmarkTigerTonic_GithubParam           500000              6560 ns/op            1484 B/op         28 allocs/op
BenchmarkTraffic_GithubParam               50000             38388 ns/op            7153 B/op         60 allocs/op

BenchmarkBeego_GithubAll                      50          33412384 ns/op          519818 B/op      11209 allocs/op
BenchmarkDenco_GithubAll                   10000            124646 ns/op           20058 B/op        341 allocs/op
BenchmarkGocraftWeb_GithubAll               5000            508037 ns/op          136608 B/op       1915 allocs/op
BenchmarkGoji_GithubAll                     2000            857689 ns/op           57346 B/op        347 allocs/op
BenchmarkGoJsonRest_GithubAll               1000           2136712 ns/op          407995 B/op       6558 allocs/op
BenchmarkGorillaMux_GithubAll                100          21210657 ns/op          153505 B/op       1420 allocs/op
BenchmarkHttpRouter_GithubAll              20000             81445 ns/op           14101 B/op        169 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            195386 ns/op           57349 B/op        347 allocs/op
BenchmarkKocha_GithubAll                   10000            186740 ns/op           24070 B/op        847 allocs/op
BenchmarkMartini_GithubAll                   100          12863026 ns/op          245797 B/op       2943 allocs/op
BenchmarkPat_GithubAll                       500           5945074 ns/op         1589140 B/op      32575 allocs/op
BenchmarkRevel_GithubAll                    5000            477564 ns/op          131274 B/op       1847 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1265563 ns/op          251049 B/op       6087 allocs/op
BenchmarkTraffic_GithubAll                   100          18400453 ns/op         3176285 B/op      24959 allocs/op
```

### [Google+](https://developers.google.com/+/api/latest/)

Last but not least the Google+ API, consisting of 13 routes. In reality this is just a subset of a much larger API.

```
BenchmarkBeego_GPlusStatic               1000000              3491 ns/op             853 B/op         18 allocs/op
BenchmarkDenco_GPlusStatic              50000000              42.5 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlusStatic          1000000              1171 ns/op             297 B/op          6 allocs/op
BenchmarkGoji_GPlusStatic                5000000               336 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_GPlusStatic           500000              7197 ns/op            1147 B/op         26 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4044 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              47.8 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        50000000              45.4 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic              50000000              72.1 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              5912 ns/op             862 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               441 ns/op              99 B/op          2 allocs/op
BenchmarkRevel_GPlusStatic               5000000               728 ns/op             164 B/op          4 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               237 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6175 ns/op            1513 B/op         19 allocs/op

BenchmarkBeego_GPlusParam                 200000             10548 ns/op            1279 B/op         23 allocs/op
BenchmarkDenco_GPlusParam                5000000               312 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_GPlusParam           1000000              2169 ns/op             674 B/op          9 allocs/op
BenchmarkGoji_GPlusParam                 1000000              1203 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlusParam            200000              9500 ns/op            1823 B/op         30 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             10240 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam          10000000               300 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2000000               894 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlusParam                5000000               527 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_GPlusParam               200000             11221 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2731 ns/op             753 B/op         14 allocs/op
BenchmarkRevel_GPlusParam                1000000              1954 ns/op             656 B/op          8 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              4309 ns/op            1103 B/op         19 allocs/op
BenchmarkTraffic_GPlusParam               200000              9977 ns/op            2044 B/op         23 allocs/op

BenchmarkBeego_GPlus2Params               200000             13947 ns/op            1340 B/op         23 allocs/op
BenchmarkDenco_GPlus2Params              5000000               647 ns/op              99 B/op          2 allocs/op
BenchmarkGocraftWeb_GPlus2Params         1000000              2667 ns/op             736 B/op         10 allocs/op
BenchmarkGoji_GPlus2Params               1000000              1771 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlus2Params          200000             11081 ns/op            2198 B/op         33 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             28219 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         5000000               355 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        1000000              1103 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlus2Params              2000000               989 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_GPlus2Params              50000             36649 ns/op            1320 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000              8776 ns/op            2403 B/op         41 allocs/op
BenchmarkRevel_GPlus2Params              1000000              2557 ns/op             752 B/op         10 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              7039 ns/op            1586 B/op         28 allocs/op
BenchmarkTraffic_GPlus2Params              50000             31630 ns/op            3629 B/op         35 allocs/op

BenchmarkBeego_GPlusAll                    10000            154279 ns/op           16504 B/op        299 allocs/op
BenchmarkDenco_GPlusAll                   500000              5955 ns/op             698 B/op         16 allocs/op
BenchmarkGocraftWeb_GPlusAll              100000             30077 ns/op            8354 B/op        117 allocs/op
BenchmarkGoji_GPlusAll                    100000             17040 ns/op            3775 B/op         22 allocs/op
BenchmarkGoJsonRest_GPlusAll               10000            127173 ns/op           24126 B/op        402 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            180824 ns/op            9732 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll             1000000              3870 ns/op             660 B/op         11 allocs/op
BenchmarkHttpTreeMux_GPlusAll             200000             11155 ns/op            3775 B/op         22 allocs/op
BenchmarkKocha_GPlusAll                   200000              8992 ns/op            1015 B/op         43 allocs/op
BenchmarkMartini_GPlusAll                  10000            204658 ns/op           15553 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             68309 ns/op           17711 B/op        346 allocs/op
BenchmarkRevel_GPlusAll                   100000             27246 ns/op            7930 B/op        107 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             70209 ns/op           15493 B/op        322 allocs/op
BenchmarkTraffic_GPlusAll                  10000            215859 ns/op           42173 B/op        447 allocs/op
```


## Conclusions
First of all, there is no reason to use net/http's default [ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is very limited and does not have especially good performance. There are enough alternatives coming in every flavor, choose the one you like best.

Secondly, the broad range of functions of some of the frameworks comes at a high price in terms of performance. For example Martini has great flexibility, but very bad performance. Martini has the worst performance of all tested routers in a lot of the benchmarks. Beego seems to have some scalability problems and easily defeats Martini with even worse performance, when the number of parameters or routes is high. I really hope, that the routing of these packages can be optimized. I think the Go-ecosystem needs great feature-rich frameworks like these.

Last but not least, we have to determine the performance champion.

Denco and its predecessor Kocha-urlrouter seem to have great performance, but are not convenient to use as a router for the net/http package. A lot of extra work is necessary to use it as a http.Handler. [The README of Denco claims](https://github.com/naoina/denco/blob/b03dbb499269a597afd0db715d408ebba1329d04/README.md), that the package is not intended as a replacement for [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).

[Goji](https://github.com/zenazn/goji/) looks very decent. It has great performance while also having a great range of features, more than any other router / framework in the top group.

Currently no router can beat the performance of the [HttpRouter](https://github.com/julienschmidt/httprouter) package, which currently dominates nearly all benchmarks.

In the end, performance can not be the (only) criterion for choosing a router. Play around a bit with some of the routers, and choose the one you like best.
