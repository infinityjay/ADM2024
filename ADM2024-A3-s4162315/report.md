
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

### 1.2 Software environment

* golang 

  Version: 1.20

Other dependencies versions is listed in go.mod

### 1.3 structure of file



## 2.Usage Instruction





## 2. Compression techniques

### 2.1 Uncompressed binary format(bin)

2.1.1 implementation

2.1.2 results & analysis

### 2.6 bit packing

write function for each packInt8toInt32 ... and run the unit test.(uint to encode)

list all the function location, and application



## 3. Results

### 3.1 Results of the execution



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
|                  |                           |                     |                      |                    |                  |                 |
| rle              | l_comment-string.csv      | 157.35              | 168.80               | 0.93               | 13268            | 12141           |
| rle              | l_commitdate-string.csv   | 62.96               | 73.48                | 0.86               | 13333            | 10214           |
| rle              | l_linestatus-string.csv   | 11.45               | 3.35                 | 3.42               | 2125             | 9489            |
| rle              | l_receiptdate-string.csv  | 62.96               | 73.97                | 0.85               | 13083            | 9902            |
| rle              | l_returnflag-string.csv   | 11.45               | 8.11                 | 1.41               | 5226             | 10067           |
| rle              | l_shipdate-string.csv     | 62.96               | 73.94                | 0.85               | 14362            | 10671           |
| rle              | l_shipinstruct-string.csv | 73.39               | 64.36                | 1.16               | 11228            | 10107           |
| rle              | l_shipmode-string.csv     | 30.25               | 35.74                | 0.85               | 12075            | 9851            |
| rle              | l_discount-int8.csv       | 11.97               | 10.41                | 1.15               | 1258             | 178             |
| rle              | l_linenumber-int8.csv     | 11.45               | 11.04                | 1.04               | 1356             | 167             |
| rle              | l_quantity-int8.csv       | 16.14               | 11.22                | 1.44               | 872              | 189             |
| rle              | l_tax-int8.csv            | 11.45               | 10.18                | 1.12               | 719              | 170             |
| rle              | l_discount-int16.csv      | 11.97               | 15.61                | 0.77               | 682              | 208             |
| rle              | l_linenumber-int16.csv    | 11.45               | 16.56                | 0.69               | 800              | 204             |
| rle              | l_quantity-int16.csv      | 16.14               | 16.83                | 0.96               | 895              | 227             |
| rle              | l_suppkey-int16.csv       | 27.98               | 17.17                | 1.63               | 1031             | 336             |
| rle              | l_tax-int16.csv           | 11.45               | 15.26                | 0.75               | 678              | 206             |
| rle              | l_discount-int32.csv      | 11.97               | 26.01                | 0.46               | 738              | 273             |
| rle              | l_extendedprice-int32.csv | 45.05               | 28.62                | 1.57               | 959              | 453             |
| rle              | l_linenumber-int32.csv    | 11.45               | 27.59                | 0.41               | 723              | 265             |
| rle              | l_orderkey-int32.csv      | 44.73               | 7.15                 | 6.25               | 865              | 317             |
| rle              | l_partkey-int32.csv       | 36.88               | 28.62                | 1.29               | 1478             | 467             |
| rle              | l_quantity-int32.csv      | 16.14               | 28.04                | 0.58               | 864              | 358             |
| rle              | l_suppkey-int32.csv       | 27.98               | 28.61                | 0.98               | 879              | 395             |
| rle              | l_tax-int32.csv           | 11.45               | 25.44                | 0.45               | 690              | 263             |
| rle              | l_discount-int64.csv      | 11.97               | 46.82                | 0.26               | 687              | 404             |
| rle              | l_extendedprice-int64.csv | 45.05               | 51.51                | 0.87               | 1211             | 581             |
| rle              | l_linenumber-int64.csv    | 11.45               | 49.67                | 0.23               | 853              | 396             |
| rle              | l_orderkey-int64.csv      | 44.73               | 12.87                | 3.47               | 1029             | 352             |
| rle              | l_partkey-int64.csv       | 36.88               | 51.51                | 0.72               | 1375             | 631             |
| rle              | l_quantity-int64.csv      | 16.14               | 50.48                | 0.32               | 845              | 422             |
| rle              | l_suppkey-int64.csv       | 27.98               | 51.50                | 0.54               | 1083             | 555             |
| rle              | l_tax-int64.csv           | 11.45               | 45.79                | 0.25               | 839              | 426             |
|                  |                           |                     |                      |                    |                  |                 |
| dic              | l_comment-string.csv      | 157.35              | 211.08               | 0.75               | 23826            | 15814           |
| dic              | l_commitdate-string.csv   | 62.96               | 26.05                | 2.42               | 11245            | 10027           |
|                  | l_linestatus-string.csv   | 11.45               | 11.45                | 1.00               | 10347            | 9395            |
|                  | l_receiptdate-string.csv  | 62.96               | 26.07                | 2.41               | 11796            | 9840            |
| dic              | l_returnflag-string.csv   | 11.45               | 11.45                | 1.00               | 10783            | 9377            |
|                  | l_shipdate-string.csv     | 62.96               | 26.07                | 2.41               | 11326            | 10834           |
| dic              | l_shipinstruct-string.csv | 74.39               | 11.45                | 6.50               | 10584            | 9932            |
| dic              | l_shipmode-string.csv     | 30.25               | 11.45                | 2.64               | 11404            | 9457            |
|                  | l_discount-int8.csv       | 11.97               | 5.72                 | 2.09               | 698              | 262             |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
|                  |                           |                     |                      |                    |                  |                 |
| dic              | l_discount-int8.csv       | 11.97               | 12.49                | 0.96               | 1073             | 10095           |
| dic              | l_extendedprice-int32.csv | 45.05               | 52.34                | 0.86               | 3576             | 15216           |
|                  |                           |                     |                      |                    |                  |                 |
| for              | l_linenumber-int8.csv     | 11.45               | 2.86                 | 4.00               | 740              | 161             |
| for              | l_discount-int8.csv       | 11.97               | 3.77                 | 3.18               | 943              | 157             |
| for              | l_tax-int8.csv            | 11.45               | 3.96                 | 2.89               | 779              | 174             |
| for              | l_quantity-int8.csv       | 16.14               | 9.33                 | 1.73               | 893              | 257             |
| for              | l_discount-int16.csv      | 11.97               | 5.06                 | 2.36               | 947              | 544             |
| for              | l_linenumber-int16.csv    | 11.45               | 2.86                 | 4.00               | 625              | 142             |
| for              | l_quantity-int16.csv      | 16.14               | 18.45                | 0.87               | 746              | 247             |
| for              | l_suppkey-int16.csv       | 27.98               | 22.88                | 1.22               | 1196             | 355             |
| for              | l_tax-int16.csv           | 11.45               | 5.55                 | 2.06               | 955              | 165             |
| for              | l_discount-int32.csv      | 11.97               | 10.13                | 1.18               | 703              | 183             |
| for              | l_extendedprice-int32.csv | 45.05               | 45.78                | 0.98               | 841              | 544             |
| for              | l_linenumber-int32.csv    | 11.45               | 5.72                 | 2.00               | 556              | 147             |
| for              | l_orderkey-int32.csv      | 44.73               | 45.78                | 0.98               | 791              | 539             |
| for              | l_partkey-int32.csv       | 36.88               | 45.76                | 0.81               | 893              | 541             |
| for              | l_quantity-int32.csv      | 16.14               | 6.70                 | 2.41               | 686              | 165             |
| for              | l_suppkey-int32.csv       | 27.98               | 45.19                | 0.62               | 726              | 493             |
| for              | l_discount-int64.csv      | 11.97               | 15.41                | 0.78               | 666              | 220             |
| for              | l_extendedprice-int64.csv | 45.05               | 91.57                | 0.49               | 966              | 806             |
| for              | l_linenumber-int64.csv    | 11.45               | 5.72                 | 2.00               | 594              | 154             |
| for              | l_orderkey-int64.csv      | 44.73               | 91.57                | 0.49               | 859              | 729             |
| for              | l_partkey-int64.csv       | 36.88               | 91.51                | 0.40               | 1015             | 820             |
| for              | l_quantity-int64.csv      | 16.14               | 7.85                 | 2.06               | 717              | 197             |
| for              | l_suppkey-int64.csv       | 27.98               | 90.37                | 0.31               | 968              | 758             |
|                  |                           |                     |                      |                    |                  |                 |
| dif              | l_discount-int8.csv       | 11.97               | 4.41                 | 2.71               | 637              | 145             |
| dif              | l_linenumber-int8.csv     | 11.45               | 3.23                 | 3.55               | 708              | 127             |
| dif              | l_quantity-int8.csv       | 16.14               | 9.53                 | 1.69               | 730              | 198             |
| dif              | l_tax-int8.csv            | 11.45               | 3.96                 | 2.89               | 586              | 143             |
| dif              | l_discount-int16.csv      | 11.97               | 6.63                 | 1.80               | 641              | 176             |
| dif              | l_linenumber-int16.csv    | 11.45               | 3.75                 | 3.05               | 648              | 148             |
| dif              | l_quantity-int16.csv      | 16.14               | 18.87                | 0.86               | 820              | 258             |
| dif              | l_suppkey-int16.csv       | 27.98               | 22.88                | 1.22               | 881              | 375             |
| dif              | l_tax-int16.csv           | 11.45               | 5.55                 | 2.06               | 619              | 161             |
| dif              | l_discount-int32.csv      | 11.97               | 9.74                 | 1.23               | 720              | 177             |
| dif              | l_extendedprice-int32.csv | 45.05               | 45.78                | 0.98               | 912              | 542             |
| dif              | l_linenumber-int32.csv    | 11.45               | 7.50                 | 1.53               | 635              | 156             |
| dif              | l_orderkey-int32.csv      | 44.73               | 5.72                 | 7.81               | 844              | 303             |
| dif              | l_partkey-int32.csv       | 36.88               | 45.76                | 0.81               | 1273             | 529             |
| dif              | l_quantity-int32.csv      | 16.14               | 6.67                 | 2.42               | 1398             | 925             |
| dif              | l_suppkey-int32.csv       | 27.98               | 45.19                | 0.62               | 784              | 497             |
| dif              | l_tax-int32.csv           | 11.45               | 10.51                | 1.09               | 622              | 174             |
| dif              | l_discount-int64.csv      | 11.97               | 14.56                | 0.82               | 706              | 213             |
| dif              | l_extendedprice-int64.csv | 45.05               | 91.57                | 0.49               | 970              | 832             |
| dif              | l_linenumber-int64.csv    | 11.45               | 9.59                 | 1.19               | 633              | 177             |
| dif              | l_orderkey-int64.csv      | 44.73               | 5.72                 | 7.81               | 814              | 323             |
| dif              | l_partkey-int64.csv       | 36.88               | 91.51                | 0.40               | 978              | 827             |
| dif              | l_quantity-int64.csv      | 16.14               | 7.79                 | 2.07               | 732              | 192             |
| dif              | l_suppkey-int64.csv       | 27.98               | 90.39                | 0.31               | 870              | 781             |

### 3.3 Analysis of different compression techs



### 3.4 Analysis of different files

? if need?







## 4. Usage of the code



