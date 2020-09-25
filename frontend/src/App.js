import React from 'react';
import { connect, sendMessage } from './api/index';
import Header from "./components/Header/Header"
import './App.css';

const App = () => {
  connect();

  const send = () => {
    console.log("hello");
    sendMessage("hello");
  }

  return (
    <div className="App">
      <Header />
      <button onClick={send}>Message</button>
    </div>
  );
}

export default App;
