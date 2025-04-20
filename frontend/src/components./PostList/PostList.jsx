"use client";
import { useEffect, useState } from "react";
import Post from "../Post/Post";
import { validbackendUrl } from "@/utils/ustil";

const PostList = () => {
  const [posts, setPosts] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch(`${validbackendUrl}/api/posts?limit=10&offset=0`, {
      cache: "no-store",
      credentials: "include",
    })
      .then(async (res) => {
        const contentType = res.headers.get("content-type");
        if (!res.ok) {
          const errorBody = contentType?.includes("application/json")
            ? await res.json()
            : await res.text();
          throw new Error(
            errorBody || `Request failed with status ${res.status}`
          );
        }
        return res.json();
      })
      .then(setPosts)
      .catch((err) => {
        console.error("Fetch error:", err.message);
        setError(err.message);
      });
  }, []);

  if (error) return <div>Error: {error}</div>;

  console.log('fetched posts:', posts)
 
  return (
    <section>
      {posts.map((post) => (
        <Post key={post.id} data={post} />
      ))}
    </section>
  );
};

export default PostList;
