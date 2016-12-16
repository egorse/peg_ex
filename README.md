# peg_ex
The benchmarks of pointlander/peg with AST and without AST

It has peg generated parser, parser with initial tokens slice size 128 and parser without peg own AST (so far manually created)

The benchcmp looks like:
```
benchmark           old ns/op     new ns/op     delta
Benchmark100-4      46888         3247          -93.07%
Benchmark200-4      52011         5785          -88.88%
Benchmark500-4      62611         13024         -79.20%
Benchmark1k-4       94656         25884         -72.65%
Benchmark10k-4      374852        247468        -33.98%
Benchmark20k-4      373827        246092        -34.17%
Benchmark50k-4      1774499       1199769       -32.39%
Benchmark100k-4     3647900       2436161       -33.22%
Benchmark200k-4     7965721       4859911       -38.99%
Benchmark500k-4     21448772      13444311      -37.32%
Benchmark1M-4       41923699      28403515      -32.25%
Benchmark2M-4       81485143      55279573      -32.16%
Benchmark4M-4       155086360     109257278     -29.55%

benchmark           old allocs     new allocs     delta
Benchmark100-4      61             41             -32.79%
Benchmark200-4      104            66             -36.54%
Benchmark500-4      231            139            -39.83%
Benchmark1k-4       456            268            -41.23%
Benchmark10k-4      4254           2440           -42.64%
Benchmark20k-4      4254           2440           -42.64%
Benchmark50k-4      21115          12076          -42.81%
Benchmark100k-4     42202          24126          -42.83%
Benchmark200k-4     84374          48225          -42.84%
Benchmark500k-4     210870         120509         -42.85%
Benchmark1M-4       431823         246769         -42.85%
Benchmark2M-4       863587         493492         -42.86%
Benchmark4M-4       1727125        986943         -42.86%

benchmark           old bytes     new bytes     delta
Benchmark100-4      394752        1328          -99.66%
Benchmark200-4      395680        2144          -99.46%
Benchmark500-4      397920        4016          -98.99%
Benchmark1k-4       402528        7984          -98.02%
Benchmark10k-4      495136        89760         -81.87%
Benchmark20k-4      495136        89760         -81.87%
Benchmark50k-4      1669664       429696        -74.26%
Benchmark100k-4     3715744       842676        -77.32%
Benchmark200k-4     7938976       1799696       -77.33%
Benchmark500k-4     30062176      4687143       -84.41%
Benchmark1M-4       61947872      10775740      -82.61%
Benchmark2M-4       124239328     21501950      -82.69%
Benchmark4M-4       248584737     42716822      -82.82%
```
