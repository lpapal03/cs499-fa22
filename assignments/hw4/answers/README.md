# Homework 4 Report

## Part 1
Configuration:
```
node0     Workload generator, Docker Swarm leader
nodes1-5  Docker Swarm workers
```

To generate workloads we used
```
./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://{node}:{port for mongodb-rate given by stack services}
```

We noticed that ``` ../wrk2/scripts/hotel-reservation/mixed-workload_type_1.lua ``` directory has the following code at the end:

```
request = function()
  cur_time = math.floor(socket.gettime())
  local search_ratio      = 0.6
  local recommend_ratio   = 0.39
  local user_ratio        = 0.005
  local reserve_ratio     = 0.005

  local coin = math.random()
  if coin < search_ratio then
    return search_hotel(url)
  elseif coin < search_ratio + recommend_ratio then
    return recommend(url)
  elseif coin < search_ratio + recommend_ratio + user_ratio then
    return user_login(url)
  else 
    return reserve(url)
  end
```
So we changed it in order to have only 200 result codes (HTTP OK):
```
request = function()
--   cur_time = math.floor(socket.gettime())
--   local search_ratio      = 0.6
--   local recommend_ratio   = 0.39
--   local user_ratio        = 0.005
--   local reserve_ratio     = 0.005

--   local coin = math.random()
--   if coin < search_ratio then
--     return search_hotel(url)
--   elseif coin < search_ratio + recommend_ratio then
--     return recommend(url)
--   elseif coin < search_ratio + recommend_ratio + user_ratio then
--     return user_login(url)
--   else 
--     return reserve(url)
--   end
    return search_hotel(url)
end
```
We commented out all the function code and we always return search_hotel(url), because we don't have the other services.

We then ran ``` ./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://127.0.0.1:8080 ``` in node0 (Manager):

```
Running 30s test @ http://127.0.0.1:8080
  2 threads and 100 connections
  Thread calibration: mean lat.: 3.956ms, rate sampling interval: 10ms
  Thread calibration: mean lat.: 3.970ms, rate sampling interval: 10ms
  Thread Stats   Avg      Stdev     99%   +/- Stdev
    Latency     3.91ms  783.75us   6.40ms   75.67%
    Req/Sec     1.05k   147.83     1.44k    63.05%
#[Mean    =        3.908, StdDeviation   =        0.784]
#[Max     =       12.816, Total count    =        39899]
#[Buckets =           27, SubBuckets     =         2048]
----------------------------------------------------------
  59889 requests in 30.00s, 15.87MB read
Requests/sec:   1996.43
Transfer/sec:    541.61KB
```
The second time we ran it, we got completely different results (except req/sec and transfer/sec):
```
Running 30s test @ http://127.0.0.1:8080
  2 threads and 100 connections
  Thread calibration: mean lat.: 10.918ms, rate sampling interval: 32ms
  Thread calibration: mean lat.: 12.299ms, rate sampling interval: 32ms
  Thread Stats   Avg      Stdev     99%   +/- Stdev
    Latency    11.39ms    3.50ms  20.08ms   68.14%
    Req/Sec     1.02k   602.00     1.61k    76.77%
#[Mean    =       11.395, StdDeviation   =        3.496]
#[Max     =       25.840, Total count    =        39900]
#[Buckets =           27, SubBuckets     =         2048]
----------------------------------------------------------
  57912 requests in 30.01s, 15.35MB read
Requests/sec:   1930.02
Transfer/sec:    523.69KB
```
So we decided to run it 25 times and get the average:
```
for i in {1..25}; do ./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://127.0.0.1:8080 | grep Latency >> lat_results.txt; done
```

We did these processes both for scaled profile service and not. We scaled our profile microservice by 3

![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/scaled.png "Scaled vs Not Scaled")

For scaled, the Search/Nearby gRPC was the bottleneck with average time ~59ms and GetProfiles had an average time ~57ms. 
For not scaled, the Search/Nearby gRPC was the bottleneck with average time ~148ms and GetProfiles had an average time ~145ms. 

## Part 2

Configuration:
```
node0     Workload generator
node1     Target server (case of single machine), Docker Swarm leader (case of multiple machines)
node2-4   Docker Swarm workers
```

To generate workloads we used
```
./wrk -t2 -c100 -d30s -R2000 -L -s ./scripts/hotel-reservation/mixed-workload_type_1.lua http://{node}:{port for mongodb-rate given by stack services}
```

Tutorial for [using swarm with docker-compose (deploy stack)](https://docs.docker.com/engine/swarm/stack-deploy/)

### MongoDB

#### Single node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/mongodb_single.png "MongoDB single node")

#### Multi node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/mongodb_multi.png "MongoDB multi node") 

### Memcached

#### Single node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/memcached_single.png "Memcached single node")

#### Multi node
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/memcached_multi.png "Memcached multi node") 

### Results analysis

#### Latency
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/latency.png "Latency")
Regarding latency, what we observe is that in both cases MongoDB and Memcached, the latency starts off small in the single node with a value of around 2 and 4 respectively, and there is a dramatic increase to around 300 and 400ms. This can be explained by the bottleneck that the network introduces and therefore is expected. What was not expected is the fact that the memcached latency is higher in both single and multinode implementations. To some degree, this can be explained by cache misses and the system trying to fill the cache with relevant values. However, with the workload we have, we would expect that by some point the cache would produce better average results than MongoDB. Maybe this difference in values is due to the implementations of each technology and not to the database and cache as concepts in a system. 

#### Throughput
![alt text](https://github.com/cseas002/cs499-fa22/blob/main/assignments/hw4/answers/results/throughput.png "Throughput")
Similarly here in throughput, both implementations MongoDB and Memcached start of well and have a drop in performance, although not as dramatic as in latency. Again, we can see that the Memcached implementation performs worse but in this case it is probably the expected result since the Memcached implementation does more operations per request and therefore takes longer for each request to get resolved.

#### Conclusion
Apparently, there is not a "good" or "bad" solution, either that being a database with a cache or not, or deplying to single or multiple nodes. Everything depends on the use case and we can only find the optimal implementation by testing

Thank you for reading :)
