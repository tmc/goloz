// goloz React web client
import React from "react";

import Game from './Game.jsx';

import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

export default function App() {
  return (
    <Router>
      <div>
        <Switch>
          <Route path="/">
            <Game />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

