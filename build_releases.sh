#!/usr/bin/env bash

echo 'Generating binary for linux...'
gox -osarch="linux/amd64"
mv pinger_linux_amd64 pinger
tar -czf pinger-$1-x86_64-linux.tar.gz pinger

echo 'Generating binary for darwin...'
gox -osarch="darwin/amd64"
mv pinger_darwin_amd64 pinger
tar -czf pinger-$1-x86_64-darwin.tar.gz pinger

echo 'Generating binary for windows...'
gox -osarch="windows/amd64"
mv pinger_windows_amd64.exe pinger.exe
zip pinger-$1-x86_64-windows.zip pinger.exe

echo 'Cleaning up...'
rm pinger pinger.exe
