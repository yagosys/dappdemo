config server-policy vserver
  edit "hyperledgerest"
    config  vip-list
      edit 1
        set interface port1
      next
    end
  next
end

config server-policy policy
  edit "hyperledge8888"
    set ssl enable
    set vserver hyperledgerest
    set service hyperledgerest8888
    set replacemsg Predefined
    set server-pool hyperledgerest3000
    config  http-content-routing-list
    end
    set tlog enable
  next
end

config server-policy server-pool
  edit "hyperledgerest3000"
    set server-balance enable
    set health HLTHCK_ICMP
    config  pserver-list
      edit 1
        set ip 172.17.0.1
        set port 3000
      next
    end
  next
end
