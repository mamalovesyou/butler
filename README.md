HeyButler.io

# Bootstrap
Some tools you need to install before jumping in

- Golang 1.17
- pre-commit

## Quickstart

### Requirements
You need to install some tools and configure them before setting up your environment.

1. Install `brew`: Follow instructions at [brew.sh](https://brew.sh/)
2. Install `docker`: Follow instructions on [Docker website](https://docs.docker.com/engine/install/) 
3. Install third party tools
    ```shell 
   $ make install-tools
    ```

# Getting Started

## Create migrations
To create a new SQL migration for a service just use the following command. All migrations for a service are stored in the `migrations` directory at the service root directory:
```
migrate create -ext sql -dir internal/services/<service_name>/migrations <migration_name>
```
E.g: Let's create table called `users` for the `auth` service:
```
migrate create -ext sql -dir internal/services/auth/migrations create_users_table
```
If there were no errors, we should have two files available under `internal/services/auth/migrations` folder:
- 20220718140058_create_users_table.down.sql
- 20220718140058_create_users_table.up.sql

Note the `sql` extension that we provided.

Best practice: When writing a rollback migration in the `.down.sql` add `IF EXISTS/IF NOT EXISTS`. By adding `IF EXISTS/IF NOT EXISTS` we are making migrations idempotent. E.g:
```
DROP TABLE IF EXISTS users;
```
