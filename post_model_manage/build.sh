# DELETE OUT
# shellcheck disable=SC2034
build_out=$(pwd)/out/post_model_manage
# delete output
rm -r "$build_out"
# BUILD
# shellcheck disable=SC2035
go build -o "$build_out" -v -x  *.go
