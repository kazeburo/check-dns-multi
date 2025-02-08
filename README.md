# check-dns-multi

Monitor multiple DNS's response.

## usage

```
Usage:
  main [OPTIONS]

Application Options:
  -v, --version            Show version
      --protocol=[tcp|udp]
  -p, --port=              Port number (default: 53)
  -H, --hostname=          DNS server hostnames (default: 127.0.0.1)
  -Q, --question=          Question hostname (default: example.com.)
  -q, --querytype=[A|AAAA]
  -E, --expect=            Expect string in result
      --timeout=           Timeout (default: 5s)
      --all                Require all resolution OK

Help Options:
  -h, --help               Show this help message

```

## Check 2 DNS server status

```
$ check-dns-multi -H 8.8.8.8 -H 1.1.1.1 -Q example.com. -q A
DNS OK: [8.8.8.8:53] HEADER-> ;; opcode: QUERY, status: NOERROR, id: 61175 ;; flags: qr rd ra;
[8.8.8.8:53] ANSWER-> example.com.	62	IN	A	23.192.228.80
[8.8.8.8:53] ANSWER-> example.com.	62	IN	A	23.215.0.138
[8.8.8.8:53] ANSWER-> example.com.	62	IN	A	96.7.128.175
[8.8.8.8:53] ANSWER-> example.com.	62	IN	A	23.192.228.84
[8.8.8.8:53] ANSWER-> example.com.	62	IN	A	23.215.0.136
[8.8.8.8:53] ANSWER-> example.com.	62	IN	A	96.7.128.198
[1.1.1.1:53] HEADER-> ;; opcode: QUERY, status: NOERROR, id: 16450 ;; flags: qr rd ra;
[1.1.1.1:53] ANSWER-> example.com.	268	IN	A	23.192.228.84
[1.1.1.1:53] ANSWER-> example.com.	268	IN	A	96.7.128.175
[1.1.1.1:53] ANSWER-> example.com.	268	IN	A	23.215.0.136
[1.1.1.1:53] ANSWER-> example.com.	268	IN	A	23.215.0.138
[1.1.1.1:53] ANSWER-> example.com.	268	IN	A	96.7.128.198
[1.1.1.1:53] ANSWER-> example.com.	268	IN	A	23.192.228.80
```