# Homework 1 - Part 2

## Experiment Results
---

#### Machines:
node0: c220g2 Wisconsin
node1: c220g1 Wisconsin

### 1. Local memory
#### Latency
After performing the mlc command, we se that the idle latency from node0 to node0 is 93.7ns.
#### Bandwidth
The mlc command gives us the bandwidth with varying read-write ratios. Calculating the average of these results gives us 29574MB/s.
#### Memory
The physical memory of the machine is 164.93GB as stated by the meminfo command.


### 2. Local Disk
#### Latency
The ioping command gave us an average latency of 169.4us (=169400ns)
#### Bandwidth
Executing the fio command, we get an average bandwidth of 186409.45KiB/s (=190.88MB/s)
#### Memory
The hdparm command suggests that the disk has size 480GB.

### 3. Remote Memory and Disk
#### Latency
Pinging node1 took 0.158ms (=158000ns). We need to add to that the latency for a disk read, which as we saw earlier is around 169400ns.
That gives us a total latency of 
#### Bandwidth
After trying 10Gbits/s, the network responded with around 5.1Gbit/s, meaning this is the max it can go. Therefore, we have 673.5MB/s (=5.1Gbit/s). However, the bottleneck of this network will be the bandwidth between the remote server and its disk, which as we saw is around
190MB/s
#### Memory
I was not able to find an accurate measurement for the collective memory of remote disks, so I assumed 1 Petabyte, a reasonable for
medium sized data center. 

### Graph observations
---
The graph suggests that as we move from local memory to remote memory, the latency has a huge jump from local memory to local disk. Something
expected since disk I/O operations take significantly longer than local I/Os. Then, from local disk to remote memory, latency is increased slightly. Again expected, since the only added delay is the network. We can also observe that from local memory to local disk the bandwidth
has also a large dip but then somewhat stabilizes. That is because the bandwidth inside a single machine is way grater than the bottleneck
that we face once we leave said machine. Finally, the memory capacity stadily increases from local memory to disk, but then has a massive jump
when we leave the machine. Again, something expected since the datacenter as a whole can hold way more data than a single disk can.  
