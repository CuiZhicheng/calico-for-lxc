# calico-for-lxc

**need root**

```
apt-get update
apt-get install docker etcd
mkdir -p /opt/bin
mkdir -p /etc/calico
mkdir -p /etc/cni/net.d
chmod a+w -R /etc/cni/net.d
cp tool/* /opt/bin/
chmod +x opt/bin/*
export PATH=$PATH:/opt/bin
```

1. create file calicoctl.cfg in /etc/calico
	* calicocfg/calicoctl.cfg.sample is an example of it, replace the ip with your etcd ip

2. create ipPool.cfg in /etc/calico if you want to set your own ip pool
	* calicocfg/ipPool.cfg.sample is an example of it, replace the cidr with your own ip pool cidr

3. create file in /etc/cni/net.d to define network
	* calicocfg/10-frontend-calico.conf is an example of it, replace the name with your own network name
	* you can create multiple .conf files to define different networks, containers in same network can communicate with each other, otherwise they can't communicate.

4. Start etcd 
```
etcd -name {NODENAME} -initial-advertise-peer-urls http://{IP}:2380 -listen-peer-urls http://0.0.0.0:2380 -listen-client-urls http://0.0.0.0:2379 -advertise-client-urls http://0.0.0.0:2379 -initial-cluster-token {CLUSTERNAME} -initial-cluster {NODENAME}=http://{IP}:2380 -initial-cluster-state new &
```
5. Configure and start docker
```
service docker stop
dockerd --cluster-store=etcd://0.0.0.0:2379 &
```

6. Run calico-node
```
calicoctl run node --name={CALICO_NODE_NAME} --ip={IP}
```
7. If you want to set your own resource, do like this:
	* calicoctl create -f /etc/calico/ipPool.cfg 

8. Operations of lxc container:
	* To create a lxc container:
```
	./lxc-create.sh {CONTAINERNAME}
```


	* To attach a container to a network:
```
	./lxc-attach-calico.sh {CONTAINERNAME} {NETWORK}
```

	* To detach a container from a network:
```
	./lxc-detach-calico.sh {CONTAINERNAME} {NETWORK}
```

	* To delete a lxc containre:
```
	./lxc-delelte.sh {CONTAINERNAME}
```
