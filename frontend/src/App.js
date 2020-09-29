import React, { useState } from "react";
import { connect, sendMessage } from "./api/index";
import Header from "./components/Header/Header"
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import "./App.css";

const App = () => {
  const [chatHistory, setChatHistory] = useState([]);

  connect((message) => {
    console.log("New message");
    const newChats = [...chatHistory, message]
    setChatHistory(newChats);
    console.log(newChats);
  });

  const send = (event) => {
    if (event.keyCode === 13) {
      sendMessage(event.target.value);
      event.target.value = "";
    }
  }

  return (
    <div className="App">
      <Header />
      <ChatHistory chatHistory={chatHistory} />
      <ChatInput send={send} />
    </div>
  );
}

export default App;
