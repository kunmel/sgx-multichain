# sgx-multichain

* 目前分别链接了fabric 1.4.3以及ethereum私链
* 使用fabric-sdk-go连接fabric网络，使用了Shitaibin/fabric-sdk-go-sample作为基础进行修改，使其能够访问fabric并且调用链码
* 构建以太坊私链并使用go、rpc访问其后端，能够使用自带方法或调用部署的智能合约
* 目前发现go-ehthereum需要很大的内存空间，因此可能无法运行在enclave内,因此可能需要单独搭建eth、fab的客户端通过grpc进行通信。
* 20210917版本使用了grpc进行通信，其中sgx端服务创建在enclave内部，fabric、eth端服务创建在本地。sgx端通过grpc向两种区块链端发送信息，从而调用fabric-sdk-go或eth后端。

## 项目结构

sgx-multichain
├── key
├── multichain_proto
├── README.md
├── sgx-multichain-eth
├── sgx-multichain-fab
└── sgx-multichain-server

其中sgx-multichain-server为部署在sgx端的服务，其所需使用的秘钥、proto文件等均在文件夹内，只需直接复制到docker内并构建occlum。而其他key、proto等则是在本地构建eth、fab时所需要的，eth、fab两个文件夹分别包含连接两个区块链的server端，使用了go mod进行包的管理。

## 通信结构

DB ── 15432 ── sgx ──  22222 ── fabric-sdk-go

​									└─── 33333 ── eth rpc							 

## 需要注意的点

* 整个项目使用了go mod来管理依赖，其中go-ethereum与fabric-sdk-go的protobuf版本有所区别

  * go-ethereum中使用了github.com/golang/protobuf v1.4.2    google.golang.org/protobuf v1.23.0
  * fabric-sdk-go中使用了github.com/golang/protobuf v1.3.3    google.golang.org/protobuf v1.23.0
  * 这会导致后续放在一起编译时由于版本不同、结构不同导致报错，所以在linker的go.mod添加`github.com/golang/protobuf => github.com/golang/protobuf v1.3.3`
  
* 转到docker内要做的操作：
  * eth的秘钥文件、ethereum/contract 复制并修改地址
  * fabric-samples下的文件复制并修改地址
  * 所有ip都要改为真实ip，不能使用localhost
  * 别忘了打开111module以及设置goproxy

