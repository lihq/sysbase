# About

System base metrics.

系统基础监控项采集实现。


## Network bandwidth

网卡流量统计

注意：多张网卡接口需要用 tags 分类。

     
Grafana 配图用 order by ts asc，平时调试 SQL 用 order by ts desc。例：

    select date_trunc('minute', "ts") as time, metric, sum(value) from datapoints where metric='net.in.bytes' and tags->>'face'='enp0s3' group by ts, metric order by ts asc;
    
    
查看网卡流量实时变化

    watch -n 1  -d 'cat /proc/net/dev'
