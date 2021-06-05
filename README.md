# 1kv-exporter

A simple exporter that scrapes the thousand node validator API every few minutes, and exposes that data as a 
prometheus endpoint. By default exposes the metrics at http://[host]:17586/metrics

_Note: this is alpha quality software (at best,) very little testing has been done to check correctness. 
If you use this at this early stage, expect it to be wrong. Contributions are most-welcome, but also feel free
to open a ticket and I'll see what I can do. It will improve as I experience various conditions and get through
a few nomination cycles._

## Options:

```
Usage of 1kv-exporter
  -c string
        Which chain to monitor, possible choices are 'kusama' or 'polkadot' (default "kusama")
  -p int
        Port to listen on (default 17586)
  -s string
        Validators to monitor, comma seperated list of stash addresses, alternate: 'VALIDATORS' env var
  -t int
        Update frequency in minutes (default 5)
```

## Example output:

```
# HELP onekv_active Validator is active
# TYPE onekv_active gauge
onekv_active{chain="kusama",validator_name="some-awesome-validator"} 1
# HELP onekv_bonded Bonded tokens
# TYPE onekv_bonded gauge
onekv_bonded{chain="kusama",validator_name="some-awesome-validator"} 64
# HELP onekv_commission Validator commission
# TYPE onekv_commission gauge
onekv_commission{chain="kusama",validator_name="some-awesome-validator"} 5
# HELP onekv_fault_events The total number of fault events
# TYPE onekv_fault_events gauge
onekv_fault_events{chain="kusama",validator_name="some-awesome-validator"} 0
# HELP onekv_faults Validator faults
# TYPE onekv_faults gauge
onekv_faults{chain="kusama",validator_name="some-awesome-validator"} 0
# HELP onekv_inclusion Validator inclusion
# TYPE onekv_inclusion gauge
onekv_inclusion{chain="kusama",validator_name="some-awesome-validator"} 0.3333333333333333
# HELP onekv_last_nominated Seconds since last nomination
# TYPE onekv_last_nominated gauge
onekv_last_nominated{chain="kusama",validator_name="some-awesome-validator"} 68674.263499
# HELP onekv_last_valid Seconds since last valid
# TYPE onekv_last_valid gauge
onekv_last_valid{chain="kusama",validator_name="some-awesome-validator"} 1095.263497
# HELP onekv_offline_accumulated Accumulated offline time
# TYPE onekv_offline_accumulated gauge
onekv_offline_accumulated{chain="kusama",validator_name="some-awesome-validator"} 0
# HELP onekv_offline_since Current offline time
# TYPE onekv_offline_since gauge
onekv_offline_since{chain="kusama",validator_name="some-awesome-validator"} 0
# HELP onekv_rank Current 1kv rank
# TYPE onekv_rank gauge
onekv_rank{chain="kusama",validator_name="some-awesome-validator"} 10
# HELP onekv_span_inclusion Span Inclusion
# TYPE onekv_span_inclusion gauge
onekv_span_inclusion{chain="kusama",validator_name="some-awesome-validator"} 0.10714285714285714
# HELP onekv_unclaimed_eras Unclaimed eras
# TYPE onekv_unclaimed_eras gauge
onekv_unclaimed_eras{chain="kusama",validator_name="some-awesome-validator"} 0
```

## Running

1. install golang
1. run `go get github.com/blockpane/1kv-exporter/cmd/1kv-exporter`
1. The binary should be in $GOPATH/bin/1kv-exporter (GOPATH is usually ~/go)

**Or with docker**

1. `git clone https://github.com/blockpane/1kv-exporter.git && cd 1kv-exporter`
1. `docker build -t 1kv-exporter .`
1. `docker run -d --restart unless-stopped --name 1kv-exporter -p 17586:17586 1kv-exporter -s EJCY3aaaa...`

Then add it to the prometheus config with something like:

```yaml
  - job_name: "1kv"
    scrape_interval: 5m
    static_configs:
     - targets: ["hostname:17586"]
```

## Contributing

Fork it, modify it, open an issue, open a pull request that references the issue, and ping me on matrix: @blockpane:matrix.org

## TODOs (contributions welcome)

* add more awesome stats
* add example grafana dashboard

## Thanks

If you find this useful delegator votes are the best way to say thanks: `EJCY3iiVFgHH3RQn5EPKA1opf1oHRAbeiqsoCnXUQxpWp9k`
