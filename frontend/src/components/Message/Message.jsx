import React, { useState, useEffect } from "react";
import "./Message.scss";

const Message = (props) => {
    const [message, setMessage] = useState("");

    useEffect(() => {
        let temp = JSON.parse(props.message)
        setMessage(temp)
    }, [props.message])

    return (
        <div className="Message">{message.body}</div>
    )
}

export default Message;
