# Maintainer: Vivek Kumar Sahu <vivekkumarsahu650@gmail.com>
pkgname=url_shortner
pkgver=1.0.0
pkgrel=0
pkgdesc="A simple URL shortener written in Go"
url="https://github.com/viveksahu26/url_shortner"
arch="all"
license="Apache-2.0"
makedepends="go"
source="${pkgname}-${pkgver}.tar.gz::https://github.com/viveksahu26/url_shortner/archive/refs/tags/v${pkgver}.tar.gz"
builddir="$srcdir/${pkgname}-${pkgver}"
export GOCACHE="${GOCACHE:-"$srcdir/go-cache"}"
export GOTMPDIR="${GOTMPDIR:-"$srcdir"}"
export GOMODCACHE="${GOMODCACHE:-"$srcdir/go"}"
export GO111MODULE=auto
export GOPATH="$srcdir"

prepare() {
  default_prepare
  mkdir -p "${builddir}/src"
  mv "${builddir}"/* "${builddir}/src/"
}

build() {
  cd "$builddir/src"
  go build -o ../url_shortner .
}

package() {
  install -Dm755 "$builddir/url_shortner" "$pkgdir"/usr/bin/url_shortner
}

sha512sums="" # You will need to generate this using 'abuild checksum'
