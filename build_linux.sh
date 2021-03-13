
#!/bin/bash
os=$(go env GOOS)
arch=$(go env GOARCH)
goversion=$(go version | awk '{print $3}')
commitid=$(git rev-parse --short HEAD)
account=$(git log --pretty=format:"%%an" -1)
branch=$(git branch --show-current)
nowtime=$(date +%Y-%m-%d.%H:%M:%S)
go build -ldflags "-X github.com/yuansudong/http_dns/cmdline._GitBranch=${branch} -X github.com/yuansudong/http_dns/cmdline._OS=${os} -X github.com/yuansudong/http_dns/cmdline._Arch=${arch} -X github.com/yuansudong/http_dns/cmdline._GoVersion=${goversion} -X github.com/yuansudong/http_dns/cmdline._GitCommit=${commitid} -X github.com/yuansudong/http_dns/cmdline._GitAccount=${account} -X github.com/yuansudong/http_dns/cmdline._DateTime=${nowtime}" -o http_dns
