#!/usr/bin/Env bash
osArray=("linux" "darwin" "freebsd" "windows")
archs=("amd64" "386")
version=${1-"0.0.1-preview1"}
out_file="yuanbootctl"

build() {
  os=$1
  arch=$2
  go env -w GOPROXY=https://goproxy.cn,direct
  GOOS=$os GOARCH=$arch go build -o "${out_file}" .
  tgzName="${out_file}_${version}_${os}_${arch}.tar.gz"
  rm -f "${tgzName}"
  tar -czf "${tgzName}" "${out_file}"
  rm -f "${out_file}"

  echo $(shasum -a 256 "${tgzName}")
}

main() {
  cd cli/yuanbootctl
  ls
  echo "mod download"
  go get -t .
  go mod download
  for os in "${osArray[@]}"; do
    for arch in "${archs[@]}"; do
      if [ os=='darwin' && arch=='386' ]; then
          continue
      fi
      build "${os}" "${arch}"
    done
  done
}

main