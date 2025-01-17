# Homework 2

## Part 1: Load Balancing

### 1.1: Load balancing static web site

For this experiment I will use a heavy workload in order to see the load balancer in action.
Workload will be 200 threads, 500 connections for 10 seconds, using the following command:

```console
$ wrk -t200 -c500 -d10s --timeout 2s http://node0/
```

#### Case 0:

1 Load balancer (node0)<br />
1 Web Server (node1)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_case0.png" alt="Case 0" title="Case 0">

#### Case 1:

1 Load balancer (node0)<br />
2 Web Servers (node1, node2)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_case1.png" alt="Case 1" title="Case 1">

#### Case 2:

1 Load balancer (node0)<br />
3 Web Servers (node1, node2, node3)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_case2.png" alt="Case 2" title="Case 2">

#### Case 3:

1 Load balancer (node0)<br />
4 Web Servers (node1, node2, node3, mode4)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_case3.png" alt="Case 3" title="Case 3">

#### Observations

##### Throughput

As we can see, in case0, the throughput is 26k requests/s but as soon as we introduce a second webserver that jumps top 161k req/s. After that, ew have a steady increase of around
20k req/s, going to 173k for case2 and 186k for case3. Probably an expected result since when we have just one webserver and the requests keep on coming faster than
the webserver can handle, a sort of request queue gets created and each request takes exponentially more time to be handled. Once the second webserver enters and the workload
is spread, no queue is formed since the load balancer can route the message to the most suitable server at the time. Same logic applies to the rest of the webservers
but now without that queue, the throughput speed is not that significant.

##### Tail Latency

The latency at the 99th percentile seems to again have a dramatic decrease of 500ms from case0 to case1 but then decrease steadily with a rate of 50ms. I would suggest
that the explanation is the same as for the throughput.

### 1.2: Load balancing HotelMap web service

For this experiment I used the setup from the first part of the experiment but replaced the contents of index.html.js
with the contents of index.html from .../labs/05-hotelapp/hotelapp/internal/frontend/static/index.html

#### Case 0:

1 Load balancer (node0)<br />
1 Web Server (node1)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_2_case0.png" alt="Case 0" title="Case 0">

#### Case 1:

1 Load balancer (node0)<br />
2 Web Servers (node1, node2)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_2_case1.png" alt="Case 1" title="Case 1">

#### Case 2:

1 Load balancer (node0)<br />
3 Web Servers (node1, node2, node3)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_2_case2.png" alt="Case 2" title="Case 2">

#### Case 3:

1 Load balancer (node0)<br />
4 Web Servers (node1, node2, node3, mode4)<br />
1 Workload provider (node5)

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part1_2_case3.png" alt="Case 3" title="Case 3">

#### Observations

It is safe to say that the observations we made on the first web app are consistent with the hotel map web app.
This is probably because we are serving static sites with no computations and therefore the content of the site
does not affect performance.

## Part 2: Web Search Characterization

### 2.1: Single- vs Multi-threaded Client

#### Outline

[frontend]
node1

[index]
node2

index_server_threads_count=10

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_cores_per_socket.png" alt="Cores per socket" title="Cores per socket">

Command: <br/>

```console
./client node1 8080 /local/websearch/ISPASS_PAPER_QUERIES_100K 1000 {thread_count} onlyHits.jsp 1 1 /tmp/out 1 > temp.txt
```

##### Case 1

Client threads = 1

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case1_cpu.png" alt="Cpu util" title="Cpu util">
<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case1_latency_throughput.png" alt="Latency and throughput" title="Latency and throughput">

##### Case 2

Client threads = 8

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case2_cpu.png" alt="Cpu util" title="Cpu util">
<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case2_latency_throughput.png" alt="Latency and throughput" title="Latency and throughput">

##### Case 3

Client threads = 64

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case3_cpu.png" alt="Cpu util" title="Cpu util">
<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case3_latency_throughput.png" alt="Latency and throughput" title="Latency and throughput">

##### Case 4

Client threads = 128

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case4_cpu.png" alt="Cpu util" title="Cpu util">
<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_case4_latency_throughput.png" alt="Latency and throughput" title="Latency and throughput">

### Observations

As we can see, cpu performance scales linearly with the number of client threads. Something we expected. Following the same pattern,
throughput is also increasing linearly and the tail latency is worsening maybe a bit faster than the client threads are getting increased.
This degradation in response times is caused by increased contentions of shared resources as the number of active cores increases.

### 2.2: Index Partitioning

Run with:

```console
./client node1 8080 /local/websearch/ISPASS_PAPER_QUERIES_100K 1000 1 onlyHits.jsp 1 1 /tmp/out 1 2> /dev/null
```

##### Case 1: hosts-2-index

[frontend]</br>
node1</br>

[index]</br>
node2</br>
node3</br>

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_2_case1.png" alt="Latency and throughput" title="Latency and throughput">

##### Case 2: hosts-2-index

[frontend]</br>
node1</br>

[index]</br>
node2</br>
node3</br>
node4</br>
node5</br>

<img src="https://github.com/lpapal03/cs499-fa22/blob/main/assignments/hw2/answers/images/part2_2_case2.png" alt="Latency and throughput" title="Latency and throughput">

### Observations

As we can see, throughput stays exactly the same and response times are quite similar except for the 99th percentile
where we have a speedup of around 32%. I was expecting to have better throughput when doubling the amount of index servers,
however this is not the case. The only thing that gets better is the response time at the 99th percentile.

Thank you for reading :)
