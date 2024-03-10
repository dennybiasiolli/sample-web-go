# Install Go

https://go.dev/doc/install

```bash
GO_VERSION=1.22.1
GO_OS=linux
GO_ARCH=arm64

wget https://golang.org/dl/go$GO_VERSION.$GO_OS-$GO_ARCH.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go$GO_VERSION.$GO_OS-$GO_ARCH.tar.gz
rm go$GO_VERSION.$GO_OS-$GO_ARCH.tar.gz

export PATH=$PATH:/usr/local/go/bin

cat >>$HOME/.profile <<EOT

# go
export PATH=$PATH:/usr/local/go/bin
EOT
cat >>$HOME/.bashrc <<EOT

# go
export PATH=$PATH:/usr/local/go/bin
EOT
```

## Update to a new version of Go

Relaunch the installation process with the new version number.
