Benchmarks
===================

Machine: Mac Book Pro-2015 2.7GHz 8GB
Go version: go1.8.1 darwin/amd64

| ➜ go test -run=XXX -bench=. -benchmem=true |           |            |           |              |
|--------------------------------------------|-----------|------------|-----------|--------------|
| Benchmark_IsAlpha-4                        | 5000000   | 323 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsAlphaDash-4                    | 3000000   | 415 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsAlphaNumeric-4                 | 5000000   | 338 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsBoolean-4                      | 100000000 | 10.6 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsCreditCard-4                   | 3000000   | 543 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsCoordinate-4                   | 2000000   | 950 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsCSSColor-4                     | 5000000   | 300 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsDate-4                         | 2000000   | 719 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsDateDDMMYY-4                   | 3000000   | 481 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsEmail-4                        | 1000000   | 1172 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsFloat-4                        | 3000000   | 432 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIn-4                           | 200000000 | 7.34 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsJSON-4                         | 1000000   | 1595 ns/op | 768 B/op  | 12 allocs/op |
| Benchmark_IsNumeric-4                      | 10000000  | 195 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsLatitude-4                     | 3000000   | 523 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsLongitude-4                    | 3000000   | 516 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIP-4                           | 1000000   | 1073 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIPV4-4                         | 3000000   | 580 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIPV6-4                         | 1000000   | 1288 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsMatchedRegex-4                 | 200000    | 7133 ns/op | 5400 B/op | 66 allocs/op |
| Benchmark_IsURL-4                          | 1000000   | 1159 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsUUID-4                         | 2000000   | 832 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsUUID3-4                        | 2000000   | 783 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsUUID4-4                        | 2000000   | 899 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsUUID5-4                        | 2000000   | 828 ns/op  | 0 B/op    | 0 allocs/op  |
| BenchmarkRoller_Start-4                    | 200000    | 6869 ns/op | 2467 B/op | 28 allocs/op |
| Benchmark_isContainRequiredField-4         | 300000000 | 4.23 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_Validate-4                       | 200000    | 9347 ns/op | 664 B/op  | 28 allocs/op |
