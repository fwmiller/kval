# kval: A key-value store

kval is a persistent key-value store written in Go.  The system keeps
key-value stores (hereafter referred to as databases) on secondary storage
(e.g. rotating disk or solid state storage) on a local machine.

The system is architected with a front-end command line interface (cli)
and a backend "library" that can be decoupled from the cli.  The commands
try to follow Redis syntax where possible.


## cli Commands

The cli is organized as a sub-package to the kval main program.  A set of
cli commands are implemented as functions in this package.  Each function
has a corresponding backend function defined by the kval interface (see
the Backend API below).  The following table lists the implemented cli
commands.


| Command	| Description					|
|---------------|-----------------------------------------------|
| select	| Set current database				|
| create	| Create a new key-value store database		|
| remove	| Remove existing database			|
| keys          | List the keys in a database                   |
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

### keys

Syntax: `keys`

### get

Syntax: `get` _key_

### del

Syntax: `del` _key_


## Backend API

The backend is organized as a sub-package and is used by the cli sub-package.
The kval package exports an interface that represents the operations that can
be performed on the persistent key-value store.  

### `kval.DB.IsDb(dbname string) string`

Check whether `dbname` is a valid database name that currently exists

### `kval.DB.CreateDb(dbname string) bool`

Create a new key-value store database called `dbname`

### `kval.DB.RemoveDb(dbname string) bool`

Remove database called `dbname`

### `kval.DB.Keys(dbname string)`

Print the keys in the database called `dbname`

### `kval.DB.Set(dbname string, key string, value string)`

Set a new `key`-`value` pair in the database `dbname`

### `kval.DB.Get(dbname string, key string) string`

Return the value associated with the specified `key` in the database `dbname`

### `kval.DB.Del(dbname string, key string)`

Delete the `key`-value pair in the database `dbname`
