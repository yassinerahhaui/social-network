import Link from 'next/link';
import styles from './ProfileFriends.module.css';
import { backendUrl } from '@/utils/ustil';
import Image from 'next/image';
import avatar from "@/assets/images/no-face.jpg"

const ProfileFriends = async () => {
    const res = await fetch(`${backendUrl}/profile-friends`, {cache: 'no-store'})
    const friends = await res.json()
    return (
        <div className={styles.card}>
            <div className={styles.cardHeader}>
                <div className="">
                    <h3 className={styles.cardTitle}>Friends</h3>
                    <span className={styles.cardText}>301 Friends</span>
                </div>
                <Link href={'/'} className={styles.cardLinkBtn}>See all friends</Link>
            </div>
            <div className={styles.cardBody}>
                {friends.map(el=> {
                    return (
                        <Link href={'/'} className={styles.friendCard} key={el.id}>
                            <Image src={avatar} className={styles.friendImg} width={200} height={200} alt='user image' />
                            <h5 className=''>{el.name}</h5> 
                        </Link>
                    )
                })}
            </div>
        </div>
    );
}

export default ProfileFriends;