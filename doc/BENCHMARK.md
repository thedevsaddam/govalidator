Benchmarks
===================

Machine: XPS 13 9370 (07E6)
Go version: go version go1.12.6 linux/amd64

| ➜ go test -run=XXX -bench=. -benchmem=true |            |              |              |              |
|--------------------------------------------|------------|--------------|--------------|--------------|
|Benchmark_IsAlpha-8                         | 10000000	  | 205 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsAlphaDash-8                     | 5000000	  | 268 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsAlphaNumeric-8                  | 10000000	  | 182 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsBoolean-8                       | 200000000  | 6.84 ns/op   | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsCreditCard-8                    | 10000000	  | 243 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsCoordinate-8             	     | 3000000	  | 482 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsCSSColor-8                      | 10000000	  | 160 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsDate-8                   	     | 3000000	  | 531 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsDateDDMMYY-8             	     | 5000000	  | 246 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsEmail-8                  	     | 3000000	  | 549 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsFloat-8                         | 10000000	  | 199 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsIn-8                            | 5000000    | 3.77 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsJSON-8                   	     | 2000000	  | 956 ns/op	 | 640 B/op	    | 12 allocs/op | 
|Benchmark_IsMacAddress-8             	     | 5000000	  | 277 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsNumeric-8                       | 20000000	  | 110 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsLatitude-8               	     | 5000000	  | 249 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsLongitude-8              	     | 5000000	  | 250 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsIP-8                     	     | 3000000	  | 578 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsIPV4-8                   	     | 5000000	  | 286 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsIPV6-8                   	     | 2000000	  | 931 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsMatchedRegex-8           	     | 200000	  | 5786 ns/op	 | 4465 B/op	| 57 allocs/op | 
|Benchmark_IsURL-8                    	     | 2000000	  | 866 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsUUID-8                   	     | 3000000	  | 455 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsUUID3-8                  	     | 3000000	  | 536 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsUUID4-8                  	     | 3000000	  | 411 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_IsUUID5-8                  	     | 3000000	  | 443 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|BenchmarkRoller_Start-8              	     | 300000	  | 4659 ns/op	 | 2468 B/op	| 28 allocs/op | 
|Benchmark_isContainRequiredField-8          | 1000000000 | 2.69 ns/op	 | 0 B/op	    | 0 allocs/op  | 
|Benchmark_Validate-8                 	     | 200000	  | 6742 ns/op	 | 727 B/op	    | 29 allocs/op | 
