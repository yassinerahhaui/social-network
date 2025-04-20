"use client";
// components/SocketComponent.js
import styles from "./style.module.css";
import { useContext, useEffect } from "react";
import { Context } from "../../context/context.js";

const ChatComponent = () => {
  // useEffect(() => {
  //   const ws = new WebSocket("ws://localhost:8080/api/ws");
  //   ws.onopen = () => {
  //     console.log("Connected to WebSocket");
  //     ws.send("Hello, WebSocket!");
  //   };
  //   ws.onmessage = (event) => {
  //     console.log("Message received:", event.data);
  //   };
  //   ws.onclose = () => {
  //     console.log("WebSocket connection closed");
  //   };
  //   return () => {
  //     ws.close();
  //   };
  // }, []);
  // const { username } = useContext(Context);

  // useEffect(() => {
    // console.log("Username in context has been set:", username); // This will log the updated username
  // }, [username]); // Track username changes

  // return <div>Current Username: {username}</div>;
  return (
    <div className="main">
      <Sidebar />
      <Main />
    </div>
  );
};

function Main() {
  return (
    <div className="main-section">
      <ChatWindow />
      <MessageInput />
    </div>
  );
}

function Sidebar() {
  return (
      <div className="sidebar">
        <div className="user-info">USER INFO</div>
        <button className="compose">+ Compose</button>
        <div className="contact-list">CONTACT list</div>
      </div>
  );
}

function ChatWindow() {
  return (
    <div className="chat-window">
      <div className="chat-header">
        <div className="chat-user-info">
          <img src="user-avatar.jpg" alt="User" className="avatar" />
          <div>
            <strong>Anthony M. Conley</strong>
            <p className="location">Timisoara, Romania</p>
          </div>
        </div>
      </div>

      <div className="messages">
        <div className="message received">
          Vivendo periculis mel ut, eam ei quem tota.
        </div>
        <div className="message received">
          In nec mucius recteque concludaturque.
        </div>
        <div className="date-separator">Today</div>
        <div className="message sent">Dicunt delectus salutatus nam te.</div>
        <div className="message sent">
          Vivendo periculis mel ut, eam ei quem tota.
        </div>
      </div>
    </div>
  );
}

function MessageInput() {
  return (
    <div className="message-input">
      <input placeholder="Write a reply..." />
      <button className="send">Send</button>
    </div>
  );
}
export default ChatComponent;
