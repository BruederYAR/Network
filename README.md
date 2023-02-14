# Peer-To-Peer Network
It is a peer-to-peer network providing a secure protocol for data transmission. Data is encrypted using RSA

The network is created when node is created. Then the nodes are united into one network

## Create node
You need to choose how you will interact with the network. 
```
# Console
go run ./cmd/console/console.go [ip]:[port] [node_name]

# Http Api
go run ./cmd/api/api.go [ip]:[port] [node_name]
```
*If the sqlite database cannot be created, then you need to download the GCC compiler*

Ip can be skipped, then it will be automatically inserted as ipv4. Instead of ip, you can also write: ipv4, ipv6

When creating a node, 3 files are created:
- log_[node_name].txt - Node logs. Here you can view who and when did the handshakes
- [node_name].db - sqlite database that stores information about nodes on the network
- privatekey - Private key for RSA encryption

## Interaction with node

### Connect
Allows you to combine networks into one \ Join the network

In the console UI, write the command:
```
/connect [ip]:[port]
```
In the http API, make a GET request:
```
http://localhost/api/control/connect/:address
```

### Nodes info
Helps to find out which nodes are present in the network

In the console UI, write the command:
```
/network
```
In the http API, make a GET request:
```
http://localhost/api/nodes/
```

### Send
Allows you to send data to another node

In the console UI, write the command:
```
/m [ip]:[port] [message]
# or
/m [node_name] [message]
```
In the http API, make a POST request:
```
http://localhost/control/send

# body
{
    "address":"[ip]:[port]",
    "message":"message"
}
```