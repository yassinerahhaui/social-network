"use client"
import { useState } from "react"
import styles from "./CreatePost.module.css"
import Form  from "next/form"
// import { HandleForm } from "@/services/CreatePost"
/*
    - content
    - image
*/


const CreatePost = () => {
    const [status,setStatus] = useState("public")
    const [friends,setFriends] = useState(["yassine","oussama","elfihry","brahim","naytderhm","ghost"])
    const [viewers,setViewers] = useState([])
    const HandleForm = async (event) => {
        const formData = event.target
        const content = formData.get("content")
        const image = formData.get("image")
        console.log(content);
        console.log(image);
    }
    return (
        <>
            <div className={styles.overlay} id="overlay"></div>
            <div className={styles.card}>
                <h3 className={styles.cardTitle}>Create Post</h3>
                <div className={styles.cardBody}>
                    <Form onSubmit={HandleForm}>
                        <textarea name="content" placeholder="What's on your mind..." className={styles.formContent}></textarea>
                        <div className={styles.formGroup}>
                            <select name="status" onChange={(e)=> setStatus(e.target.value)} value={status} className={styles.status}>
                                <option value="public">Public</option>
                                <option value="private">Private</option>
                                <option value="almost-private">Almost Private</option>
                            </select>
                            <label htmlFor="imageUpload" className={styles.imgUploadLabel}>Upload 📷</label>
                            <input type="file" name="image" className={styles.formImage} id="imageUpload" />
                        </div>
                        {status === "private" ? 
                            <select name="viewers" className={styles.viewers}  onChange={(e)=> {setViewers([...viewers,e.target.value])}}>
                                <option value="" className={styles.select}>select</option>
                                {friends.map(el=> <option key={el} value={el}>{el}</option>)}
                            </select> : 
                        ""}
                        {status === "private" ? <span className={styles.seeText}>Only you can see this post {viewers.length > 0 ? 'and' : ''}</span>: ''}
                        {status === "private" ? 
                        <div className="">
                            {viewers.map(el=> <span key={el} className={styles.selectedViewer}>{el}</span>)}
                        </div>:
                        ""}
                        <hr className={styles.line} />
                        <button type="submit" className={styles.formSubmit}>submit</button>
                    </Form>
                </div>
                <div className={styles.cardFooter}></div>
            </div>
        </>
    );
}

export default CreatePost;