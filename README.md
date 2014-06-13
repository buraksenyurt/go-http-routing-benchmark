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
 * [Kocha-urlrouter](github.com/naoina/kocha-urlrouter)
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
| HttpServeMux |__18256 B__|         -  |        -  |        -  |
| Beego        |  79992 B  |  502168 B  |  25128 B  |  39400 B  |
| Denco        |__15872 B__| __50016 B__|  10432 B  |   6080 B  |
| Gocraft Web  |  55736 B  |   95080 B  |   7768 B  |  12352 B  |
| Goji         |  35008 B  |   60856 B  | __3456 B__| __5952 B__|
| Go-Json-Rest | 146440 B  |  142112 B  |  12248 B  |  14960 B  |
| Gorilla Mux  | 332136 B  |  742280 B  |  34464 B  |  60176 B  |
| HttpRouter   |  29880 B  | __43736 B__| __3128 B__| __5768 B__|
| HttpTreeMux  |  74440 B  |   76672 B  |   7160 B  |   7664 B  |
| Kocha        | 117808 B  |  795344 B  | 129808 B  | 185104 B  |
| Martini      | 317008 B  |  530032 B  |  24848 B  |  48584 B  |
| Pat          |__17624 B__| __16488 B__| __1448 B__| __2072 B__|
| Revel        |  66088 B  |  104320 B  |   9264 B  |  14688 B  |
| TigerTonic   |  84296 B  |  103840 B  |  10040 B  |  10520 B  |
| Traffic      | 300440 B  |  500920 B  |  24480 B  |  43336 B  |

The first place goes to [Pat](https://github.com/bmizerany/pat), followed by [HttpRouter](https://github.com/julienschmidt/httprouter) and [Goji](https://github.com/zenazn/goji/). Now, before everyone starts reading the documentation of Pat, `[SPOILER]` this low memory consumption comes at the price of relatively bad routing performance. The routing structure of Pat is simple - probably too simple. `[/SPOILER]`.

Moreover main memory is cheap and usually not a scarce resource. As long as the router doesn't require Megabytes of memory, it should be no deal breaker. But it gives us a first hint how efficient or wasteful a router works.


### Static Routes

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

The only intention of this benchmark is to allow a comparison with the default router of Go's net/http package, [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is limited to static routes and does not support parameters in the route pattern.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.8% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.

```
BenchmarkHttpServeMux_StaticAll     2000           1201290 ns/op             104 B/op          8 allocs/op

BenchmarkBeego_StaticAll            1000           2063983 ns/op          521120 B/op      15180 allocs/op
BenchmarkDenco_StaticAll          100000             15445 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_StaticAll      10000            212954 ns/op           49238 B/op        951 allocs/op
BenchmarkGoji_StaticAll            10000            102123 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_StaticAll       2000           1286144 ns/op          183132 B/op       4129 allocs/op
BenchmarkGorillaMux_StaticAll        500           5628269 ns/op           72403 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll     100000             25028 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll    100000             25080 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_StaticAll          100000             26600 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll           500           4176818 ns/op          145788 B/op       2521 allocs/op
BenchmarkPat_StaticAll              1000           2177115 ns/op          554319 B/op      11250 allocs/op
BenchmarkRevel_StaticAll           10000            147621 ns/op           30751 B/op        633 allocs/op
BenchmarkTigerTonic_StaticAll      20000             76099 ns/op            7781 B/op        158 allocs/op
BenchmarkTraffic_StaticAll           100          14340193 ns/op         3798582 B/op      27957 allocs/op
```

### Micro Benchmarks

The following benchmarks measure the cost of some very basic operations.

In the first benchmark, only a single route, containing a parameter, is loaded into the routers. Then a request for a URL matching this pattern is made and the router has to call the respective registered handler function. End.
```
BenchmarkBeego_Param              200000              7526 ns/op            1213 B/op         21 allocs/op
BenchmarkDenco_Param            10000000               264 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_Param        1000000              2057 ns/op             674 B/op          9 allocs/op
BenchmarkGoji_Param              1000000              1069 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Param         200000              9139 ns/op            1804 B/op         30 allocs/op
BenchmarkGorillaMux_Param         500000              5846 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_Param       10000000               187 ns/op              33 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param       2000000               802 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param             5000000               452 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_Param            500000              7719 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_Param               1000000              2344 ns/op             687 B/op         14 allocs/op
BenchmarkRevel_Param             1000000              1822 ns/op             639 B/op          8 allocs/op
BenchmarkTigerTonic_Param         500000              3852 ns/op            1027 B/op         19 allocs/op
BenchmarkTraffic_Param            500000              7531 ns/op            2030 B/op         23 allocs/op
```

Same as before, but now with multiple parameters, all in the same single route. The intention is to see how the routers scale with the number of parameters. The values of the parameters must be passed to the handler function somehow, which requires allocations. Let's see how clever the routers solve this task with a route containing 5 and 20 parameters:
```
BenchmarkBeego_Param5             100000             25516 ns/op            1343 B/op         21 allocs/op
BenchmarkDenco_Param5            1000000              1460 ns/op             490 B/op          4 allocs/op
BenchmarkGocraftWeb_Param5       1000000              3191 ns/op             948 B/op         12 allocs/op
BenchmarkGoji_Param5             1000000              1548 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Param5        200000             14113 ns/op            3291 B/op         41 allocs/op
BenchmarkGorillaMux_Param5        200000             11210 ns/op             916 B/op          7 allocs/op
BenchmarkHttpRouter_Param5       5000000               494 ns/op             163 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param5      1000000              1235 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param5            1000000              1857 ns/op             449 B/op         10 allocs/op
BenchmarkMartini_Param5           100000             16115 ns/op            1317 B/op         13 allocs/op
BenchmarkPat_Param5               500000              4981 ns/op            1494 B/op         25 allocs/op
BenchmarkRevel_Param5            1000000              3251 ns/op             981 B/op         15 allocs/op
BenchmarkTigerTonic_Param5        200000             12353 ns/op            2640 B/op         53 allocs/op
BenchmarkTraffic_Param5           200000             12632 ns/op            2357 B/op         31 allocs/op

BenchmarkBeego_Param20             10000            137161 ns/op            3741 B/op         24 allocs/op
BenchmarkDenco_Param20            500000              4694 ns/op            2054 B/op          6 allocs/op
BenchmarkGocraftWeb_Param20       200000             10268 ns/op            3865 B/op         17 allocs/op
BenchmarkGoji_Param20             500000              4577 ns/op            1261 B/op          2 allocs/op
BenchmarkGoJsonRest_Param20        50000             36797 ns/op           10673 B/op         77 allocs/op
BenchmarkGorillaMux_Param20       100000             20776 ns/op            3313 B/op         10 allocs/op
BenchmarkHttpRouter_Param20      1000000              1522 ns/op             653 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param20      500000              7017 ns/op            2220 B/op          4 allocs/op
BenchmarkKocha_Param20            500000              5995 ns/op            1839 B/op         27 allocs/op
BenchmarkMartini_Param20           50000             68846 ns/op            3714 B/op         16 allocs/op
BenchmarkPat_Param20              500000              4994 ns/op            1494 B/op         25 allocs/op
BenchmarkRevel_Param20            200000             12094 ns/op            4564 B/op         35 allocs/op
BenchmarkTigerTonic_Param20        50000             47878 ns/op           11270 B/op        179 allocs/op
BenchmarkTraffic_Param20           50000             34868 ns/op            8251 B/op         68 allocs/op
```

Now let's see how expensive it is to access a parameter. The handler function reads the value (by the name of the parameter, e.g. with a map lookup; depends on the router) and writes it to our [web scale storage](https://www.youtube.com/watch?v=b2F-DItXtZs) (`/dev/null`).
```
BenchmarkBeego_ParamWrite         200000              9030 ns/op            1654 B/op         26 allocs/op
BenchmarkDenco_ParamWrite        5000000               341 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_ParamWrite   1000000              2201 ns/op             683 B/op         10 allocs/op
BenchmarkGoji_ParamWrite         1000000              1129 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_ParamWrite    200000             10914 ns/op            2286 B/op         35 allocs/op
BenchmarkGorillaMux_ParamWrite    500000              6152 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite  10000000               247 ns/op              33 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParamWrite  2000000               891 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParamWrite        5000000               533 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_ParamWrite       200000              8971 ns/op            1287 B/op         16 allocs/op
BenchmarkPat_ParamWrite           500000              4002 ns/op            1129 B/op         19 allocs/op
BenchmarkRevel_ParamWrite        1000000              1880 ns/op             639 B/op          8 allocs/op
BenchmarkTigerTonic_ParamWrite    500000              6135 ns/op            1485 B/op         25 allocs/op
BenchmarkTraffic_ParamWrite       200000              9122 ns/op            2464 B/op         27 allocs/op
```

### [Parse.com](https://parse.com/docs/rest#summary)

Enough of the micro benchmark stuff. Let's play a bit with real APIs. In the first set of benchmarks, we use a clone of the structure of [Parse](https://parse.com)'s decent medium-sized REST API, consisting of 26 routes.

The tasks are 1.) routing a static URL (no parameters), 2.) routing a URL containing 1 parameter, 3.) same with 2 parameters, 4.) route all of the routes once (like the StaticAll benchmark, but the routes now contain parameters).

Worth noting is, that the requested route might be a good case for some routing algorithms, while it is a bad case for another algorithm. The values might vary slightly depending on the selected route.

```
BenchmarkBeego_ParseStatic                500000              4994 ns/op            1295 B/op         22 allocs/op
BenchmarkDenco_ParseStatic              50000000              49.2 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParseStatic          1000000              1254 ns/op             314 B/op          6 allocs/op
BenchmarkGoji_ParseStatic                5000000               425 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_ParseStatic           500000              7520 ns/op            1147 B/op         26 allocs/op
BenchmarkGorillaMux_ParseStatic           200000              8193 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              47.6 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        20000000              82.4 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic              20000000              77.9 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              6934 ns/op             862 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1072 ns/op             249 B/op          5 allocs/op
BenchmarkRevel_ParseStatic               2000000               804 ns/op             180 B/op          4 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               352 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000              9600 ns/op            2396 B/op         25 allocs/op

BenchmarkBeego_ParseParam                 200000             11286 ns/op            1827 B/op         35 allocs/op
BenchmarkDenco_ParseParam               10000000               303 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_ParseParam           1000000              2210 ns/op             689 B/op          9 allocs/op
BenchmarkGoji_ParseParam                 1000000              1364 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_ParseParam            200000              9470 ns/op            1809 B/op         30 allocs/op
BenchmarkGorillaMux_ParseParam            200000              9505 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam          10000000               250 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParseParam          2000000               891 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParseParam                5000000               508 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_ParseParam               200000              9881 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              3900 ns/op            1197 B/op         20 allocs/op
BenchmarkRevel_ParseParam                1000000              1992 ns/op             655 B/op          8 allocs/op
BenchmarkTigerTonic_ParseParam            500000              4203 ns/op            1084 B/op         19 allocs/op
BenchmarkTraffic_ParseParam               200000              9704 ns/op            2329 B/op         25 allocs/op

BenchmarkBeego_Parse2Params               200000             13104 ns/op            1988 B/op         35 allocs/op
BenchmarkDenco_Parse2Params              5000000               586 ns/op              99 B/op          2 allocs/op
BenchmarkGocraftWeb_Parse2Params         1000000              2508 ns/op             736 B/op         10 allocs/op
BenchmarkGoji_Parse2Params               1000000              1338 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Parse2Params          200000             10920 ns/op            2167 B/op         33 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             10274 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params        10000000               288 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_Parse2Params        1000000              1008 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Parse2Params              2000000               863 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_Parse2Params             200000             10188 ns/op            1220 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              3868 ns/op             908 B/op         21 allocs/op
BenchmarkRevel_Parse2Params              1000000              2375 ns/op             722 B/op         10 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              6629 ns/op            1488 B/op         28 allocs/op
BenchmarkTraffic_Parse2Params             200000              9847 ns/op            2135 B/op         25 allocs/op

BenchmarkBeego_ParseAll                    10000            271131 ns/op           40497 B/op        777 allocs/op
BenchmarkDenco_ParseAll                   500000              7808 ns/op             733 B/op         19 allocs/op
BenchmarkGocraftWeb_ParseAll               50000             51273 ns/op           14319 B/op        210 allocs/op
BenchmarkGoji_ParseAll                    100000             27420 ns/op            5491 B/op         33 allocs/op
BenchmarkGoJsonRest_ParseAll               10000            237745 ns/op           41644 B/op        759 allocs/op
BenchmarkGorillaMux_ParseAll                5000            367092 ns/op           17274 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              500000              5382 ns/op             665 B/op         16 allocs/op
BenchmarkHttpTreeMux_ParseAll             100000             17190 ns/op            5491 B/op         33 allocs/op
BenchmarkKocha_ParseAll                   200000             12000 ns/op            1160 B/op         54 allocs/op
BenchmarkMartini_ParseAll                  10000            255956 ns/op           27712 B/op        333 allocs/op
BenchmarkPat_ParseAll                      50000             73982 ns/op           18294 B/op        385 allocs/op
BenchmarkRevel_ParseAll                    50000             45078 ns/op           12646 B/op        176 allocs/op
BenchmarkTigerTonic_ParseAll               20000             92848 ns/op           20871 B/op        420 allocs/op
BenchmarkTraffic_ParseAll                  10000            307768 ns/op           70689 B/op        763 allocs/op
```


### [GitHub](http://developer.github.com/v3/)

The GitHub API is rather large, consisting of 203 routes. The tasks are basically the same as in the benchmarks before.

```
BenchmarkBeego_GithubStatic               500000              5376 ns/op            1197 B/op         38 allocs/op
BenchmarkDenco_GithubStatic             50000000              72.4 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GithubStatic         1000000              1234 ns/op             314 B/op          6 allocs/op
BenchmarkGoji_GithubStatic               5000000               471 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_GithubStatic          500000              7619 ns/op            1163 B/op         26 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             46661 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000              98.8 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       20000000              83.0 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic             20000000               101 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             21218 ns/op             862 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 200000             14353 ns/op            3789 B/op         76 allocs/op
BenchmarkRevel_GithubStatic              2000000               821 ns/op             180 B/op          4 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               417 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             81277 ns/op           23368 B/op        172 allocs/op

BenchmarkBeego_GithubParam                 50000             61400 ns/op            3038 B/op         57 allocs/op
BenchmarkDenco_GithubParam               5000000               668 ns/op              99 B/op          2 allocs/op
BenchmarkGocraftWeb_GithubParam          1000000              2572 ns/op             737 B/op         10 allocs/op
BenchmarkGoji_GithubParam                1000000              1897 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GithubParam           200000             11114 ns/op            2193 B/op         33 allocs/op
BenchmarkGorillaMux_GithubParam            50000             33390 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          5000000               477 ns/op              98 B/op          1 allocs/op
BenchmarkHttpTreeMux_GithubParam         1000000              1132 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GithubParam               2000000               948 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_GithubParam              100000             28201 ns/op            1221 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             10322 ns/op            2628 B/op         56 allocs/op
BenchmarkRevel_GithubParam               1000000              2457 ns/op             739 B/op         10 allocs/op
BenchmarkTigerTonic_GithubParam           500000              6557 ns/op            1484 B/op         28 allocs/op
BenchmarkTraffic_GithubParam               50000             38129 ns/op            7153 B/op         60 allocs/op

BenchmarkBeego_GithubAll                      50          33276310 ns/op          519865 B/op      11210 allocs/op
BenchmarkDenco_GithubAll                   10000            126465 ns/op           20057 B/op        341 allocs/op
BenchmarkGocraftWeb_GithubAll               5000            503853 ns/op          136608 B/op       1915 allocs/op
BenchmarkGoji_GithubAll                     2000            861253 ns/op           57348 B/op        347 allocs/op
BenchmarkGoJsonRest_GithubAll               1000           2173548 ns/op          407929 B/op       6558 allocs/op
BenchmarkGorillaMux_GithubAll                100          20876261 ns/op          153518 B/op       1420 allocs/op
BenchmarkHttpRouter_GithubAll              20000             80131 ns/op           14101 B/op        169 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            192816 ns/op           57351 B/op        347 allocs/op
BenchmarkKocha_GithubAll                   10000            183658 ns/op           24070 B/op        847 allocs/op
BenchmarkMartini_GithubAll                   100          12670407 ns/op          245781 B/op       2943 allocs/op
BenchmarkPat_GithubAll                       500           5940985 ns/op         1589223 B/op      32575 allocs/op
BenchmarkRevel_GithubAll                    5000            472542 ns/op          131276 B/op       1847 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1250034 ns/op          251056 B/op       6087 allocs/op
BenchmarkTraffic_GithubAll                   100          17813831 ns/op         3176162 B/op      24958 allocs/op
```

### [Google+](https://developers.google.com/+/api/latest/)

Last but not least the Google+ API, consisting of 13 routes. In reality this is just a subset of a much larger API.

```
BenchmarkBeego_GPlusStatic               1000000              3398 ns/op             853 B/op         18 allocs/op
BenchmarkDenco_GPlusStatic              50000000              45.6 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlusStatic          1000000              1166 ns/op             297 B/op          6 allocs/op
BenchmarkGoji_GPlusStatic                5000000               324 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_GPlusStatic           500000              7346 ns/op            1147 B/op         26 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4024 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              47.6 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        50000000              45.2 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic              50000000              71.3 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              5983 ns/op             862 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               437 ns/op              99 B/op          2 allocs/op
BenchmarkRevel_GPlusStatic               5000000               727 ns/op             164 B/op          4 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               237 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              6028 ns/op            1513 B/op         19 allocs/op

BenchmarkBeego_GPlusParam                 200000             10409 ns/op            1279 B/op         23 allocs/op
BenchmarkDenco_GPlusParam                5000000               312 ns/op              33 B/op          1 allocs/op
BenchmarkGocraftWeb_GPlusParam           1000000              2162 ns/op             674 B/op          9 allocs/op
BenchmarkGoji_GPlusParam                 1000000              1210 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlusParam            200000              9681 ns/op            1822 B/op         30 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             10115 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam          10000000               294 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlusParam          2000000               908 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlusParam                5000000               533 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_GPlusParam               200000             11088 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              2665 ns/op             753 B/op         14 allocs/op
BenchmarkRevel_GPlusParam                1000000              1946 ns/op             656 B/op          8 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              4278 ns/op            1103 B/op         19 allocs/op
BenchmarkTraffic_GPlusParam               200000             10078 ns/op            2044 B/op         23 allocs/op

BenchmarkBeego_GPlus2Params               200000             13594 ns/op            1340 B/op         23 allocs/op
BenchmarkDenco_GPlus2Params              5000000               672 ns/op              99 B/op          2 allocs/op
BenchmarkGocraftWeb_GPlus2Params         1000000              2631 ns/op             737 B/op         10 allocs/op
BenchmarkGoji_GPlus2Params               1000000              1774 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlus2Params          200000             11283 ns/op            2197 B/op         33 allocs/op
BenchmarkGorillaMux_GPlus2Params          100000             28169 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         5000000               354 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        1000000              1086 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlus2Params              2000000               967 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_GPlus2Params              50000             35769 ns/op            1320 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000              8645 ns/op            2403 B/op         41 allocs/op
BenchmarkRevel_GPlus2Params              1000000              2534 ns/op             752 B/op         10 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              6957 ns/op            1586 B/op         28 allocs/op
BenchmarkTraffic_GPlus2Params              50000             31056 ns/op            3629 B/op         35 allocs/op

BenchmarkBeego_GPlusAll                    10000            153642 ns/op           16505 B/op        299 allocs/op
BenchmarkDenco_GPlusAll                   500000              5923 ns/op             698 B/op         16 allocs/op
BenchmarkGocraftWeb_GPlusAll              100000             29503 ns/op            8354 B/op        117 allocs/op
BenchmarkGoji_GPlusAll                    100000             17107 ns/op            3775 B/op         22 allocs/op
BenchmarkGoJsonRest_GPlusAll               10000            130358 ns/op           24121 B/op        402 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            177879 ns/op            9732 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll             1000000              3864 ns/op             660 B/op         11 allocs/op
BenchmarkHttpTreeMux_GPlusAll             200000             11074 ns/op            3775 B/op         22 allocs/op
BenchmarkKocha_GPlusAll                   200000              8916 ns/op            1015 B/op         43 allocs/op
BenchmarkMartini_GPlusAll                  10000            200905 ns/op           15552 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      50000             67146 ns/op           17709 B/op        346 allocs/op
BenchmarkRevel_GPlusAll                   100000             27046 ns/op            7930 B/op        107 allocs/op
BenchmarkTigerTonic_GPlusAll               50000             69351 ns/op           15492 B/op        322 allocs/op
BenchmarkTraffic_GPlusAll                  10000            213909 ns/op           42170 B/op        447 allocs/op
```


## Conclusions
First of all, there is no reason to use net/http's default [ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is very limited and does not have especially good performance. There are enough alternatives coming in every flavor, choose the one you like best.

Secondly, the broad range of functions of some of the frameworks comes at a high price in terms of performance. For example Martini has great flexibility, but very bad performance. Martini has the worst performance of all tested routers in a lot of the benchmarks. Beego seems to have some scalability problems and easily defeats Martini with even worse performance, when the number of parameters or routes is high. I really hope, that the routing of these packages can be optimized. I think the Go-ecosystem needs great feature-rich frameworks like these.

Last but not least, we have to determine the performance champion.

Denco and its predecessor Kocha-urlrouter seem to have great performance, but are not convenient to use as a router for the net/http package. A lot of extra work is necessary to use it as a http.Handler. [The README of Denco claims](https://github.com/naoina/denco/blob/b03dbb499269a597afd0db715d408ebba1329d04/README.md), that the package is not intended as a replacement for [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).

[Goji](https://github.com/zenazn/goji/) looks very decent. It has great performance while also having a great range of features, more than any other router / framework in the top group.

Currently no router can beat the performance of the [HttpRouter](https://github.com/julienschmidt/httprouter) package, which currently dominates nearly all benchmarks.

In the end, performance can not be the (only) criterion for choosing a router. Play around a bit with some of the routers, and choose the one you like best.
