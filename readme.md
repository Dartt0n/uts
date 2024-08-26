# uts

Simple [Unix timestamp](https://en.wikipedia.org/wiki/Unix_time) to human readable date converter. Convert unix timestamp in seconds or nanoseconds to [RFC1123](https://datatracker.ietf.org/doc/html/rfc1123) format.

## Install
```
go install github.com/dartt0n/uts@latest
```

## Usage
```
Usage: uts <unix timestamp>
```

### seconds precision
```bash
$ uts 1724692825
> Mon, 26 Aug 2024 20:20:25 UTC
```

### nanoseconds precision
```bash
$ uts 1723140436809000000
> Thu, 08 Aug 2024 21:07:16 UTC
```

### pipe from stdin
```bash
$ echo 1724692825 | uts
> Mon, 26 Aug 2024 20:20:25 UTC
```

## Development

### Tools
- [Go](https://go.dev/)
- [Just](https://github.com/casey/just)

### Run
```bash
just run
```
