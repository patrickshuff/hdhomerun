# hdhomerun cli tools
A commandline tool written to provide an alternative open source equivalent to the hdhomerun_config tool provided by SiliconDust for interacting with their tv tuners.

# Examples
## Discover where your HDHR tuner is on your network
```
$ go run hdhomerun.go discover       
hdhomerun device 1322f2f9 found at 192.168.174.249
```
## Get a listing of all channels
```
$ go run hdhomerun.go channels | head      
hdhomerun device 1322f2f9 found at 192.168.174.249
  2     KTVU                    http://192.168.174.249:5004/auto/v2
  3     KNTV                    http://192.168.174.249:5004/auto/v3
  4     KRON                    http://192.168.174.249:5004/auto/v4
  5     KPIX                    http://192.168.174.249:5004/auto/v5
  6     KICU                    http://192.168.174.249:5004/auto/v6
  7     KGO                     http://192.168.174.249:5004/auto/v7
  8     KTSF                    http://192.168.174.249:5004/auto/v8
  9     KQED                    http://192.168.174.249:5004/auto/v9
 10     KTEH                    http://192.168.174.249:5004/auto/v10
 ```
 
 # Getting started
```

# Clone the repo
git clone https://github.com/patrickshuff/hdhomerun.git

# cd into the dir
cd hdhomerun

# Build it/grab the dependencies
go build hdhomerun.go

# Run it
go run hdhomerun.go discover

```
