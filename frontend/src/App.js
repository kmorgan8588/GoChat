import React, { useState } from 'react';
import { connect, sendMessage } from './api/index';
import Header from "./components/Header/Header"
import ChatHistory from './components/ChatHistory/ChatHistory';
import './App.css';

const App = () => {
  const [chatHistory, setChatHistory] = useState([]);

  connect((message) => {
    console.log("New message");
    const newChats = [...chatHistory, message]
    setChatHistory(newChats);
    console.log(newChats);
  });

  const send = () => {
    console.log("hello");
    sendMessage("hello");
  }

  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <button onClick={send}>Message</button>
    </div>
  );
}

export default App;
