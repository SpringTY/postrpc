# DELETE OUT
# shellcheck disable=SC2034
build_out=$(pwd)/out/
# delete output
rm -r "$build_out"
# BUILD
go build -o "$build_out" -v -x
