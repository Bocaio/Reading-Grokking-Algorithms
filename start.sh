set -e 
(
    cd "$(dirname "$0")" 
    go build -o /tmp/main main.go
)
exec /tmp/main "$@"