## Protobuf gRPC Go

### 소개
> grpc protocol buffers submodule 저장소
- [Go gRPC](https://grpc.io/docs/languages/go/quickstart)
- [Protocol Buffers](https://protobuf.dev)

### 구조
```
- gen
: generate 된 파일 장소
- plugins
: generate 를 위한 언어별 플러그인 모음
- protos
: proto 파일 모음
```

### 시작하기
```bash
# 모든 proto generate
make gen.proto

# 특정 proto golang generate
make gen.go.file

# 모든 proto golang generate
make gen.go

# 특정 proto c# generate
make gen.csharp.file

# 모든 proto c# generate
make gen.csharp
```
