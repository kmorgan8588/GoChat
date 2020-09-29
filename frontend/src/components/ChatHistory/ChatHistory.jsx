import React from 'react';
import Message from "../Message/Message";
import "./ChatHistory.scss";

const ChatHistory = (props) => {
    const { chatHistory } = props;

    const messages = chatHistory.map((msg, index) => (<Message message={msg.data} key={index} />))



    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    )
};

export default ChatHistory;
