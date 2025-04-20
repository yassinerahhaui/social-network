'use client';
import React from 'react';
import CreatePost from '@/components./CreatePost/CreatePost';
import NewPostCard from '@/components./NewPostCard/NewPostCard';

export default async function GroupPage() {
  //logic
  const { groupId } = useParams();
  // // Fetch group data using groupId
  // const groupData = await GetGroup(groupId);
  console.log(groupId);
  return (
    <>
      <div>
        <CreatePost />
        <NewPostCard />
      </div>
    </>
  );
}
