[TOC]

# Report for Assignment 1

## 0. Group member

The group number is 6.

| **Name** | **Student  Number** | **Tasks**         |
| -------- | ------------------- | ----------------- |
| Jie Chen | S4162315            | All code & report |

## 1. Introduction

 

 

## 2. The environment for the experiment

### 2.1 Hardware environment

Here is the hardware environment of my experiment, and you don't need to use the same environment as mine. 

* CPU: vendor, model, generation, clockspeed, cache size

  

* Main memory

* Disk: size & speed

### 2.2 Software environment

Download all of the software and config them properly to repeat the experiment.

* MonetDB

  Version: v11.51.3 (Aug2024)

* MySQL

  Version: Ver 8.3.0 for macos14.2 on arm64

* Python

  Version: 3.12.2

* DBeaver – as the client to MonetDB and MySQL

  Version: 24.2.0.202409011551

## 3.Procedures to verify the SF-1 results

First, I download the `MonetDB` and `MySQL`, and follow the official instruction to set up the username, password and port. Then I use the tool `DBeaver` to connect to the database of `MonetDB` and `MySQL`. `DBeaver` is a user friendly database client tool, and I will execute all of the sql script through it.

I use the provided scripts and data from BrightSpace directly, and follow the instructions in the task 1.

### 3.1 MonetDB

#### 3.1.1 Load data

Create 2 databases called `SF1` and `SF3`. 

Use the DBeaver to connect to these 2 databases separately. 

Import the create table script `.../dbgen/MonetDB/0-create_tables.sql` into the `SF1`, and then run the script to create table.

Import the load data script `.../dbgen/MonetDB/1-load_data.SF-1.sql` into the `SF1`. I change the data location to the absolute file path (`…/dbgen/SF-1/data/xxx.tbl`) for each line and then run the script to load the SF-1 data.

Finally, I import the add_constraints  script `.../dbgen/MonetDB/2-add_constraints.sql` into the `SF1`, and then run it to set constraints.

####  3.1.2 Run queries

Import all of the queries into the database `SF1`, and then run these sql scripts one by one. Once we get the results of each script, we can use the tool of `DBeaver` to export the data with the specific format.

![image-20240922110634205](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922110634205.png)

 Just click on the button on the bottom, and then config the format as the screenshot shows. Then we get all of the output file from the queries. Check the `README` file located in the `.../dbgen/check_answers/` to find the file format requiements, and adjust the configration in the export file process if necessary.

Make sure the output file format satisfy the requirements and then store them into the folder `/Users/jay/Desktop/monetdb_sf1_out`.

#### 3.1.3 Verify the results

I enter the file path `.../dbgen/check_answers/`, and run the command.

```bash
# command to verify the results
./pairs.sh /Users/jay/Desktop/monetdb_sf1_out /Users/jay/Desktop/lessons material/2024-2/Advanced data management/assignment 01/TPC-H_V3.0.1++/dbgen/answers
```

The file path `/Users/jay/Desktop/monetdb_sf1_out` is the place I store the output files of previous queries. And the file path `/Users/jay/Desktop/lessons material/2024-2/Advanced data management/assignment 01/TPC-H_V3.0.1++/dbgen/answers` is the folder to store the official correct output files.

Then the script generate folders including `./data_01`, `./data_management`, `./management`, `./material`. After I check the content of these files, I find all of these folders are used to store the comparison log files and result files which is shown as follows,

![A screenshot of a phone  Description automatically generated](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/clip_image001.png)

At the mean time, the terminal shows all of these logs which indicate all the results of 22 queries find 0 unacceptable missmatches. And we can verify that all the results are correct.

![image-20240922111523727](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922111523727.png)

### 3.2 MySQL

#### 3.2.1 Load data

When I try to run the load data scripts in the MySQL database, I find some errors and constraints. So I choose to copy import the data from MonetDB tables to MySQL tables with the DBeaver. An example is shown as below.

![image-20240922115618420](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922115618420.png)

To make sure all of the data have been successfully imported, I run the query command to check the total row in these tables.

For example, I find 1500000 rows in table orders in the mysql database which is same as that in MonetDB.

![image-20240922120913064](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922120913064.png)

Finally, I import the add_constraints  script `.../dbgen/MonetDB/2-add_constraints.sql` into the `SF1`, and then run it to set constraints. But it takes far more time in MySQL than in MonetDB, almost 537s!

#### 3.2.2 Run queries

Use the same sql scripts as the MonetDB to execute the queries. Several sql scripts need to change the syntax to MySQL 8.0. I have to change the syntax in `q01.sql`, `q15.sql`. Then I store all of the output files into the folder `/Users/jay/Desktop/mysql_sf1_out`

#### 3.2.3 Verify the results

I use the same script to verify the results. I also move the answers into the foleder `/Users/jay/Desktop/answers`.

Run the command:

```bash
# command to verify the results
./pairs.sh /Users/jay/Desktop/mysql_sf1_out /Users/jay/Desktop/answers 
```

Similarly, I get the results from terminal, which means all the results are correct.

![image-20240922125301897](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922125301897.png)

## 4. Compare query execution

Repeat the previous steps of `SF1`, and load all of the data and scripts into database `SF3` of MonetDB and MySQL separately. 

Then loading data process is really slow in MySQL database, and the table `lineitem` costs almost 32min to complete.

Then I run all of these scripts and record the execution time for each. For each query, I add a `TRACE` to obtain the execution time, for example:

```sql
TRACE select
	l_returnflag,
	l_linestatus,
	sum(l_quantity) as sum_qty,
	sum(l_extendedprice) as sum_base_price,
	sum(l_extendedprice * (1 - l_discount)) as sum_disc_price,
	sum(l_extendedprice * (1 - l_discount) * (1 + l_tax)) as sum_charge,
	avg(l_quantity) as avg_qty,
	avg(l_extendedprice) as avg_price,
	avg(l_discount) as avg_disc,
	count(*) as count_order
from
	lineitem
where
	l_shipdate <= date '1998-12-01' - interval '90' day (3)
group by
	l_returnflag,
	l_linestatus
order by
	l_returnflag,
	l_linestatus;
```

To run the experiment in **warm memory**, I run each query for more then five times and then calculate the mean execution time (except the first execution time) of the query as is shown in the following figure. The first execution is the **cold run**, and we should set all the experiment in **hot run**.

![image-20240922170856400](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922170856400.png)

 Some queries may cost many more times, and I will run for fewer times and also calculate the mean execution time. Then I get the execution time table.

| Query | DBMS    | Dataset | Execution time/ms           |
| ----- | ------- | ------- | --------------------------- |
| q01   | MonetDB | SF1     | 174                         |
| q02   | MonetDB | SF1     | 70                          |
| q03   | MonetDB | SF1     | 73                          |
| q04   | MonetDB | SF1     | 59                          |
| q05   | MonetDB | SF1     | 68                          |
| q06   | MonetDB | SF1     | 50                          |
| q01   | MonetDB | SF3     | 395                         |
| q02   | MonetDB | SF3     | 75                          |
| q03   | MonetDB | SF3     | 127                         |
| q04   | MonetDB | SF3     | 112                         |
| q05   | MonetDB | SF3     | 89                          |
| q06   | MonetDB | SF3     | 62                          |
| q01   | MySQL   | SF1     | 6120                        |
| q02   | MySQL   | SF1     | 119                         |
| q03   | MySQL   | SF1     | 2250                        |
| q04   | MySQL   | SF1     | 568                         |
| q05   | MySQL   | SF1     | 943                         |
| q06   | MySQL   | SF1     | 1383                        |
| q01   | MySQL   | SF3     | 20760                       |
| q02   | MySQL   | SF3     | longer than 1h, use 3600000 |
| q03   | MySQL   | SF3     | 9413                        |
| q04   | MySQL   | SF3     | 12457                       |
| q05   | MySQL   | SF3     | 617194                      |
| q06   | MySQL   | SF3     | 6683                        |

Then I plot the results to compare the execution time with different queries and different DBMS. The value of execution time spread widely, and I use the log scale to put all these data into one figure.

![image-20240922202646003](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922202646003.png)

![image-20240922202656154](https://raw.githubusercontent.com/infinityjay/myImageHost/main/typora/image-20240922202656154.png)

From the figures above, we can find that for MonetDB, the query 1 is the most time consuming. For MySQL, the query performance are very unstable

## 5. Implementation of the queries in Python

 

 

## 6. Performance comparison between DBMS and Python implementation

 

 

 

 

 

## *Appendix

Some notes for the configuration and commands

### *.1 MonetDB

#### *.1.1 Configuration

| Description                 | Configuration   |
| --------------------------- | --------------- |
| Configuration file location | ~/.monetdb      |
| Default username/password   | monetdb/monetdb |
| Port                        | 54321           |
| Language                    | sql             |

 

#### *.1.2 Commands

| Description                            | Command                         |
| -------------------------------------- | ------------------------------- |
| Create workspace                       | `monetdbd create  ~/my-dbfarm`  |
| Check configuration                    | `monetdbd get all  ~/my-dbfarm` |
| Start server                           | `monetdbd start  ~/my-dbfarm`   |
| Create database                        | `monetdb create  my-first-db `  |
| Start database                         | `monetdb start  my-first-db`    |
| Check database status                  | `monetdb status`                |
| Release database (if locked)           | `monetdb release  my-first-db`  |
| Connect to the database                | `mclient -dmy-first-db`         |
| Stop monetdb daemon process completely | `monetdbd stop  ~/my-dbfarm`    |