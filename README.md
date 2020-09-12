# FundingsService: NSF Analyzing Tool as Service

## Quick Start

1. `sysctl -w vm.max_map_count=262144` to make sure elasticsearch can start-up.
2. configure mysql connection settings.
3. set env `BROKERS` `DB_PASSWD` `DB_HOST` `TOPIC`
4. run docker-compose in mini_cluster
5. get start with the service.

