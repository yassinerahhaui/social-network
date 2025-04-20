import Image from "next/image";
import styles from './page.module.css'
import coverImg from '@/assets/images/coverImg.jpg'
import avatar from '@/assets/images/no-face.jpg'
import NewPostCard from "@/components/NewPostCard/NewPostCard";
import ProfileFriends from "@/components/ProfileFriends/ProfileFriends";
import PostList from "@/components/PostList/PostList";
import CreatePost from "@/components/CreatePost/CreatePost";


export async function generateMetadata({ params }) {
    const { username } = await params;
    return {
        title: username,
    }
}

const ProfilePage = async ({ params }) => {
    const { username } = await params;
    return (
        <>
            <div className={styles.header}>
                <Image className={styles.coverImg} src={coverImg} width={1200} height={500} alt="" />
                <div className={styles.headerInfo}>
                    <div className={styles.headerInfoStart}>
                        <Image className={styles.avatar} src={avatar} width={150} height={150} alt="" />
                        <h1 className={styles.headerInfoUsername}>{username}</h1>
                    </div>
                    {/* <div className={styles.headerInfoEnd}>
                        <button className={styles.editBtn}>Edit Profile</button>
                        <button className={styles.editBtn}>Create Post</button>
                    </div> */}
                </div>
                <CreatePost />
            </div>
            <main className={styles.content}>
                <aside className={styles.sideBar}>
                    <ProfileFriends />
                </aside>
                <div className={styles.profilePosts}>
                    <NewPostCard />
                    <PostList />
                </div>
            </main>
        </>
    );
}

export default ProfilePage;