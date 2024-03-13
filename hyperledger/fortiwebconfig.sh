#!/bin/bash -x

echo "please use fortiweb GUI to upload hyperledger_fabric_gateway.proto as grpc idl file"
echo "change serverip to match your actual ip for hyperledger machine"
serverip="172.31.49.189"
echo process ctrl-c to break if you want change serverip
sleep 10
cat << EOF | tee fortiwebconfig.txt 

config waf grpc-idl file
  edit "hyperledger_fabric_gateway.proto"
  next
end

config waf grpc-security rule
  edit "hyperledgernew"
    set url /gateway.Gateway/*
    set idl-file hyperledger_fabric_gateway.proto
    set rate-limit 1
    set size-limit 100
    set req-message-name gateway.EvaluateRequest
    set rsp-message-name gateway.EvaluateResponse
  next
end

config waf grpc-security policy
  edit "hyperledger-new"
    set enable-signature-detection enable
    config  rule-list
      edit 1
        set rule hyperledgernew
      next
    end
  next
end


config server-policy vserver
  edit "hyperledge3000"
    config  vip-list
      edit 1
        set interface port1
        set use-interface-ip enable
      next
    end
  next
end

config server-policy server-pool
  edit "hyperledgerest3000"
    set comment hyperledgerrestapi
    config  pserver-list
      edit 1
        set ip $serverip
        set port 3000
      next
    end
  next
end

config server-policy service custom
  edit "hyperledgerrest3000"
    set port 3000 
  next
end

config waf custom-access rule
  edit "hyperledger_rest_rule"
    set action alert_deny
    config  source-ip-filter
    end
    config  geo-filter
    end
    config  time-range-filter
    end
    config  method
    end
    config  url-filter
    end
    config  http-header-filter
    end
    config  parameter
      edit 1
        set name function
        set location-check enable
      next
    end
    config  user-filter
    end
    config  access-limit-filter
    end
    config  http-transaction
    end
    config  response-code
    end
    config  content-type
    end
    config  packet-interval
    end
    config  main-class
    end
    config  sub-class
    end
    config  custom-signature
    end
    config  occurrence
    end
  next
end

config waf custom-access policy
  edit "hyperledger_rest_rule"
    config  rule
      edit 1
        set rule-name hyperledger_rest_rule
      next
    end
  next
end

config waf web-protection-profile inline-protection
  edit "hyperledger_rest_policy"
    set custom-access-policy hyperledger_rest_rule
    set ip-intelligence enable
  next
end

config server-policy policy
  edit "hyperledgerest3000"
    set ssl enable
    set vserver hyperledge3000
    set service hyperledgerrest3000
    set web-protection-profile hyperledger_rest_policy
    set replacemsg Predefined
    set server-pool hyperledgerest3000
    config  http-content-routing-list
    end
    set tlog enable
  next
end

config server-policy server-pool
  edit "hyperledge9051"
    set server-pool-id 3707592029084994372
    config  pserver-list
      edit 1
        set $serverip
        set port 9051
        set ssl enable
        set client-certificate grpcsample
        set tls-v10 disable
        set tls-v11 disable
        set http2 enable
      next
    end
  next
end



config server-policy vserver
  edit "hyperledger9052"
    config  vip-list
      edit 1
        set interface port1
        set use-interface-ip enable
      next
    end
  next
end

config server-policy policy
  edit "hyperledger9052"
    set ssl enable
    set vserver hyperledger9052
    set web-protection-profile hyperledge-new
    set replacemsg Predefined
    set server-pool hyperledge9051
    set https-service hyperledger9052
    set certificate grpcsample
    set tls-v10 disable
    set tls-v11 disable
    set ssl-noreg disable
    config  http-content-routing-list
    end
    set http2 enable
    set tlog enable
  next
end
EOF

