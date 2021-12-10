RPC_REPO_PATH=post_model_manage
CLIENT_PATH=rpc_sdk
IDL_PATH=idl/post_model_manage.proto

protoc  --go_out=$RPC_REPO_PATH\
  --go_opt=paths=source_relative\
   --go-grpc_out=$RPC_REPO_PATH \
    --go-grpc_opt=paths=source_relative\
      $IDL_PATH   