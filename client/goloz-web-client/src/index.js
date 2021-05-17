import ReactDOM from 'react-dom';
import App from './App.jsx';

// WASM polyfill.
if (!WebAssembly.instantiateStreaming) {
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    return await WebAssembly.instantiate(source, importObject);
  };
}

const app = App();
ReactDOM.render(app, document.querySelector('#container'));

