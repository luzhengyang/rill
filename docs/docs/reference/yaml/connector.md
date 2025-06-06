---
note: GENERATED. DO NOT EDIT.
title: Connector YAML
sidebar_position: 35
---

When you add olap_connector to your rill.yaml file, you will need to set up a `<connector_name>.yaml` file in the 'connectors' directory. This file requires the following parameters,type and driver (see below for more parameter options). Rill will automatically test the connectivity to the OLAP engine upon saving the file. This can be viewed in the connectors tab in the UI.

:::tip Did you know?

Starting from Rill 0.46, you can directly create OLAP engines from the UI! Select + Add -> Data -> Connect an OLAP engine

:::


## Properties

### `type`

_[string]_ - Refers to the resource type and must be `connector` _(required)_

## Common Properties

### `name`

_[string]_ - Name is usually inferred from the filename, but can be specified manually. 

### `refs`

_[array of string]_ - List of resource references 

### `dev`

_[object]_ - Overrides any properties in development environment. 

### `prod`

_[object]_ - Overrides any properties in production environment. 

## One of Properties Options
- [athena](#athena)
- [azure](#azure)
- [bigquery](#bigquery)
- [clickhouse](#clickhouse)
- [druid](#druid)
- [duckdb](#duckdb)
- [gcs](#gcs)
- [https](#https)
- [local_file](#local_file)
- [motherduck](#motherduck)
- [mysql](#mysql)
- [pinot](#pinot)
- [postgres](#postgres)
- [redshift](#redshift)
- [s3](#s3)
- [salesforce](#salesforce)
- [slack](#slack)
- [snowflake](#snowflake)
- [sqlite](#sqlite)

## athena

Configuration properties specific to the athena

### `driver`

_[string]_ - Refers to the driver type and must be driver `athena` _(required)_

### `aws_access_key_id`

_[string]_ - AWS Access Key ID for Athena access 

### `aws_secret_access_key`

_[string]_ - AWS Secret Access Key for Athena access 

### `aws_access_token`

_[string]_ - Optional AWS session token for temporary credentials 

### `external_id`

_[string]_ - Optional External ID for assuming a role 

### `role_arn`

_[string]_ - Optional AWS Role ARN to assume when accessing Athena 

### `role_session_name`

_[string]_ - Optional Session name when assuming the role 

### `allow_host_access`

_[boolean]_ - Allow access to host environment configuration 

## azure

Configuration properties specific to the azure

### `driver`

_[string]_ - Refers to the driver type and must be driver `azure` _(required)_

### `azure_storage_account`

_[string]_ - Azure storage account name 

### `azure_storage_key`

_[string]_ - Azure storage access key 

### `azure_storage_sas_token`

_[string]_ - Optional azure SAS token for authentication 

### `azure_storage_connection_string`

_[string]_ - Optional azure connection string for storage account 

### `allow_host_access`

_[boolean]_ - Allow access to host environment configuration 

## bigquery

Configuration properties specific to the bigquery

### `driver`

_[string]_ - Refers to the driver type and must be driver `bigquery` _(required)_

### `google_application_credentials`

_[string]_ - Path to the Google Cloud credentials JSON file 

### `allow_host_access`

_[boolean]_ - Allow access to host environment configuration 

## clickhouse

Configuration properties specific to the clickhouse

### `driver`

_[string]_ - Refers to the driver type and must be driver `clickhouse` _(required)_

### `managed`

_[boolean]_ - `true` means Rill will provision the connector using the default provisioner. `false` disables automatic provisioning. 

### `dsn`

_[string]_ - DSN(Data Source Name) for the ClickHouse connection 

### `username`

_[string]_ - Username for authentication 

### `password`

_[string]_ - Password for authentication 

### `host`

_[string]_ - Host where the ClickHouse instance is running 

### `port`

_[integer]_ - Port where the ClickHouse instance is accessible 

### `database`

_[string]_ - Name of the ClickHouse database within the cluster 

### `ssl`

_[boolean]_ - Indicates whether a secured SSL connection is required 

### `cluster`

_[string]_ - Cluster name, required for running distributed queries 

### `log_queries`

_[boolean]_ - Controls whether to log raw SQL queries 

### `settings_override`

_[string]_ - override the default settings used in queries. example `readonly = 1, session_timezone = 'UTC'` 

### `embed_port`

_[integer]_ - Port to run ClickHouse locally (0 for random port) 

### `can_scale_to_zero`

_[boolean]_ - Indicates if the database can scale to zero 

### `max_open_conns`

_[integer]_ - Maximum number of open connections to the database 

### `max_idle_conns`

_[integer]_ - Maximum number of idle connections in the pool 

### `dial_timeout`

_[string]_ - Timeout for dialing the ClickHouse server 

### `conn_max_lifetime`

_[string]_ - Maximum time a connection may be reused 

### `read_timeout`

_[string]_ - Maximum time for a connection to read data 

## druid

Configuration properties specific to the druid

### `driver`

_[string]_ - Refers to the driver type and must be driver `druid` _(required)_

### `dsn`

_[string]_ - Data Source Name (DSN) for connecting to Druid _(required)_

### `username`

_[string]_ - Username for authenticating with Druid 

### `password`

_[string]_ - Password for authenticating with Druid 

### `host`

_[string]_ - Hostname of the Druid coordinator or broker 

### `port`

_[integer]_ - Port number of the Druid service 

### `ssl`

_[boolean]_ - Enable SSL for secure connection 

### `log_queries`

_[boolean]_ - Log raw SQL queries sent to Druid 

### `max_open_conns`

_[integer]_ - Maximum number of open database connections (0 = default, -1 = unlimited) 

### `skip_version_check`

_[boolean]_ - Skip checking Druid version compatibility 

## duckdb

Configuration properties specific to the duckdb

### `driver`

_[string]_ - Refers to the driver type and must be driver `duckdb` _(required)_

### `pool_size`

_[integer]_ - Number of concurrent connections and queries allowed 

### `allow_host_access`

_[boolean]_ - Whether access to the local environment and file system is allowed 

### `cpu`

_[integer]_ - Number of CPU cores available to the database 

### `memory_limit_gb`

_[integer]_ - Amount of memory in GB available to the database 

### `read_write_ratio`

_[number]_ - Ratio of resources allocated to the read database; used to divide CPU and memory 

### `boot_queries`

_[string]_ - SQL to run when initializing a new connection, before extensions and defaults 

### `init_sql`

_[string]_ - SQL to run when initializing a new connection, after extensions and defaults 

### `log_queries`

_[boolean]_ - Whether to log raw SQL queries executed through OLAP 

## gcs

Configuration properties specific to the gcs

### `driver`

_[string]_ - Refers to the driver type and must be driver `gcs` _(required)_

### `google_application_credentials`

_[string]_ - Google Cloud credentials JSON string 

### `allow_host_access`

_[boolean]_ - Allow access to host environment configuration 

### `key_id`

_[string]_ - Optional S3-compatible Key ID when used in compatibility mode 

### `secret`

_[string]_ - Optional S3-compatible Secret when used in compatibility mode 

## https

Configuration properties specific to the https

### `driver`

_[string]_ - Refers to the driver type and must be driver `https` _(required)_

### `path`

_[string]_ - The full HTTPS URI to fetch data from _(required)_

### `headers`

_[object]_ - HTTP headers to include in the request 

## local_file

Configuration properties specific to the local_file

### `driver`

_[string]_ - Refers to the driver type and must be driver `local_file` _(required)_

### `dsn`

_[string]_ - Data Source Name (DSN) indicating the file path or location of the local file _(required)_

### `allow_host_access`

_[boolean]_ - Flag to indicate if access to host-level file paths is permitted 

## motherduck

Configuration properties specific to the motherduck

### `driver`

_[string]_ - Refers to the driver type and must be driver `motherduck` _(required)_

### `dsn`

_[string]_ - Data Source Name (DSN) specifying the MotherDuck connection endpoint _(required)_

### `token`

_[string]_ - Authentication token for accessing MotherDuck (secret) _(required)_

## mysql

Configuration properties specific to the mysql

### `driver`

_[string]_ - Refers to the driver type and must be driver `mysql` _(required)_

### `dsn`

_[string]_ - DSN(Data Source Name) for the mysql connection _(required)_

## pinot

Configuration properties specific to the pinot

### `driver`

_[string]_ - Refers to the driver type and must be driver `pinot` _(required)_

### `dsn`

_[string]_ - DSN(Data Source Name) for the Pinot connection _(required)_

### `username`

_[string]_ - Username for authenticating with Pinot 

### `password`

_[string]_ - Password for authenticating with Pinot 

### `broker_host`

_[string]_ - Hostname of the Pinot broker _(required)_

### `broker_port`

_[integer]_ - Port number for the Pinot broker 

### `controller_host`

_[string]_ - Hostname of the Pinot controller _(required)_

### `controller_port`

_[integer]_ - Port number for the Pinot controller 

### `ssl`

_[boolean]_ - Enable SSL connection to Pinot 

### `log_queries`

_[boolean]_ - Log raw SQL queries executed through Pinot 

### `max_open_conns`

_[integer]_ - Maximum number of open connections to the Pinot database 

## postgres

Configuration properties specific to the postgres

### `driver`

_[string]_ - Refers to the driver type and must be driver `postgres` _(required)_

### `dsn`

_[string]_ - DSN(Data Source Name) for the postgres connection _(required)_

## redshift

Configuration properties specific to the redshift

### `driver`

_[string]_ - Refers to the driver type and must be driver `redshift` _(required)_

### `aws_access_key_id`

_[string]_ - AWS access key ID for authentication 

### `aws_secret_access_key`

_[string]_ - AWS secret access key for authentication 

### `aws_access_token`

_[string]_ - AWS session token for temporary credentials (optional) 

### `allow_host_access`

_[boolean]_ - Allow access to host environment configuration 

## s3

Configuration properties specific to the s3

### `driver`

_[string]_ - Refers to the driver type and must be driver `s3` _(required)_

### `aws_access_key_id`

_[string]_ - AWS Access Key ID used for authentication 

### `aws_secret_access_key`

_[string]_ - AWS Secret Access Key used for authentication 

### `aws_access_token`

_[string]_ - Optional AWS session token for temporary credentials 

### `endpoint`

_[string]_ - Optional custom endpoint URL for S3-compatible storage 

### `region`

_[string]_ - AWS region of the S3 bucket 

### `allow_host_access`

_[boolean]_ - Allow access to host environment configuration 

### `retain_files`

_[boolean]_ - Whether to retain intermediate files after processing 

## salesforce

Configuration properties specific to the salesforce

### `driver`

_[string]_ - Refers to the driver type and must be driver `salesforce` _(required)_

### `username`

_[string]_ - Salesforce account username _(required)_

### `password`

_[string]_ - Salesforce account password (secret) 

### `key`

_[string]_ - Authentication key for Salesforce (secret) 

### `endpoint`

_[string]_ - Salesforce API endpoint URL _(required)_

### `client_id`

_[string]_ - Client ID used for Salesforce OAuth authentication 

## slack

Configuration properties specific to the slack

### `driver`

_[string]_ - Refers to the driver type and must be driver `slack` _(required)_

### `bot_token`

_[string]_ - Bot token used for authenticating Slack API requests _(required)_

## snowflake

Configuration properties specific to the snowflake

### `driver`

_[string]_ - Refers to the driver type and must be driver `snowflake` _(required)_

### `dsn`

_[string]_ - DSN (Data Source Name) for the Snowflake connection _(required)_

### `parallel_fetch_limit`

_[integer]_ - Maximum number of concurrent fetches during query execution 

## sqlite

Configuration properties specific to the sqlite

### `driver`

_[string]_ - Refers to the driver type and must be driver `sqlite` _(required)_

### `dsn`

_[string]_ - DSN(Data Source Name) for the sqlite connection _(required)_