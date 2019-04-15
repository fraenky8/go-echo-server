# go-echo-server

a simple echo server written in go to play around with (Docker, AWS EC2, ...)

## Notes

#### ssh:

`ssh -i ~/.ssh/fme_aws_dev.pem ec2-user@<IP_ADDRESS>`

#### scp:
* copy the build for linux/x64 to the EC2 instance

`scp ~/Coding/Go/src/github.com/fraenky8/go-echo-server/go-echo-server ec2-user@<IP_ADDRESS>:/home/ec2-user`


#### Apache Benchmark
* 10 clients with 100 requests each

`ab -c 10 -n 100 http://<IP_ADDRESS>/ab`


#### Linux general

```
sudo ./go-echo-server > out.log 2>&1 &

tail -f ./out.log
```
