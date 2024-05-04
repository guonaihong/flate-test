# flate-test

测试flate性能的仓库

goos: linux
goarch: amd64
pkg: github.com/guonaihong/flate-test
cpu: AMD Ryzen 7 5800H with Radeon Graphics
Benchmark_Encode_9_1KB-16                                  41922             28382 ns/op
Benchmark_Encode_8_1KB-16                                  42033             28814 ns/op
Benchmark_Encode_7_1KB-16                                  41984             28268 ns/op
Benchmark_Encode_6_1KB-16                                  41331             29326 ns/op
Benchmark_Encode_5_1KB-16                                  42123             28455 ns/op
Benchmark_Encode_4_1KB-16                                  42778             27857 ns/op
Benchmark_Encode_3_1KB-16                                  43016             27110 ns/op
Benchmark_Encode_1_1KB-16                                  74268             16065 ns/op
Benchmark_Encode_WithPool_9_13KB-16                         4788            250575 ns/op
Benchmark_Encode_WithPool_8_13KB-16                         4724            250559 ns/op
Benchmark_Encode_WithPool_7_13KB-16                         4719            251931 ns/op
Benchmark_Encode_WithPool_6_13KB-16                         4918            243240 ns/op
Benchmark_Encode_WithPool_5_13KB-16                         4911            242923 ns/op
Benchmark_Encode_WithPool_4_13KB-16                         6398            184439 ns/op
Benchmark_Encode_WithPool_3_13KB-16                         6807            175030 ns/op
Benchmark_Encode_WithPool_1_13KB-16                        14581             84988 ns/op
Benchmark_Encode_WithPool_9_1KB-16                         42796             27724 ns/op
Benchmark_Encode_WithPool_8_1KB-16                         43936             27686 ns/op
Benchmark_Encode_WithPool_7_1KB-16                         43330             27563 ns/op
Benchmark_Encode_WithPool_6_1KB-16                         43492             27669 ns/op
Benchmark_Encode_WithPool_5_1KB-16                         42798             29274 ns/op
Benchmark_Encode_WithPool_4_1KB-16                         44089             27115 ns/op
Benchmark_Encode_WithPool_3_1KB-16                         44806             27006 ns/op
Benchmark_Encode_WithPool_1_1KB-16                         76774             15600 ns/op
Benchmark_Encode_thirdpart_9_13KB-16                        4053            282761 ns/op
Benchmark_Encode_thirdpart_8_13KB-16                        5253            219571 ns/op
Benchmark_Encode_thirdpart_7_13KB-16                        6211            190900 ns/op
Benchmark_Encode_thirdpart_6_13KB-16                       13134             91583 ns/op
Benchmark_Encode_thirdpart_5_13KB-16                       13143             91194 ns/op
Benchmark_Encode_thirdpart_4_13KB-16                       17906             67192 ns/op
Benchmark_Encode_thirdpart_3_13KB-16                       18805             63743 ns/op
Benchmark_Encode_thirdpart_1_13KB-16                       23798             49432 ns/op
Benchmark_Encode_thirdpart_9_1KB-16                        41011             28425 ns/op
Benchmark_Encode_thirdpart_8_1KB-16                        51123             23701 ns/op
Benchmark_Encode_thirdpart_7_1KB-16                        50995             23460 ns/op
Benchmark_Encode_thirdpart_6_1KB-16                        85968             13836 ns/op
Benchmark_Encode_thirdpart_5_1KB-16                        85766             13773 ns/op
Benchmark_Encode_thirdpart_4_1KB-16                        93132             12370 ns/op
Benchmark_Encode_thirdpart_3_1KB-16                        98865             11898 ns/op
Benchmark_Encode_thirdpart_1_1KB-16                       110989             10684 ns/op
Benchmark_Encode_WithPool_thirdpart_9_13KB-16               4218            276667 ns/op
Benchmark_Encode_WithPool_thirdpart_8_13KB-16               5479            214980 ns/op
Benchmark_Encode_WithPool_thirdpart_7_13KB-16               6219            185042 ns/op
Benchmark_Encode_WithPool_thirdpart_6_13KB-16              14497             81490 ns/op
Benchmark_Encode_WithPool_thirdpart_5_13KB-16              14426             82806 ns/op
Benchmark_Encode_WithPool_thirdpart_4_13KB-16              19776             60454 ns/op
Benchmark_Encode_WithPool_thirdpart_3_13KB-16              20791             58013 ns/op
Benchmark_Encode_WithPool_thirdpart_1_13KB-16              26470             45441 ns/op
Benchmark_Encode_WithPool_thirdpart_9_1KB-16               43581             27239 ns/op
Benchmark_Encode_WithPool_thirdpart_8_1KB-16               53634             23823 ns/op
Benchmark_Encode_WithPool_thirdpart_7_1KB-16               53194             22517 ns/op
Benchmark_Encode_WithPool_thirdpart_6_1KB-16               90272             13188 ns/op
Benchmark_Encode_WithPool_thirdpart_5_1KB-16               90769             13195 ns/op
Benchmark_Encode_WithPool_thirdpart_4_1KB-16              101108             11784 ns/op
Benchmark_Encode_WithPool_thirdpart_3_1KB-16              105606             11319 ns/op
Benchmark_Encode_WithPool_thirdpart_1_1KB-16              117613             10163 ns/op
PASS
ok      github.com/guonaihong/flate-test        82.094s
