import React from 'react';
import logo from './logo.svg';
import { connect, sendMessage } from './api/index';
import './App.css';

const App = () => {
  connect();

  const send = () => {
    console.log("hello");
    sendMessage("hello");
  }

  return (
    <div className="App">
      <button onClick={send}>Message</button>
    </div>
  );
}

export default App;
