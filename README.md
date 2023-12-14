# Golang学习代码

## 规范
- 命名
  - 文件名：必须小写，使用_做单词间隔
  - 变量、函数：
    - 驼峰规范，例如GetIdGroup/GetUrlGroup，不要写GetIDGroup/GetURLGroup（提升单词间区分度，避免歧义）
    - 词组间有层级关系的使用_间隔，例如枚举中：RoomType_Good、RoomType_Bad
- int类型的使用
  - proto、model、dao 中，避免直接使用uint/int，应当使用 uint32/int32、uint64/int64
  - 其他和存储、传输无关场景（工具、算法），默认使用uint/int
  - 对于数量、定价等不可能为负数的情况，使用 uint/uint32/uint64
  - 对于差额、浮动数量、浮动价格等情况，使用 int/int32/int64
- float类型使用
  - 绝大多数对精度有要求场景（金额，比例等），避免使用float，优先使用int代替，备注中定义好单位1对应的数值
  - 对精度误差不敏感算法场景可使用float


## idl文件（proto文件）
编写后，需要执行下述命令，生成go代码
```bash
protoc --go_out=. --go-grpc_out=. ./idl/sched/*.proto
protoc --go_out=. --go-grpc_out=. ./idl/ota/*.proto
protoc --go_out=plugins=grpc:./ *.proto
```

## idl准备环境
```bash
brew install protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

