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
| HttpServeMux |__18064 B__|         -  |        -  |        -  |
| Beego        |  79472 B  |  497248 B  |  26480 B  |  38768 B  |
| Denco        |  44752 B  |  107632 B  |  54896 B  |  36368 B  |
| Gocraft Web  |  57976 B  |   95736 B  |   8024 B  |  13120 B  |
| Goji         |  32400 B  | __58424 B__| __3392 B__| __6704 B__|
| Go-Json-Rest | 152608 B  |  148352 B  |  11696 B  |  13712 B  |
| Gorilla Mux  | 685152 B  | 1557216 B  |  80240 B  | 125480 B  |
| HttpRouter   |__26232 B__| __44344 B__| __3144 B__| __5792 B__|
| HttpTreeMux  |  75624 B  |   81408 B  |   7712 B  |   7616 B  |
| Kocha        | 130336 B  |  811744 B  | 139968 B  | 191632 B  |
| Martini      | 312592 B  |  579472 B  |  27520 B  |  50608 B  |
| Pat          |__21272 B__| __18968 B__| __1448 B__| __2360 B__|
| Revel        |  65320 B  |  104528 B  |   9264 B  |  14688 B  |
| TigerTonic   |  85264 B  |   99392 B  |  10576 B  |  11008 B  |
| Traffic      | 649568 B  | 1124704 B  |  57984 B  |  98168 B  |

The first place goes to [Pat](https://github.com/bmizerany/pat), followed by [HttpRouter](https://github.com/julienschmidt/httprouter) and [Goji](https://github.com/zenazn/goji/). Now, before everyone starts reading the documentation of Pat, `[SPOILER]` this low memory consumption comes at the price of relatively bad routing performance. The routing structure of Pat is simple - probably too simple. `[/SPOILER]`.

Moreover main memory is cheap and usually not a scarce resource. As long as the router doesn't require Megabytes of memory, it should be no deal breaker. But it gives us a first hint how efficient or wasteful a router works.


### Static Routes

The `Static` benchmark is not really a clone of a real-world API. It is just a collection of random static paths inspired by the structure of the Go directory. It might not be a realistic URL-structure.

The only intention of this benchmark is to allow a comparison with the default router of Go's net/http package, [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is limited to static routes and does not support parameters in the route pattern.

In the `StaticAll` benchmark each of 157 URLs is called once per repetition (op, *operation*). If you are unfamiliar with the `go test -bench` tool, the first number is the number of repetitions the `go test` tool made, to get a test running long enough for measurements. The second column shows the time in nanoseconds that a single repetition takes. The third number is the amount of heap memory allocated in bytes, the last one the average number of allocations made per repetition.

The logs below show, that http.ServeMux has only medium performance, compared to more feature-rich routers. The fastest router only needs 1.8% of the time http.ServeMux needs.

[HttpRouter](https://github.com/julienschmidt/httprouter) was the first router (I know of) that managed to serve all the static URLs without a single heap allocation. Since [the first run of this benchmark](https://github.com/julienschmidt/go-http-routing-benchmark/blob/0eb78904be13aee7a1e9f8943386f7c26b9d9d79/README.md) more routers followed this trend and were optimized in the same way.

```
BenchmarkHttpServeMux_StaticAll     1000           1509664 ns/op             104 B/op          8 allocs/op

BenchmarkBeego_StaticAll            1000           2373893 ns/op          521254 B/op      15181 allocs/op
BenchmarkDenco_StaticAll          100000             19112 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_StaticAll      10000            246574 ns/op           49249 B/op        951 allocs/op
BenchmarkGoji_StaticAll            10000            128310 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_StaticAll       2000           1513146 ns/op          183219 B/op       4130 allocs/op
BenchmarkGorillaMux_StaticAll        500           6787297 ns/op           72408 B/op        966 allocs/op
BenchmarkHttpRouter_StaticAll      50000             31463 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_StaticAll     50000             31474 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_StaticAll           50000             33299 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_StaticAll           500           5169444 ns/op          145810 B/op       2522 allocs/op
BenchmarkPat_StaticAll              1000           2519000 ns/op          554338 B/op      11250 allocs/op
BenchmarkRevel_StaticAll           10000            171800 ns/op           30753 B/op        633 allocs/op
BenchmarkTigerTonic_StaticAll      20000             93011 ns/op            7782 B/op        158 allocs/op
BenchmarkTraffic_StaticAll           100          16320267 ns/op         3798960 B/op      27960 allocs/op
```

### Micro Benchmarks

The following benchmarks measure the cost of some very basic operations.

In the first benchmark, only a single route, containing a parameter, is loaded into the routers. Then a request for a URL matching this pattern is made and the router has to call the respective registered handler function. End.
```
BenchmarkBeego_Param              200000              9306 ns/op            1213 B/op         21 allocs/op
BenchmarkDenco_Param             5000000               448 ns/op              50 B/op          2 allocs/op
BenchmarkGocraftWeb_Param        1000000              2350 ns/op             674 B/op          9 allocs/op
BenchmarkGoji_Param              1000000              1231 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Param         200000             10754 ns/op            1804 B/op         30 allocs/op
BenchmarkGorillaMux_Param         500000              7824 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_Param       10000000               221 ns/op              33 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param       2000000               905 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param             5000000               541 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_Param            200000              9264 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_Param               1000000              2711 ns/op             687 B/op         14 allocs/op
BenchmarkRevel_Param             1000000              2056 ns/op             639 B/op          8 allocs/op
BenchmarkTigerTonic_Param         500000              4484 ns/op            1027 B/op         19 allocs/op
BenchmarkTraffic_Param            200000              8754 ns/op            2030 B/op         23 allocs/op
```

Same as before, but now with multiple parameters, all in the same single route. The intention is to see how the routers scale with the number of parameters. The values of the parameters must be passed to the handler function somehow, which requires allocations. Let's see how clever the routers solve this task with a route containing 5 and 20 parameters:
```
BenchmarkBeego_Param5                  50000             31199 ns/op            1343 B/op         21 allocs/op
BenchmarkDenco_Param5                1000000              1676 ns/op             409 B/op          5 allocs/op
BenchmarkGocraftWeb_Param5            500000              3676 ns/op             948 B/op         12 allocs/op
BenchmarkGoji_Param5                 1000000              1842 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Param5            100000             16482 ns/op            3291 B/op         41 allocs/op
BenchmarkGorillaMux_Param5            200000             13985 ns/op             916 B/op          7 allocs/op
BenchmarkHttpRouter_Param5           5000000               571 ns/op             163 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param5          1000000              1432 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Param5                1000000              2178 ns/op             449 B/op         10 allocs/op
BenchmarkMartini_Param5               100000             19271 ns/op            1317 B/op         13 allocs/op
BenchmarkPat_Param5                   500000              5760 ns/op            1494 B/op         25 allocs/op
BenchmarkRevel_Param5                 500000              3757 ns/op             981 B/op         15 allocs/op
BenchmarkTigerTonic_Param5            200000             14786 ns/op            2640 B/op         53 allocs/op
BenchmarkTraffic_Param5               200000             14909 ns/op            2358 B/op         31 allocs/op

BenchmarkBeego_Param20                 10000            173208 ns/op            3739 B/op         24 allocs/op
BenchmarkDenco_Param20                500000              4884 ns/op            1679 B/op          7 allocs/op
BenchmarkGocraftWeb_Param20           200000             11571 ns/op            3864 B/op         17 allocs/op
BenchmarkGoji_Param20                 500000              5333 ns/op            1260 B/op          2 allocs/op
BenchmarkGoJsonRest_Param20            50000             42345 ns/op           10672 B/op         77 allocs/op
BenchmarkGorillaMux_Param20           100000             24837 ns/op            3313 B/op         10 allocs/op
BenchmarkHttpRouter_Param20          1000000              1716 ns/op             653 B/op          1 allocs/op
BenchmarkHttpTreeMux_Param20          200000              8266 ns/op            2219 B/op          4 allocs/op
BenchmarkKocha_Param20                500000              7057 ns/op            1839 B/op         27 allocs/op
BenchmarkMartini_Param20               20000             86362 ns/op            3715 B/op         16 allocs/op
BenchmarkPat_Param20                  500000              5775 ns/op            1494 B/op         25 allocs/op
BenchmarkRevel_Param20                200000             13848 ns/op            4565 B/op         35 allocs/op
BenchmarkTigerTonic_Param20            50000             56685 ns/op           11268 B/op        179 allocs/op
BenchmarkTraffic_Param20               50000             40647 ns/op            8250 B/op         68 allocs/op
```

Now let's see how expensive it is to access a parameter. The handler function reads the value (by the name of the parameter, e.g. with a map lookup; depends on the router) and writes it to our [web scale storage](https://www.youtube.com/watch?v=b2F-DItXtZs) (`/dev/null`).
```
BenchmarkBeego_ParamWrite         200000             10809 ns/op            1654 B/op         26 allocs/op
BenchmarkDenco_ParamWrite        5000000               547 ns/op              50 B/op          2 allocs/op
BenchmarkGocraftWeb_ParamWrite   1000000              2545 ns/op             683 B/op         10 allocs/op
BenchmarkGoji_ParamWrite         1000000              1315 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_ParamWrite    200000             12879 ns/op            2286 B/op         35 allocs/op
BenchmarkGorillaMux_ParamWrite    500000              7451 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParamWrite  10000000               298 ns/op              33 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParamWrite  1000000              1007 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParamWrite        5000000               640 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_ParamWrite       200000             10829 ns/op            1287 B/op         16 allocs/op
BenchmarkPat_ParamWrite           500000              4586 ns/op            1129 B/op         19 allocs/op
BenchmarkRevel_ParamWrite        1000000              2178 ns/op             639 B/op          8 allocs/op
BenchmarkTigerTonic_ParamWrite    500000              7212 ns/op            1485 B/op         25 allocs/op
BenchmarkTraffic_ParamWrite       200000             10614 ns/op            2464 B/op         27 allocs/op
```

### [Parse.com](https://parse.com/docs/rest#summary)

Enough of the micro benchmark stuff. Let's play a bit with real APIs. In the first set of benchmarks, we use a clone of the structure of [Parse](https://parse.com)'s decent medium-sized REST API, consisting of 26 routes.

The tasks are 1.) routing a static URL (no parameters), 2.) routing a URL containing 1 parameter, 3.) same with 2 parameters, 4.) route all of the routes once (like the StaticAll benchmark, but the routes now contain parameters).

Worth noting is, that the requested route might be a good case for some routing algorithms, while it is a bad case for another algorithm. The values might vary slightly depending on the selected route.

```
BenchmarkBeego_ParseStatic                500000              5902 ns/op            1295 B/op         22 allocs/op
BenchmarkDenco_ParseStatic              50000000              60.5 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_ParseStatic          1000000              1449 ns/op             314 B/op          6 allocs/op
BenchmarkGoji_ParseStatic                5000000               537 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_ParseStatic           200000              8858 ns/op            1147 B/op         26 allocs/op
BenchmarkGorillaMux_ParseStatic           200000             10315 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_ParseStatic         50000000              61.0 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_ParseStatic        20000000               101 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_ParseStatic              20000000              97.2 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_ParseStatic              500000              8215 ns/op             862 B/op         12 allocs/op
BenchmarkPat_ParseStatic                 1000000              1255 ns/op             249 B/op          5 allocs/op
BenchmarkRevel_ParseStatic               2000000               937 ns/op             180 B/op          4 allocs/op
BenchmarkTigerTonic_ParseStatic          5000000               421 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_ParseStatic              200000             11173 ns/op            2396 B/op         25 allocs/op

BenchmarkBeego_ParseParam                 200000             13410 ns/op            1827 B/op         35 allocs/op
BenchmarkDenco_ParseParam                5000000               493 ns/op              50 B/op          2 allocs/op
BenchmarkGocraftWeb_ParseParam           1000000              2515 ns/op             690 B/op          9 allocs/op
BenchmarkGoji_ParseParam                 1000000              1589 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_ParseParam            200000             11119 ns/op            1809 B/op         30 allocs/op
BenchmarkGorillaMux_ParseParam            200000             11774 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_ParseParam          10000000               288 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_ParseParam          2000000               981 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_ParseParam                5000000               582 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_ParseParam               200000             11952 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_ParseParam                   500000              4430 ns/op            1197 B/op         20 allocs/op
BenchmarkRevel_ParseParam                1000000              2221 ns/op             655 B/op          8 allocs/op
BenchmarkTigerTonic_ParseParam            500000              4841 ns/op            1084 B/op         19 allocs/op
BenchmarkTraffic_ParseParam               200000             11207 ns/op            2330 B/op         25 allocs/op

BenchmarkBeego_Parse2Params               100000             15447 ns/op            1989 B/op         35 allocs/op
BenchmarkDenco_Parse2Params              2000000               822 ns/op             116 B/op          3 allocs/op
BenchmarkGocraftWeb_Parse2Params         1000000              2858 ns/op             736 B/op         10 allocs/op
BenchmarkGoji_Parse2Params               1000000              1576 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_Parse2Params          200000             12735 ns/op            2167 B/op         33 allocs/op
BenchmarkGorillaMux_Parse2Params          200000             12513 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_Parse2Params         5000000               333 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_Parse2Params        1000000              1144 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_Parse2Params              1000000              1019 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_Parse2Params             200000             12469 ns/op            1220 B/op         13 allocs/op
BenchmarkPat_Parse2Params                 500000              4520 ns/op             908 B/op         21 allocs/op
BenchmarkRevel_Parse2Params              1000000              2666 ns/op             722 B/op         10 allocs/op
BenchmarkTigerTonic_Parse2Params          500000              7759 ns/op            1488 B/op         28 allocs/op
BenchmarkTraffic_Parse2Params             200000             11562 ns/op            2135 B/op         25 allocs/op

BenchmarkBeego_ParseAll                    10000            324418 ns/op           40499 B/op        777 allocs/op
BenchmarkDenco_ParseAll                   200000             11777 ns/op            1008 B/op         35 allocs/op
BenchmarkGocraftWeb_ParseAll               50000             59215 ns/op           14319 B/op        210 allocs/op
BenchmarkGoji_ParseAll                     50000             32191 ns/op            5491 B/op         33 allocs/op
BenchmarkGoJsonRest_ParseAll               10000            279400 ns/op           41644 B/op        759 allocs/op
BenchmarkGorillaMux_ParseAll                5000            457558 ns/op           17274 B/op        175 allocs/op
BenchmarkHttpRouter_ParseAll              500000              6466 ns/op             665 B/op         16 allocs/op
BenchmarkHttpTreeMux_ParseAll             100000             19673 ns/op            5491 B/op         33 allocs/op
BenchmarkKocha_ParseAll                   200000             14527 ns/op            1160 B/op         54 allocs/op
BenchmarkMartini_ParseAll                  10000            308871 ns/op           27715 B/op        333 allocs/op
BenchmarkPat_ParseAll                      20000             85603 ns/op           18296 B/op        385 allocs/op
BenchmarkRevel_ParseAll                    50000             51380 ns/op           12646 B/op        176 allocs/op
BenchmarkTigerTonic_ParseAll               10000            108252 ns/op           20872 B/op        420 allocs/op
BenchmarkTraffic_ParseAll                  10000            359819 ns/op           70704 B/op        763 allocs/op
```


### [GitHub](http://developer.github.com/v3/)

The GitHub API is rather large, consisting of 203 routes. The tasks are basically the same as in the benchmarks before.

```
BenchmarkBeego_GithubStatic               500000              6292 ns/op            1197 B/op         38 allocs/op
BenchmarkDenco_GithubStatic             20000000              84.7 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GithubStatic         1000000              1457 ns/op             314 B/op          6 allocs/op
BenchmarkGoji_GithubStatic               5000000               604 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_GithubStatic          200000              9061 ns/op            1163 B/op         26 allocs/op
BenchmarkGorillaMux_GithubStatic           50000             58162 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GithubStatic        20000000               125 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GithubStatic       20000000               105 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GithubStatic             20000000               132 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GithubStatic             100000             26241 ns/op             862 B/op         12 allocs/op
BenchmarkPat_GithubStatic                 100000             16814 ns/op            3789 B/op         76 allocs/op
BenchmarkRevel_GithubStatic              2000000               962 ns/op             180 B/op          4 allocs/op
BenchmarkTigerTonic_GithubStatic         5000000               494 ns/op              49 B/op          1 allocs/op
BenchmarkTraffic_GithubStatic              20000             95631 ns/op           23369 B/op        172 allocs/op

BenchmarkBeego_GithubParam                 20000             77679 ns/op            3038 B/op         57 allocs/op
BenchmarkDenco_GithubParam               2000000               967 ns/op             116 B/op          3 allocs/op
BenchmarkGocraftWeb_GithubParam          1000000              2989 ns/op             737 B/op         10 allocs/op
BenchmarkGoji_GithubParam                1000000              2277 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GithubParam           200000             13102 ns/op            2194 B/op         33 allocs/op
BenchmarkGorillaMux_GithubParam            50000             40468 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_GithubParam          5000000               568 ns/op              98 B/op          1 allocs/op
BenchmarkHttpTreeMux_GithubParam         1000000              1318 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GithubParam               1000000              1156 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_GithubParam               50000             34692 ns/op            1221 B/op         13 allocs/op
BenchmarkPat_GithubParam                  200000             12177 ns/op            2628 B/op         56 allocs/op
BenchmarkRevel_GithubParam               1000000              2826 ns/op             739 B/op         10 allocs/op
BenchmarkTigerTonic_GithubParam           500000              7700 ns/op            1484 B/op         28 allocs/op
BenchmarkTraffic_GithubParam               50000             44948 ns/op            7153 B/op         60 allocs/op

BenchmarkBeego_GithubAll                      50          40861674 ns/op          519988 B/op      11210 allocs/op
BenchmarkDenco_GithubAll                   10000            174017 ns/op           21363 B/op        508 allocs/op
BenchmarkGocraftWeb_GithubAll               5000            585437 ns/op          136615 B/op       1915 allocs/op
BenchmarkGoji_GithubAll                     2000           1038040 ns/op           57349 B/op        347 allocs/op
BenchmarkGoJsonRest_GithubAll               1000           2555637 ns/op          408055 B/op       6558 allocs/op
BenchmarkGorillaMux_GithubAll                100          25697013 ns/op          153510 B/op       1420 allocs/op
BenchmarkHttpRouter_GithubAll              20000             96329 ns/op           14102 B/op        169 allocs/op
BenchmarkHttpTreeMux_GithubAll             10000            225263 ns/op           57350 B/op        347 allocs/op
BenchmarkKocha_GithubAll                   10000            222596 ns/op           24072 B/op        847 allocs/op
BenchmarkMartini_GithubAll                   100          15716582 ns/op          245764 B/op       2943 allocs/op
BenchmarkPat_GithubAll                       500           6950182 ns/op         1589281 B/op      32576 allocs/op
BenchmarkRevel_GithubAll                    5000            556683 ns/op          131270 B/op       1847 allocs/op
BenchmarkTigerTonic_GithubAll               2000           1487732 ns/op          251075 B/op       6087 allocs/op
BenchmarkTraffic_GithubAll                   100          21345637 ns/op         3176442 B/op      24961 allocs/op
```

### [Google+](https://developers.google.com/+/api/latest/)

Last but not least the Google+ API, consisting of 13 routes. In reality this is just a subset of a much larger API.

```
BenchmarkBeego_GPlusStatic               1000000              4016 ns/op             853 B/op         18 allocs/op
BenchmarkDenco_GPlusStatic              50000000              54.6 ns/op               0 B/op          0 allocs/op
BenchmarkGocraftWeb_GPlusStatic          1000000              1354 ns/op             297 B/op          6 allocs/op
BenchmarkGoji_GPlusStatic                5000000               427 ns/op               0 B/op          0 allocs/op
BenchmarkGoJsonRest_GPlusStatic           200000              8791 ns/op            1147 B/op         26 allocs/op
BenchmarkGorillaMux_GPlusStatic           500000              4902 ns/op             460 B/op          6 allocs/op
BenchmarkHttpRouter_GPlusStatic         50000000              59.1 ns/op               0 B/op          0 allocs/op
BenchmarkHttpTreeMux_GPlusStatic        50000000              56.3 ns/op               0 B/op          0 allocs/op
BenchmarkKocha_GPlusStatic              20000000              89.0 ns/op               0 B/op          0 allocs/op
BenchmarkMartini_GPlusStatic              500000              7149 ns/op             862 B/op         12 allocs/op
BenchmarkPat_GPlusStatic                 5000000               514 ns/op              99 B/op          2 allocs/op
BenchmarkRevel_GPlusStatic               2000000               845 ns/op             164 B/op          4 allocs/op
BenchmarkTigerTonic_GPlusStatic         10000000               286 ns/op              33 B/op          1 allocs/op
BenchmarkTraffic_GPlusStatic              500000              7105 ns/op            1513 B/op         19 allocs/op

BenchmarkBeego_GPlusParam                 200000             12606 ns/op            1279 B/op         23 allocs/op
BenchmarkDenco_GPlusParam                5000000               530 ns/op              50 B/op          2 allocs/op
BenchmarkGocraftWeb_GPlusParam           1000000              2472 ns/op             674 B/op          9 allocs/op
BenchmarkGoji_GPlusParam                 1000000              1407 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlusParam            200000             12421 ns/op            1822 B/op         30 allocs/op
BenchmarkGorillaMux_GPlusParam            200000             12751 ns/op             786 B/op          7 allocs/op
BenchmarkHttpRouter_GPlusParam           5000000               349 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlusParam          1000000              1008 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlusParam                5000000               639 ns/op              58 B/op          3 allocs/op
BenchmarkMartini_GPlusParam               200000             13275 ns/op            1188 B/op         13 allocs/op
BenchmarkPat_GPlusParam                  1000000              3094 ns/op             753 B/op         14 allocs/op
BenchmarkRevel_GPlusParam                1000000              2212 ns/op             656 B/op          8 allocs/op
BenchmarkTigerTonic_GPlusParam            500000              5044 ns/op            1103 B/op         19 allocs/op
BenchmarkTraffic_GPlusParam               200000             11654 ns/op            2044 B/op         23 allocs/op

BenchmarkBeego_GPlus2Params               100000             16441 ns/op            1340 B/op         23 allocs/op
BenchmarkDenco_GPlus2Params              2000000               953 ns/op             116 B/op          3 allocs/op
BenchmarkGocraftWeb_GPlus2Params         1000000              3045 ns/op             737 B/op         10 allocs/op
BenchmarkGoji_GPlus2Params               1000000              2095 ns/op             343 B/op          2 allocs/op
BenchmarkGoJsonRest_GPlus2Params          200000             13315 ns/op            2197 B/op         33 allocs/op
BenchmarkGorillaMux_GPlus2Params           50000             34663 ns/op             819 B/op          7 allocs/op
BenchmarkHttpRouter_GPlus2Params         5000000               421 ns/op              65 B/op          1 allocs/op
BenchmarkHttpTreeMux_GPlus2Params        1000000              1246 ns/op             343 B/op          2 allocs/op
BenchmarkKocha_GPlus2Params              1000000              1163 ns/op             132 B/op          5 allocs/op
BenchmarkMartini_GPlus2Params              50000             43900 ns/op            1320 B/op         17 allocs/op
BenchmarkPat_GPlus2Params                 200000             10062 ns/op            2403 B/op         41 allocs/op
BenchmarkRevel_GPlus2Params              1000000              2911 ns/op             753 B/op         10 allocs/op
BenchmarkTigerTonic_GPlus2Params          500000              8245 ns/op            1586 B/op         28 allocs/op
BenchmarkTraffic_GPlus2Params              50000             37030 ns/op            3629 B/op         35 allocs/op

BenchmarkBeego_GPlusAll                    10000            186509 ns/op           16506 B/op        299 allocs/op
BenchmarkDenco_GPlusAll                   200000              8759 ns/op             887 B/op         27 allocs/op
BenchmarkGocraftWeb_GPlusAll               50000             33987 ns/op            8354 B/op        117 allocs/op
BenchmarkGoji_GPlusAll                    100000             20068 ns/op            3775 B/op         22 allocs/op
BenchmarkGoJsonRest_GPlusAll               10000            153430 ns/op           24122 B/op        402 allocs/op
BenchmarkGorillaMux_GPlusAll               10000            219266 ns/op            9732 B/op         91 allocs/op
BenchmarkHttpRouter_GPlusAll              500000              4644 ns/op             660 B/op         11 allocs/op
BenchmarkHttpTreeMux_GPlusAll             200000             12616 ns/op            3775 B/op         22 allocs/op
BenchmarkKocha_GPlusAll                   200000             10847 ns/op            1015 B/op         43 allocs/op
BenchmarkMartini_GPlusAll                  10000            243172 ns/op           15553 B/op        194 allocs/op
BenchmarkPat_GPlusAll                      20000             78577 ns/op           17710 B/op        346 allocs/op
BenchmarkRevel_GPlusAll                    50000             31080 ns/op            7929 B/op        107 allocs/op
BenchmarkTigerTonic_GPlusAll               20000             83006 ns/op           15492 B/op        322 allocs/op
BenchmarkTraffic_GPlusAll                  10000            255141 ns/op           42171 B/op        447 allocs/op
```


## Conclusions
First of all, there is no reason to use net/http's default [ServeMux](http://golang.org/pkg/net/http/#ServeMux), which is very limited and does not have especially good performance. There are enough alternatives coming in every flavor, choose the one you like best.

Secondly, the broad range of functions of some of the frameworks comes at a high price in terms of performance. For example Martini has great flexibility, but very bad performance. Martini has the worst performance of all tested routers in a lot of the benchmarks. Beego seems to have some scalability problems and easily defeats Martini with even worse performance, when the number of parameters or routes is high. I really hope, that the routing of these packages can be optimized. I think the Go-ecosystem needs great feature-rich frameworks like these.

Last but not least, we have to determine the performance champion.

Denco and its predecessor Kocha-urlrouter seem to have great performance, but are not convenient to use as a router for the net/http package. A lot of extra work is necessary to use it as a http.Handler. [The README of Denco claims](https://github.com/naoina/denco/blob/b03dbb499269a597afd0db715d408ebba1329d04/README.md), that the package is not intended as a replacement for [http.ServeMux](http://golang.org/pkg/net/http/#ServeMux).

[Goji](https://github.com/zenazn/goji/) looks very decent. It has great performance while also having a great range of features, more than any other router / framework in the top group.

Currently no router can beat the performance of the [HttpRouter](https://github.com/julienschmidt/httprouter) package, which currently dominates nearly all benchmarks.

In the end, performance can not be the (only) criterion for choosing a router. Play around a bit with some of the routers, and choose the one you like best.
