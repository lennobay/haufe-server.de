# Routing - Network Docs
DESCRIPTION: A docs article to routing and the usage of the traceroute tool in Linux ...
HEADIMAGE:Screenshot_20260926_090446.png
DATE: 26.09.2026
Process of finding the best path from source to destination for an IP packet.
## Operations
1. Router Operates on Level 3 & selects best route to travel
2. Data packet have IP in their header
3. Nearest router gets the packet
4. Other routers route it further
Step 3 & 4 are repeated
## Types of Routing
1. Static Routing = routes are added manually & simple but complicated in large networks
2. Dynamic Routing = route are searched by algorithms & routes adapt to network change (update automatically)
3. Defaul Routing = packets sent to gateways when no specific route is available
## Working example 
1. Communication start with an protocoll (source route)
2. Packets are broken into small packets
3. Ip Address in packtes header
4. Shortest path found by routing table
5. Travell through multiple routers (hops)
6. Max hop number & overshoot the hops = retransmittion
7. Packets are reasambled (dest. node)
