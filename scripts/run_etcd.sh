#!/usr/bin/env bash

# ref: https://github.com/etcd-io/etcd/releases/tag/v3.5.0

# start a local etcd server
/tmp/etcd-download-test/etcd

# write,read to etcd
#/tmp/etcd-download-test/etcdctl --endpoints=localhost:2379 put foo bar
#/tmp/etcd-download-test/etcdctl --endpoints=localhost:2379 get foo
