# Maintainer: Vivek Kumar Sahu <vivekkumarsahu650@gmail.com>
pkgname=url_shortner
pkgver=0.2.0
pkgrel=0
pkgdesc="A simple URL shortener written in Go"
url="https://github.com/viveksahu26/url_shortner"
arch="all"
license="Apache-2.0"
makedepends="go tar"
source="https://github.com/viveksahu26/url_shortner/archive/v0.2.0/url_shortner-0.2.0.tar.gz"
builddir="$srcdir"


export GOCACHE="${GOCACHE:-"$srcdir/go-cache"}"
export GOTMPDIR="${GOTMPDIR:-"$srcdir"}"
export GOMODCACHE="${GOMODCACHE:-"$srcdir/go"}"


build() {
  echo "++ Buider: " $builddir
  echo "++ GOCACHE: " $GOCACHE
  echo "++ GOTMPDIR: " $GOTMPDIR
  echo "++ GOMODCACHE: " $GOMODCACHE
  cd $builddir
  echo "Go Version: " $(go version)
  tar -xvzf url_shortner-0.2.0.tar.gz
  cd url_shortner-0.2.0
  go mod tidy
  go build -o url_shortner  ./cmd/url_shortner 
}

package() {
  install -Dm755 "$builddir/url_shortner" "$pkgdir"/usr/bin/url_shortner
}

sha512sums="087fb6a6766bb0de4afbd5c6ef64542606255403792d02a56577737c049c745941ba6989b902700b2b41c9cd0a4eed8e3b5366ff132eeba1f7e56e6bdf8d6fea  url_shortner-0.2.0.tar.gz"
