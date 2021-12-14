# DELETE OUT
# shellcheck disable=SC2034
build_out=$(dirname $(readlink -f "$0"))/out/post_model_predict
# delete output
rm -r "$build_out"
# BUILD
# shellcheck disable=SC2035
go build -o "$build_out" -v -x  *.go
