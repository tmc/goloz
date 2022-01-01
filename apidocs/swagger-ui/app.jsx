import * as React from 'react'
import * as ReactDOM from 'react-dom'


import SwaggerUI from "swagger-ui-react"
import "swagger-ui-react/swagger-ui.css"

let App = () => <SwaggerUI url="api.openapi.json" />

ReactDOM.render(
  <App />,
  document.getElementById('root')
);
