# Repository

The repository package exposes various ways to interact with storage backends.  In particular, the intent is to expose implementations for the API's Repository interface, but the functionality exists independent of the API itself.

## Postgres
Interact with a Postgres database by issuing various `SELECT` and `INSERT` commands.  Useful for persistence of data - production and local testing.

## Memory
Expose the same interface, but store all data in memory.  There's no persistence across application restarts, but start each session with a familiar set of hardcoded data points.
