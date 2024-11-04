all: install

install:
    go install ./cmd/mychaind
    go install ./cmd/mychaindcli

build:
    go build -o build/mychaind ./cmd/mychaind
    go build -o build/mychaindcli ./cmd/mychaindcli

clean:
    rm -rf build/
