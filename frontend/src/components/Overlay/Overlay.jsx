"use client"
import { useContext } from "react";
import styles from "./Overlay.module.css"
import { GlobalContext } from "@/contexts/GlobalContext";

const Overlay = () => {
    const {overlay, setOverlay, showNav, setShowNav, showCreatePost, setShowCreatePost} = useContext(GlobalContext)
    return (
        <div onClick={()=> {
            setOverlay(false);
            showNav ? setShowNav(false) : '';
            showCreatePost ? setShowCreatePost(false) : '';
        }} className={overlay ? styles.display : styles.hidden}></div>
    );
}

export default Overlay;