import React, { useState } from "react";
import { useWasm } from "react-wasm";

function Game(){
  const [go, setGo] = useState(null);
  if (!go)  {
    setGo(new Go());
  }
  const importObject = (go || {}).importObject;
  const {
    loading,
    error,
    data
  } = useWasm({
    url: '/goloz.wasm',
    importObject: importObject,
  });

  if (loading) return "Loading Goloz..";
  if (error) return "An error has occurred";
  const { module, instance } = data;
  go.run(data.instance);
  return <div></div>;
}

export default Game;
