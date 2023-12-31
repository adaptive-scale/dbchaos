# DBChaos

Stress-test your database with random or pre-defined queries. Validate slow and expensive queries that break your database.

## Installation

```shell
go install github.com/adaptive-scale/dbchaos@v0.4.1
```

## Run your first test

Create a file name `config.yaml` with the following content:
```yaml
dbType: postgres
connection: "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
query: |
  SELECT pg_database.datname as "Database", pg_size_pretty(pg_database_size(pg_database.datname)) as "Size"
  FROM pg_database;
parallelRuns: 100
runFor: 30m
```

To run the above config file:

```shell
dbchaos runTest config.yaml
```

## Run bunch of queries in parallel

Create a file called `scenario.yaml` with the following content:

```yaml
dbType: mysql
connection: "root:root@tcp(host:port)/db"
scenarios:
  - query: select * from information_schema.statistics
    parallelRuns: 10000
    runFor: 15m
  - query: |
      SELECT table_schema "Database", ROUND(SUM(data_length + index_length) / 1024 / 1024, 2) "Size (MB)"
      FROM information_schema.tables
      GROUP BY table_schema;
    parallelRuns: 10000
    runFor: 15m
```

To run the above scenario file:

```shell
dbchaos runScenario scenario.yaml
```
