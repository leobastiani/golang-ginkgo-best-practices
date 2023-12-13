# ginkgo example

```bash
GOFLAGS=-mod=mod GOBIN=$(git rev-parse --show-toplevel)/bin go install github.com/onsi/ginkgo/v2/ginkgo
export PATH="$PWD/bin:$PATH"
go test -v ./pkg/...
ginkgo -r
```
