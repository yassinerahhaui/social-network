import styles from './NewPostCard.module.css'
import avatar from '@/assets/images/no-face.jpg'
import Image from 'next/image';

const NewPostCard = () => {
    return (
        <div className={styles.card}>
            <div className={styles.cardHeader}>
                <Image className={styles.createPostAvatar} src={avatar} width={150} height={150} alt="" />
                <form method="post" className={styles.form}>
                    <input type="text" name="" placeholder="What's on your mind..." className={styles.createPost} />
                </form>
            </div>
            <div className={styles.cardBody}>
                <button className={styles.editBtn}>Create Post</button>
                <button className={styles.editBtn}>Create Event</button>
            </div>
        </div>
    );
}

export default NewPostCard;