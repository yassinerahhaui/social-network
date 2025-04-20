"use client"
import { createContext, useState } from "react"

export const GlobalContext = createContext()


export const GlobalProvider = ({ children }) => {
    const [overlay, setOverlay] = useState(false)
    const [showNav, setShowNav] = useState(false)
    const [showCreatePost, setShowCreatePost] = useState(false)
    return (
        <GlobalContext.Provider value={{ 
            overlay, 
            setOverlay, 
            showNav, 
            setShowNav, 
            showCreatePost, 
            setShowCreatePost 
        }}>
            {children}
        </GlobalContext.Provider>
    )
}