# kval: A key-value store

kval is a persistent key-value store written in Go.  The system keeps
key-value stores (hereafter referred to as databases) on secondary storage
(e.g. rotating disk or solid state storage) on a local machine.

The system is architected with a front-end command line interface (cli)
and a backend "library" that can be decoupled from the cli.

## cli Commands

| Command	| Description					|
|---------------|-----------------------------------------------|
| db		| Set current database				|
| create	| Create a new key-value store database		|
| remove	| Remove existing database			|
| set		| Set a key-value pair				|
| get		| Get the value associated with a key		|
| del		| Delete a key-value pair			|
| quit		| Quit						|
| help		| Help						|

## Backend API

### `KvalInit()`

Initialize the kval backend

### `KvalIsDb(dbname string) string`

Check whether `dbname` is a valid database name that currently exists

### `KvalCreateDb(dbname string)`

Create a new key-value store database called `dbname`

### `KvalSet(dbname string, key string, value string)`

Set a new `key`-`value` pair in the database `dbname`
