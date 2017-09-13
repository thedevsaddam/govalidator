Benchmarks
===================

Machine: Mac Book Pro-2015 2.7GHz 8GB
Go version: go1.8.1 darwin/amd64

| ➜ go test -run=XXX -bench=. -benchmem=true |           |            |           |              |
|--------------------------------------------|-----------|------------|-----------|--------------|
| Benchmark_isContainRequiredField-4         | 200000000 | 7.22 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsAlpha-4                        | 5000000   | 374 ns/op  | 16 B/op   | 1 allocs/op  |
| Benchmark_IsAlphaDash-4                    | 3000000   | 458 ns/op  | 16 B/op   | 1 allocs/op  |
| Benchmark_IsAlphaNumeric-4                 | 5000000   | 355 ns/op  | 16 B/op   | 1 allocs/op  |
| Benchmark_IsBoolean-4                      | 200000000 | 9.65 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsCreditCard-4                   | 3000000   | 459 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsCoordinate-4                   | 2000000   | 870 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsCSSColor-4                     | 5000000   | 302 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsDate-4                         | 2000000   | 679 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsDateDDMMYY-4                   | 3000000   | 473 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsEmail-4                        | 1000000   | 1065 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsFloat-4                        | 3000000   | 401 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIn-4                           | 200000000 | 7.39 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsJSON-4                         | 1000000   | 1635 ns/op | 768 B/op  | 12 allocs/op |
| Benchmark_IsNumeric-4                      | 10000000  | 230 ns/op  | 16 B/op   | 1 allocs/op  |
| Benchmark_IsLatitude-4                     | 3000000   | 513 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsLongitude-4                    | 3000000   | 493 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIP-4                           | 1000000   | 1051 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIPV4-4                         | 3000000   | 515 ns/op  | 0 B/op    | 0 allocs/op  |
| Benchmark_IsIPV6-4                         | 1000000   | 1389 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsMatchedRegex-4                 | 200000    | 9488 ns/op | 6800 B/op | 74 allocs/op |
| Benchmark_IsURL-4                          | 1000000   | 1117 ns/op | 0 B/op    | 0 allocs/op  |
| Benchmark_IsUUID-4                         | 2000000   | 884 ns/op  | 16 B/op   | 1 allocs/op  |
| Benchmark_IsUUID3-4                        | 2000000   | 909 ns/op  | 16 B/op   | 1 allocs/op  |
| Benchmark_IsUUID4-4                        | 1000000   | 1090 ns/op | 16 B/op   | 1 allocs/op  |
| Benchmark_IsUUID5-4                        | 2000000   | 954 ns/op  | 16 B/op   | 1 allocs/op  |
| Benchmark_ValidateMapJSON-4                | 2000000   | 1007 ns/op | 1376 B/op | 5 allocs/op  |
| Benchmark_ValidateStructJSON-4             | 1000000   | 1140 ns/op | 1376 B/op | 5 allocs/op  |
| Benchmark_Validate-4                       | 200000    | 7606 ns/op | 1120 B/op | 31 allocs/op |
