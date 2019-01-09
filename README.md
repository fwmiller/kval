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


| Command	| Description						|
|---------------|-------------------------------------------------------|
| select	| Set current database					|
| create	| Create a new key-value store database			|
| remove	| Remove existing database				|
| keys          | List the keys in a database				|
| set		| Set a key-value pair					|
| get		| Get the value associated with a key			|
| del		| Delete a key-value pair				|
| list		| List the databases for this server			|
| time		| Get current server time				|
| exists	| Determine if a key exists				|
| TODO: copy	| Copy a key-value pair to another database		|
| TODO: append	| Append a value to a key				|
| TODO: dbsize	| Get the number of keys in a database			|
| TODO: dump	| Return a serialized version of a key-value pair	|
| TODO: expire	| Set a key's time-to-live as a duration		|
| TODO: expireat| Set a key's time-to-live as an absolute time		|
| TODO: flushall| Remove all keys from all databases			|
| TODO: flushdb	| Remove all keys from current database			|
| TODO: getset	| Set the value of a key and return its old value	|
| TODO: mget	| Get the values for a list of keys			|
| TODO: move	| Move a key-value pair to another database		|
| TODO: pttl	| Get the time-to-live for a key in milliseconds	|
| TODO: rename	| Rename a key						|
| TODO: renamenx| Rename a key if the new key does not already exist	|
| TODO: restore	| Create a key-value pair using serialized version	|
| TODO: setnx	| Set key-value pair if the key does not already exist	|
| TODO: strlen	| Get the length of the value associated with a key	|
| quit		| Quit							|
| help		| Help							|

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

### list

Syntax: `list`

### time

Syntax: `time`

### exists

Syntax: `exists` _key_


## TODO: Allow for multiple simultaneous network client connections


## Backend API

The backend is organized as a sub-package and is used by the cli sub-package.
The kval package exports an interface that represents the operations that can
be performed on the persistent key-value store.  

### `IsDb(dbname string) (string, error)`

Check whether `dbname` is a valid database name that currently exists

### `CreateDb(dbname string) error`

Create a new key-value store database called `dbname`

### `RemoveDb(dbname string) error`

Remove database called `dbname`

### `Keys(dbname string) ([]string, error)`

Generate a slice of strings containing the keys in the database called
`dbname`

### `Set(dbname string, key string, value string) error`

Set a new `key`-`value` pair in the database `dbname`

### `Get(dbname string, key string) (string, error)`

Return the value associated with the specified `key` in the database `dbname`

### `Del(dbname string, key string) error`

Delete the `key`-value pair in the database `dbname`

### `List() ([]string, error)`

Generate a slice of strings containing the names of the databases on this
server

### `Time() string`

Return a string containing the current time at the server

### `Exists(dbname string, key string) bool`

Return `true` if the key exists in the database `dbname` and `false` otherwise
