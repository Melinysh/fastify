# Fastify
### Download Acceleration as a Service

Everyone gets tired waiting for their large downloads to complete. 
BitTorrent is awesome, but you may not have a bunch of peers ready to seed it.
Fastify, a download accelerator as a service, solves both these problems and regularly enables 4x download speeds. 
In one case, Fastify delivered a file 24x faster than a conventional download. 

**Fastify won [Best Use Of AWS](https://devpost.com/software/fastify) at QHacks!**

### How do we do it?

Fastify has a thin web frontend that send your requested link to our backend.
This backend is actually a high performance EC2 cluster.
Each server in the backend receives the users link and begin to download the file locally. Once the file download is complete we generate a `.torrent` for the downloaded file and begin to seed it.
We send this `.torrent` file back to the user.
The user can open this `.torrent` file in thier favourite Bittorrent client and enjoy a much faster download. 

We even cache some downloads so popular downloads will be able to be pulled from Fastify even speedier!

Without any cache hits, we saw the following improvements in download speeds with our test files:

```
|                   | 512Mb    | 1Gb    | 2Gb     | 5Gb     |
|-------------------|----------|--------|---------|---------|
| Regular Download  | 3 mins   | 7 mins | 13 mins | 30 mins |
| Fastify           | 1.5 mins | 3 mins | 5 mins  | 9 mins  |
|-------------------|----------|--------|---------|---------|
| Effective Speedup | 2x       | 2.33x  | 2.6x    | 3.3x    |
```
_test was performed with slices of the ubuntu 16.04 iso file, on the eduroam network_
