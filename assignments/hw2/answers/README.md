# Homework 2

## Part 1: Load Balancing

### 1.1: Load balancing static web site

For this experimewnt I will use a heavy workload in order to see the load balancer in action.
Workload will be 200 threads, 500 connections for 10 seconds, using the following comand:
...
wrk -t200 -c500 -d10s --timeout 2s http://node0/
...

Case 1:
1 Load balancer (node0)
2 Web Servers (node1, node2)
1 Workload provider (node6)

Case 2:
1 Load balancer (node0)
3 Web Servers (node1, node2, node3)
1 Workload provider (node6)

Case 3:
1 Load balancer (node0)
5 Web Servers (node1, node2, node3, mode4, node5)
1 Workload provider (node6)

### 1.2: Load balancing HotelMap web service

## Part 2: Web Search Characterization

throughput (requests/sec) and tail latency changes with increasing number of web servers (1 to 4
