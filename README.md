# sgx-multichain

* 目前分别链接了fabric 1.4.3以及ethereum私链
* 使用fabric-sdk-go连接fabric网络，使用了Shitaibin/fabric-sdk-go-sample作为基础进行修改，使其能够访问fabric并且调用链码
* 构建以太坊私链并使用go、rpc访问其后端，能够使用自带方法或调用部署的智能合约

## 需要注意的点

* 整个项目使用了go mod来管理依赖，其中go-ethereum与fabric-sdk-go的protobuf版本有所区别

  * go-ethereum中使用了github.com/golang/protobuf v1.4.2    google.golang.org/protobuf v1.23.0
  * fabric-sdk-go中使用了github.com/golang/protobuf v1.3.3    google.golang.org/protobuf v1.23.0
  * 这会导致后续放在一起编译时由于版本不同、结构不同导致报错，所以在linker的go.mod添加`github.com/golang/protobuf => github.com/golang/protobuf v1.3.3`

  

