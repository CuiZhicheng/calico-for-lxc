#!/bin/bash

apt-get update
apt-get -y install docker docker.io lxc etcd aufs-tools conntrack
mkdir -p /opt/bin
mkdir -p /etc/calico
mkdir -p /etc/cni/net.d
chmod a+w -R /etc/cni/net.d
cp tool/* /opt/bin/
chmod +x /opt/bin/calico
chmod +x /opt/bin/calico-ipam
chmod +x /opt/bin/calicoctl
chmod +x /opt/bin/cnitool
export PATH=$PATH:/opt/bin
