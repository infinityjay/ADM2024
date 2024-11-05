
# Report for Assignment 3

## 0. Group member

The group number is 6.

| **Name** | **Student  Number** | **Tasks**         |
| -------- | ------------------- | ----------------- |
| Jie Chen | S4162315            | All code & report |

## 1. Introduction

### 1.1 Hardware environment

Here is the hardware environment of my experiment.

* Chip 

  Apple M1 Pro chip, 200GB/s memory bandwidth

* CPU

  Clock rate: 2064-3220MHz, 24MB Level 3 Cache

* Main memory

  32GB

* Disk

  512GB SSD, 4900 MB/s read speed and 3951 MB/s write speed

### 1.2 Software

* golang 

  Version: 1.20

Other dependencies versions is listed in go.mod

### 1.2 Main task



## 2. Compression techniques

### 2.1 Uncompressed binary format(bin)

2.1.1 implementation

2.1.2 results & analysis

### 2.6 bit packing

write function for each packInt8toInt32 ... and run the unit test.(uint to encode)

list all the function location, and application



## Conclusion



*Compression ratio = uncompressed file size / compressed file size

| Compression tech | Input file name           | Input file size(MB) | Output file size(MB) | Compression ratio* | Encode  time(ms) | Decode time(ms) |
| ---------------- | ------------------------- | ------------------- | -------------------- | ------------------ | ---------------- | --------------- |
| bin              | l_discount-int8.csv       | 11.97               | 7.87                 | 1.52               | 581              | 276             |
| bin              | l_discount-int16.csv      | 11.97               | 7.87                 | 1.52               | 617              | 256             |
| bin              | l_discount-int32.csv      | 11.97               | 7.87                 | 1.52               | 600              | 274             |
| bin              | l_discount-int64.csv      | 11.97               | 7.87                 | 1.52               | 597              | 273             |
| bin              | l_orderkey-int32.csv      | 44.73               | too large...         |                    |                  |                 |
| bin              | l_partkey-int64.csv       | 36.88               | 143081.62            | 0.00               | 141873           | Super long..    |
| bin              | l_tax-int8.csv            | 11.45               | 6.44                 | 1.78               | 505              | 248             |
| rle              | l_comment-string.csv      | 157.35              | 302.39               | 0.52               | 16079            | 12082           |
| rle              | l_commitdate-string.csv   | 62.96               | 107.13               | 0.59               | 11369            | 10010           |
| rle              | l_returnflag-string.csv   | 11.45               | 17.17                | 0.67               | 9496             | 9333            |
| rle              | l_discount-int8.csv       | 11.97               | 21.28                | 0.56               | 640              | 702             |
| rle              | l_discount-int16.csv      | 11.97               | 21.28                | 0.56               | 663              | 746             |
| rle              | l_orderkey-int32.csv      | 44.73               | 14.04                | 3.19               | 787              | 305             |
| rle              | l_partkey-int64.csv       | 36.88               | 48.33                | 0.76               | 886              | 919             |
| rle              | l_tax-int8.csv            | 11.45               | 20.35                | 0.56               | 647              | 718             |
| dic              | l_comment-string.csv      | 157.35              | 204.39               | 0.77               | 4268             | 15544           |
| dic              | l_commitdate-string.csv   | 62.96               | 26.05                | 2.42               | 1611             | 10248           |
| dic              | l_returnflag-string.csv   | 11.45               | 11.45                | 1.00               | 985              | 9738            |
| dic              | l_shipinstruct-string.csv | 74.39               | 11.45                | 6.50               | 1393             | 11590           |
| dic              | l_shipmode-string.csv     | 30.25               | 11.45                | 2.64               | 1349             | 9773            |
| dic              | l_discount-int8.csv       | 11.97               | 12.49                | 0.96               | 1073             | 10095           |
| dic              | l_extendedprice-int32.csv | 45.05               | 52.34                | 0.86               | 3576             | 15216           |
| for              | l_linenumber-int8.csv     | 11.45               | 2.86                 | 4.00               | 740              | 161             |
| for              | l_discount-int8.csv       | 11.97               | 3.77                 | 3.18               | 943              | 157             |
| for              | l_tax-int8.csv            | 11.45               | 3.96                 | 2.89               | 779              | 174             |
| for              | l_quantity-int8.csv       | 16.14               | 9.33                 | 1.73               | 893              | 257             |
| for              | l_discount-int16.csv      | 11.97               | 5.06                 | 2.36               | 947              | 544             |
| for              | l_linenumber-int16.csv    | 11.45               | 2.86                 | 4.00               | 625              | 142             |
| for              | l_quantity-int16.csv      | 16.14               | 18.45                | 0.87               | 746              | 247             |
| for              | l_suppkey-int16.csv       | 27.98               | 22.88                | 1.22               | 1196             | 355             |
| for              | l_tax-int16.csv           | 11.45               | 5.55                 | 2.06               | 955              | 165             |
| for              | l_discount-int32.csv      | 11.97               | 15.06                | 0.79               | 623              | 205             |
| for              | l_extendedprice-int32.csv | 45.05               | 45.58                | 0.99               | 1208             | 517             |
| for              | l_linenumber-int32.csv    | 11.45               | 11.45                | 1.00               | 583              | 181             |
| for              | l_orderkey-int32.csv      | 44.73               | 45.60                | 0.98               | 1360             | 531             |
| for              | l_partkey-int32.csv       | 36.88               | 36.43                | 1.01               | 1080             | 493             |
| for              | l_quantity-int32.csv      | 16.14               | 12.25                | 1.32               | 1149             | 216             |
| for              | l_suppkey-int32.csv       | 27.98               | 11.45                | 2.44               | 931              | 311             |
| for              | l_discount-int64.csv      | 11.97               | 20.26                | 0.59               | 670              | 244             |
| for              | l_extendedprice-int64.csv | 45.05               | 91.15                | 0.49               | 964              | 814             |
| for              | l_linenumber-int64.csv    | 11.45               | 11.45                | 1.00               | 619              | 184             |
| for              | l_orderkey-int64.csv      | 44.73               | 91.13                | 0.49               | 1403             | 802             |
| for              | l_partkey-int64.csv       | 36.88               | 71.75                | 0.51               | 1111             | 715             |
| for              | l_quantity-int64.csv      | 16.14               | 13.40                | 1.20               | 724              | 233             |
| for              | l_suppkey-int64.csv       | 27.98               | 11.46                | 2.44               | 1241             | 319             |
| dif              | l_discount-int8.csv       | 11.97               | 4.41                 | 2.71               | 637              | 145             |
| dif              | l_linenumber-int8.csv     | 11.45               | 3.23                 | 3.55               | 708              | 127             |
| dif              | l_quantity-int8.csv       | 16.14               | 9.53                 | 1.69               | 730              | 198             |
| dif              | l_tax-int8.csv            | 11.45               | 3.96                 | 2.89               | 586              | 143             |
| dif              | l_discount-int16.csv      | 11.97               | 6.63                 | 1.80               | 641              | 176             |
| dif              | l_linenumber-int16.csv    | 11.45               | 3.75                 | 3.05               | 648              | 148             |
| dif              | l_quantity-int16.csv      | 16.14               | 18.87                | 0.86               | 820              | 258             |
| dif              | l_suppkey-int16.csv       | 27.98               | 22.88                | 1.22               | 881              | 375             |
| dif              | l_tax-int16.csv           | 11.45               | 5.55                 | 2.06               | 619              | 161             |
| dif              |                           |                     |                      |                    |                  |                 |
| dif              |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |



## 4. Usage of the code



