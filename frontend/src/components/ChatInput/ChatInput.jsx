import React from 'react';
import "./ChatInput.scss";


const ChatInput = (props) => {
    const { send } = props;
    const placeholder = "Type a message... Hit Enter to Send";
    return (
        <div className="ChatInput">
            <input onKeyDown={send} placeholder={placeholder} />
        </div>
    )
}

export default ChatInput;
