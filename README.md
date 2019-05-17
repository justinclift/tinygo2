Just a small util to (hopefully) waste ram.  For testing wasm gc.

Compiles using:

    $ tinygo build -target wasm -gc dumb -no-debug -wasm-abi generic -o wasteram.wasm main.go
