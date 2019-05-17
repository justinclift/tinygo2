A copy of the TinyGo gc ram allocation util, used while working on the TinyGo gc

  https://github.com/tinygo-org/tinygo/blob/11567c62d45418685161cd484b55c225f1d1f65d/testdata/gc.go

Compiled using:

    $ tinygo build -target wasm -gc dumb -no-debug -wasm-abi generic -o wasteram.wasm main.go
    $ wasm2wat -f --generate-names -o wasteram.wat wasteram.wasm
