# Cacherus 
##### Extremely simple LRU cache

## Description
Uses the command line to read in a two commands: GET & PUT

- PUT: Adds a key value pair to the cache
- GET: Retrieves a value using the key

The cache has a mix size of 3 that's set in the source code. Once this limit is exceeded the last accessed/used key value pair will be evicted.
Concurrent access to cache has not been optimized but was considered.
