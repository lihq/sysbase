# About

Linux system base monitor metrics collector.

Linux 系统基础监控项采集实现。


## Network bandwidth

网卡流量统计

注意：多张网卡接口需要用 tags 分类。

     
Grafana 配图用 order by ts asc，平时调试 SQL 用 order by ts desc。例：

TimescaleDB SQL

    select date_trunc('minute', "ts") as time, metric, sum(value) from datapoints where metric='net.in.bytes' and tags->>'face'='enp0s3' group by ts, metric order by ts asc;

  
Clickhouse SQL

    SELECT
        $timeSeries as t,
        value as `net.in.bytes`
    FROM $table
    WHERE
        $timeFilter
        and endpoint= '$endpoint'
        and metric = 'net.in.bytes'	
        and visitParamExtractString(tags, 'face') = '$iface'
    ORDER BY t

    
查看网卡流量实时变化

    watch -n 1  -d 'cat /proc/net/dev'
