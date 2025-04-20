import styles from '@/components/Post/Post.module.css'
import avatar from '@/assets/images/no-face.jpg'
import postImage from '@/assets/images/post1.jpg'
import likes from '@/assets/icons/like.svg'
// import like2 from '@/assets/icons/like2.svg'
import comment from '@/assets/icons/comment.svg'
import Image from 'next/image';
import PostLike from '../PostLike/PostLike'

const Post = ({data}) => {
    return (
        <article className={styles.card}>
            <div className={styles.card_header}>
                {/* avatar - username - creation_date */}
                <Image width={100} height={100} alt='avatar' src={avatar} className={styles.post_avatar} />
                <div className="">
                    <h4 className={styles.post_username}>{data.username}</h4>
                    <span className={styles.post_created_at}>{data.created_at}</span>
                </div>
            </div>
            <div className={styles.card_body}>
                {/* text - image - reactions_len - comments_len */}
                <p className={styles.post_text}>{data.content}</p>
                {/* <img src="" className={styles.post_image} alt="post image" /> */}
                <br />
                <Image width={500} height={500} alt='post' src={postImage} className={styles.post_image} />
                <div className={styles.postInfo}>
                    <div className={styles.post_reactions_len}>
                        <Image src={likes} width={24} height={24} alt='likes' className={styles.postLikesSvg} />
                        <span className={styles.postLikesNum}>{data.likes}</span>
                    </div>
                    <div className={styles.post_comments_len}>
                        <span className={styles.postCommentsNum}>{data.comments}</span>
                        <span className={styles.postCommentsText}>comments</span>
                    </div>
                </div>
            </div>
            <div className={styles.card_footer}>
                {/* add_reaction - add_comment */}
                <PostLike data={data} />
                <button className={styles.addComment}>
                <Image src={comment} width={20} height={20} alt='likes' className={styles.button} />
                    Comment</button>
            </div>
        </article>
    );
}

export default Post;