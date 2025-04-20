"use client";

import styles from "./style.module.css";
import { useEffect} from "react";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  console.log('login page ')
  // const { username, setUsername } = useContext(Context);
  // console.log('we started with this username:', username )
  const router = useRouter();

  useEffect(() => {
    const form = document.getElementById("login-form");
    const errorMsg = document.getElementById("error-message");

    form.addEventListener("submit", async (e) => {
      e.preventDefault();

      const nickname = form.nickname.value;
      const password = form.password.value;

      console.warn(nickname, password);

      try {
        const response = await fetch("http://localhost:8080/api/login", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ nickname, password }),
          credentials: "include", // Correctly sends cookies
        });

        if (response.ok) {
          //setUsername(nickname); 
          //localStorage.setItem('username', nickname); // Persist to localStorage
          router.push("/"); // Use client-side navigation // navigat is for server side
        }
      } catch (err) {
        console.error(err.message);
        errorMsg.textContent = err.message; // Display error message to the user
      }
    });
  }, []);

  return (
    <div className={styles.loginContainer}>
      <h1>Login</h1>
      <form id="login-form" className={styles.loginForm}>
        <label htmlFor="nickname" className={styles.label}>
          nickname:
        </label>
        <input
          type="nickname"
          id="nickname"
          name="nickname"
          required
          className={styles.input}
        />

        <label htmlFor="password" className={styles.label}>
          Password:
        </label>
        <input
          type="password"
          id="password"
          name="password"
          required
          className={styles.input}
        />

        <button type="submit" className={styles.button}>
          Login
        </button>
        <p className={styles.linkText}>
          Don't have an account?{" "}
          <a href="/register" className={styles.link}>
            Register here
          </a>
        </p>
        <p id="error-message" style={{ color: "red" }}></p>
      </form>
    </div>
  );
}
