#!/bin/bash
name=$1
network=$2
sudo lxc-create -t ubuntu -n $name -- -r trusty

sudo tee /var/lib/lxc/$name/config <<EOF
# Common configuration
lxc.include = /usr/share/lxc/config/ubuntu.common.conf
# Container specific configuration
lxc.rootfs = /var/lib/lxc/$name/rootfs
lxc.rootfs.backend = dir
lxc.utsname = $name
lxc.arch = amd64
# Network configuration
lxc.network.type = empty
lxc.network.flags = up
EOF

sudo lxc-start -n $name -d
pid=$(sudo lxc-info -n $name | grep 'PID' | cut -d ':' -f 2 |  tr -d '[[:space:]]')
sudo mkdir -p /var/run/netns
sudo ln -s /proc/$pid/ns/net /var/run/netns/$pid

