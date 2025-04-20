"use client"
import { useState } from 'react';
import styles from './PostLike.module.css'
import Image from 'next/image';
import likes from '@/assets/icons/like.svg'
import like2 from '@/assets/icons/like2.svg'

const PostLike = ({ data }) => {
    const [like, setLike] = useState(data.user_like);

    return (
        <button className={styles.setLike} onClick={()=> setLike((prevLike) => !prevLike)}>
            <Image src={like ? likes : like2} width={20} height={20} alt='likes' className={styles.button} />
            Like
        </button>
    );
}

export default PostLike;