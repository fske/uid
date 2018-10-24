# uid

ID is composed of

x bytes for prefix
8 bytes for time in units of 10 ms
2 bytes for a instance id, not greater than 1000
2 bytes for a loop number, not greater than 1000

# Installation

```
go get github.com/fske/uid
```

# Usage

```
idGen, err := NewUIDGenerator("foo", int32(1))
id := idGen.ID()
```
