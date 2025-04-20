"use client"
import React , { useState , createContext } from "react";

export const Context = createContext();

export const ContextProvider = (props) => {
  const [username, setUsername] = useState(localStorage.getItem('username') || null);

  const value = {
    username,
    setUsername
  };

  return <Context.Provider value={value}>{props.children}</Context.Provider>;
};
