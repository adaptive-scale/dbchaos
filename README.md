# DBChaos

Stress-test your database with pre-defined queries. Validate slow and expensive queries that breaks your database.
  
### Installation

```shell
go install github.com/adaptive-scale/dbchaos@v0.4.1
```

#### Supported Databases

- Postgres
- MySQL
- MongoDB
- SQL Server

### Run your first test

Create a file named `config.yaml` with the following content:
```yaml
dbType: postgres
connection: "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
query: |
  SELECT pg_database.datname as "Database", pg_size_pretty(pg_database_size(pg_database.datname)) as "Size"
  FROM pg_database;
parallelRuns: 100
runFor: 30m
```

For MongoDB, the connection string should be in the following format:
```yaml
dbType: postgres
connection: "mongodb://root:example@localhost:27017/"
query: |
    {"insert": "users", "documents": [{ "user": "abc123", "status": "A" }]}
parallelRuns: 100
runFor: 30m
```

To run the above config file:

```shell
dbchaos runTest 
```

### Run bunch of queries in parallel

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
dbchaos runScenario 
```

### MongoDB Specific:
Example `scenario.yaml` file: 
```yaml
dbType: mongodb
connection: "mongodb://root:example@localhost:27017/"
scenarios:
  - query: '{"insert": "users", "documents": [{ "user": "abc123", "status": "A" }]}'
    parallelRuns: 10000
    runFor: 15m
dbName: users   #(MongoDB only)
```

### Static Synthetic Data Generation

DBChaos can also generate full schema and synthetic data for your database.

```yaml
connection: 
  dbType: postgres
  connection: "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
dryRun: false
schema: 
  numberOfSchema: 10
  generateTables: true
  language: en
tables:
  numberOfTables: 10
  minColumns: 5
  maxColumns: 10
  populateTable: true
rows:
  minRows: 100
  maxRows: 1000
```