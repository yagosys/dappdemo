#!/bin/bash -xe
echo map fortiweb admin port to 8443. 
echo map fortiweb vip port 8544 to external port 8546 , 8546 is the port for metamask to connect
echo default password is empty
myip=$(ip -4 -o -br -j add show eth0 | jq .[].addr_info[].local -r)
echo $myip
sleep 30
#docker run -it --rm --privileged -p 172.31.80.71:8443:43 --cap-add=ALL interbeing/myfmg:fweb70577
#8444:8888 is reversed for other applicaiton like nginx
docker run -it --rm --privileged -p $myip:8443:43 -p $myip:8444:8888 -p $myip:8546:8544 --cap-add=ALL interbeing/myfmg:fweb70577
