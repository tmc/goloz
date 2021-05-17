// goloz React web client

if (!WebAssembly.instantiateStreaming) {
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}
const go = new Go();
WebAssembly.instantiateStreaming(fetch("goloz.wasm"), go.importObject).then(result => {
  go.run(result.instance);
});
