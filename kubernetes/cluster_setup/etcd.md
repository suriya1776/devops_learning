# ETCD

Etcd is a distributed, consistent key-value store that provides a reliable way to store data across a cluster of machines. It's used primarily as a backend storage system for distributed systems, providing configuration management, service discovery, and distributed coordination

# Installing and Running Etcd

Considering the ec2 machine or any machine is already provisioned

* Update the system packages:
```
sudo yum update -y
```
* Install necessary packages
```
sudo yum install -y wget tar
```
* Download and Install etcd
```
ETCD_VERSION=v3.5.0
wget https://github.com/etcd-io/etcd/releases/download/${ETCD_VERSION}/etcd-${ETCD_VERSION}-linux-amd64.tar.gz
```
* Extract the tarball
```
tar -xvf etcd-${ETCD_VERSION}-linux-amd64.tar.gz
```
* Move etcd and etcdctl binaries to /usr/local/bin
```
sudo mv etcd-${ETCD_VERSION}-linux-amd64/etcd /usr/local/bin/
sudo mv etcd-${ETCD_VERSION}-linux-amd64/etcdctl /usr/local/bin/
```
* Verify the installation
```
etcd --version
etcdctl version
```
* Create a directory for etcd data
```
sudo mkdir -p /var/lib/etcd
```
* Create a systemd service file for etcd
```
sudo tee /etc/systemd/system/etcd.service <<EOF
[Unit]
Description=etcd key-value store
Documentation=https://github.com/etcd-io/etcd
After=network.target

[Service]
User=root
Type=notify
ExecStart=/usr/local/bin/etcd \\
  --name etcd-server \\
  --data-dir /var/lib/etcd \\
  --listen-client-urls http://0.0.0.0:2379 \\
  --advertise-client-urls http://0.0.0.0:2379
Restart=always
RestartSec=5
LimitNOFILE=40000

[Install]
WantedBy=multi-user.target
EOF
```

* Reload the systemd daemon
```
sudo systemctl daemon-reload
```
* Enable and start the etcd service
```
sudo systemctl enable etcd
sudo systemctl start etcd
```
* Check the status of the etcd service
```
sudo systemctl status etcd
```
* Verify etcd is running
```
etcdctl --endpoints=http://localhost:2379 endpoint status
```
* Put and get a key-value pair to ensure etcd is working
```
etcdctl --endpoints=http://localhost:2379 put key mykey
etcdctl --endpoints=http://localhost:2379 get key
```
Response will be like
```
key
mykey
```
