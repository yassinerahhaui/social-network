"use client"
import styles from './NewPostCard.module.css'
import avatar from '@/assets/images/no-face.jpg'
import { GlobalContext } from '@/contexts/GlobalContext';
import Image from 'next/image';
import { useContext } from 'react';

const NewPostCard = () => {
    const { setOverlay, setShowCreatePost } = useContext(GlobalContext)
    return (
        <div className={styles.card}>
            <div className={styles.cardHeader}>
                <Image className={styles.createPostAvatar} src={avatar} width={150} height={150} alt="" />
                <form method="post" className={styles.form}>
                    <input
                        type="text"
                        name=""
                        placeholder="What's on your mind..."
                        className={styles.createPost}
                        onClick={() => {
                            setOverlay(true);
                            setShowCreatePost(true);
                        }}
                    />
                </form>
            </div>
            <div className={styles.cardBody}>
                <button className={styles.editBtn} onClick={() => {
                    setOverlay(true);
                    setShowCreatePost(true);
                }}>Create Post</button>
                <button className={styles.editBtn}>Create Event</button>
            </div>
        </div>
    );
}

export default NewPostCard;