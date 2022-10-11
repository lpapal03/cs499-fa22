# Homework 2

## Part 1: Load Balancing

### 1.1: Load balancing static web site

For this experimewnt I will use a heavy workload in order to see the load balancer in action.
Workload will be 200 threads, 500 connections for 10 seconds, using the following comand:

```console
$ wrk -t200 -c500 -d10s --timeout 2s http://node0/
```

#### Case 0:

1 Load balancer (node0)<br />
1 Web Server (node1)<br />
1 Workload provider (node5)

<img src="/images/part1_case0.png" alt="Alt text" title="Optional title">

#### Case 1:

1 Load balancer (node0)<br />
2 Web Servers (node1, node2)<br />
1 Workload provider (node5)

#### Case 2:

1 Load balancer (node0)<br />
3 Web Servers (node1, node2, node3)<br />
1 Workload provider (node5)

#### Case 3:

1 Load balancer (node0)<br />
4 Web Servers (node1, node2, node3, mode4)<br />
1 Workload provider (node5)

### 1.2: Load balancing HotelMap web service

## Part 2: Web Search Characterization

throughput (requests/sec) and tail latency changes with increasing number of web servers (1 to 4
