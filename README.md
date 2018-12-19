# kval: A key-value store

kval is a persistent key-value store written in Go.  The system keeps
key-value stores (hereafter referred to as databases) on secondary storage
(e.g. rotating disk or solid state storage) on a local machine.

The system is architected with a front-end command line interface (cli)
and a backend "library" that can be decoupled from the cli.  The commands
try to follow Redis syntax where possible.


## cli Commands

| Command	| Description					|
|---------------|-----------------------------------------------|
| select	| Set current database				|
| create	| Create a new key-value store database		|
| remove	| Remove existing database			|
| set		| Set a key-value pair				|
| get		| Get the value associated with a key		|
| del		| Delete a key-value pair			|
| quit		| Quit						|
| help		| Help						|

### select

Syntax: `select` _dbname_

### create

Syntax: `create` _dbname_

### remove

Syntax: `remove` _dbname_

### set

Syntax: `set` _key_ _value_

### get

Syntax: `get` _key_

### del

Syntax: `del` _key_


## Backend API

### `KvalInit()`

Initialize the kval backend

### `KvalIsDb(dbname string) string`

Check whether `dbname` is a valid database name that currently exists

### `KvalCreateDb(dbname string) bool`

Create a new key-value store database called `dbname`

### `KvalRemoveDb(dbname string) bool`

Remove database called `dbname`

### `KvalSet(dbname string, key string, value string)`

Set a new `key`-`value` pair in the database `dbname`

### `KvalGet(dbname string, key string) string`

Return the value associated with the specified `key` in the database `dbname`
