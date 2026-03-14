# go-nplus1-example

## Simulation results

![](./benchmark/result.png)

## How to simulate

```bash
> docker compose up -d

[+] Running 1/2
 ⠹ Network go-nplus1-example_default  Created                                                                                                                                                               0.2s 
 ✔ Container go-nplus1-example-db-1   Started                                                                                                                                                               0.2s 

> go run ./benchmark
nplus1 10 0.13553379200000001
nplus1 100 0.428294708
nplus1 1000 3.456485125
nplus1 10000 33.357696791
in_clause 10 0.050162584
in_clause 100 0.053320957999999995
in_clause 1000 0.08453812499999999
in_clause 10000 0.09685408300000001
join 10 0.050517916999999996
join 100 0.053924
join 1000 0.0851305
join 10000 0.08542274999999999
./benchmark/scaling.png generated

> open ./benchmark/scaling.png

> docker compose down -v

[+] Running 2/2
 ✔ Container go-nplus1-example-db-1   Removed                                                                                                                                                               0.1s 
 ✔ Network go-nplus1-example_default  Removed 
```