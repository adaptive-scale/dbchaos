# DBChaos
 
Stress-test your database with pre-defined queries, generate synthetic data in your database. Validate slow and expensive queries that breaks your database.

### Features
- Synthetic Event Generation
- Synthentic Data Generation
  
### Installation

```shell
go install github.com/adaptive-scale/dbchaos@v0.4.4
```

### Supported Databases

| Database  | Synthetic Event Generation | Synthetic Data Generation    |
| ------------- | -----------------------|-------------------------------|
| Postgres | ✅ | ✅ |
| MySQL  | ✅  | ✅ |
| SQL Server  | ✅  | ✅ |
| MongoDB  | ✅  | ⛔ |

### Synthetic Event Generation
With DBChaos, you can run parallel queries on your target database. There two ways you could do that - test and scenarios.

With test, you can run single query on your target database. It would run the query parallely for the given amount of time. With Scenario, you could run multiple queries with different timeout and rates creating diverse load patterns.

We are planning to add more features around creating various load patterns.

#### Run your first test

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
dbType: mongodb
connection: "mongodb://root:example@localhost:27017/"
query: |
    {"insert": "users", "documents": [{ "user": "abc123", "status": "A" }]}
parallelRuns: 100
runFor: 30m
dbName: users
```

To run the above config file:

```shell
dbchaos runTest 
```

#### Run bunch of queries in parallel

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

For MongoDB Specific, an example `scenario.yaml` file would look as follows : 
```yaml
dbType: mongodb
connection: "mongodb://root:example@localhost:27017/"
scenarios:
  - query: '{"insert": "users", "documents": [{ "user": "abc123", "status": "A" }]}'
    parallelRuns: 10000
    runFor: 15m
dbName: users   #(MongoDB only)
```

### Synthetic Data Generation
DBChaos can generate full schema and synthetic data for your database. In DBChaos, there are two kinds of data generation techniques - Static and GPT-based. 

In static data generation, dbchaos randomly generates schema, schema name, column names and data. It comes very handy and is inexpensive if you want to create huge schemas and generate large amount data. For instance, At [Adaptive](https://adaptive.live) we use this is create unrealistic sized databases and schema to load testing our services and processes.

In GPT based data generation, you can create hyper-realistic databases and data. However, you would need an API key from OPENAI as well as it will cost you credits if you which to generate huge amount to data. We have tried to build a known schema cache in the product, which we will keep improving as well built out more features. 

#### Static Data Generation

A configuration for static generation looks as follows:
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

Save above config as `config.yaml` and run the following command:
```shell
dbchaos generate
```

#### GPT-based Synthetic Data Generation 

Configuration for GPT-based synthetic data looks as follows:

```yaml
connection: 
  dbType: postgres
  connection: "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
dryRun: false
provider: openai
model: gpt-3.5-turbo
schema_type: webshop # can be anything word like ecommerce, webshop, hospital etc
```

You have to set your OpenAI API key as an environment variable `OPENAI_API_KEY`.

Save above config as `config.yaml` and run the following command:
```shell
dbchaos generateWithLLM
```

This will generate the schemas, insert commands and persist it in the database.
