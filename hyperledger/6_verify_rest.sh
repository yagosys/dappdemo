#!/bin/bash -xe 
endpoint="172.17.0.2"
curl --request GET   --url "http://$endpoint:8888/query?channelid=mychannel&chaincodeid=basic&function=ReadAsset&args=Asset123"
#curl --request GET   --url 'http://127.0.0.1:3000/query?channelid=mychannel&chaincodeid=basic&function=ReadAsset&args=Asset123'
