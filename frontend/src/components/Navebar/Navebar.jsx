"use client";

import { usePathname } from 'next/navigation';
import { useContext, useEffect, useState } from "react";
import styles from "./Navbar.module.css";
import Image from "next/image";
import Logo from "@/assets/logo.png";
import HomeIcon from "@/assets/icons/home.png";
import GroupsIcon from "@/assets/icons/groups.png";
import EventIcon from "@/assets/icons/event.png";
import FollowersIcon from "@/assets/icons/followers.png";
import ProfileIcon from "@/assets/icons/profile.png";
import LoginIcon from "@/assets/icons/login.png";
import RegisterIcon from "@/assets/icons/pen.png";
import LogoutIcon from "@/assets/icons/logout.png";
import NotifIcon from "@/assets/icons/notification.png";
import MenuIcon from "@/assets/icons/menu.png";
import CloseIcon from "@/assets/icons/close.png";
import Link from "next/link";
import { GlobalContext } from "@/contexts/GlobalContext";

const Navebar = () => {
  const pathname = usePathname();
  const hideOn = ["/login", "/register"];

  if (hideOn.includes(pathname)) return null;
  const [isAuth, setIsAuth] = useState(true);
  const [scWidth, setScWidth] = useState(
    typeof window !== "undefined" ? window.innerWidth : 0
  );
  const [showMenu, setShowMenu] = useState("");
  const { setOverlay, showNav, setShowNav } = useContext(GlobalContext);

  useEffect(() => {
    showNav
      ? setShowMenu("block")
      : scWidth < 960
      ? setShowMenu("none")
      : setShowMenu("flex");
  }, [showNav]);

  useEffect(() => {
    const handleResize = () => {
      setScWidth(window.innerWidth);
    };
    if (scWidth > 960 && showNav) setShowNav(false);
    showNav
      ? setShowMenu("block")
      : scWidth < 960
      ? setShowMenu("none")
      : setShowMenu("flex");

    window.addEventListener("resize", handleResize);
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, [scWidth]);

  return (
    <header className={styles.navbar}>
      <div className={styles.navbar_start}>
        {/* [logo] */}
        <Link href={"/"}>
          <Image
            src={Logo}
            className={styles.logo}
            width={60}
            height={50}
            alt="logo"
          />
        </Link>
      </div>
      <menu className={styles.menu} style={{ display: showMenu }}>
        <button
          className={styles.closeBtn}
          onClick={() => {
            setShowNav(false);
            setOverlay(false);
          }}
        >
          <Image src={CloseIcon} width={20} height={20} alt="close button" />
        </button>
        <nav
          className={styles.navbar_middle}
          style={showNav ? { display: "block" } : { display: "flex" }}
          onClick={() =>
            showNav ? (setShowNav(false), setOverlay(false)) : ""
          }
        >
          {/* [home - groups - followers - events] */}
          <Link href={"/"} className={styles.link}>
            <Image
              className={styles.pageIcon}
              src={HomeIcon}
              width={36}
              height={36}
              alt="home icon link"
            />
            <span className={styles.pageName}>Home</span>
          </Link>
          <Link href={"/groups"} className={styles.link}>
            <Image
              className={styles.pageIcon}
              src={GroupsIcon}
              width={36}
              height={36}
              alt="groups icon link"
            />
            <span className={styles.pageName}>Groups</span>
          </Link>
          <Link href={"/events"} className={styles.link}>
            <Image
              className={styles.pageIcon}
              src={EventIcon}
              width={36}
              height={36}
              alt="events icon link"
            />
            <span className={styles.pageName}>Events</span>
          </Link>
          <Link href={"/followers"} className={styles.link}>
            <Image
              className={styles.pageIcon}
              src={FollowersIcon}
              width={36}
              height={36}
              alt="followers icon link"
            />
            <span className={styles.pageName}>Followers</span>
          </Link>
        </nav>
        <div
          className={styles.navbar_end}
          style={showNav ? { display: "block" } : { display: "flex" }}
          onClick={() =>
            showNav ? (setShowNav(false), setOverlay(false)) : ""
          }
        >
          {/* [login - register] OR [profile - logout] */}
          {isAuth ? (
            <>
              <Link href={"/notification"} className={styles.link}>
                <Image
                  className={styles.pageIcon}
                  src={NotifIcon}
                  width={38}
                  height={38}
                  alt="notification icon link"
                />
                <span className={styles.pageName}>Notifications</span>
              </Link>
              <Link href={`/profile/yrahhaou`} className={styles.link}>
                <Image
                  className={styles.pageIcon}
                  src={ProfileIcon}
                  width={38}
                  height={38}
                  alt="profile icon link"
                />
                <span className={styles.pageName}>Profile</span>
              </Link>
              <Link href={"/logout"} className={styles.link}>
                <Image
                  className={styles.pageIcon}
                  src={LogoutIcon}
                  width={38}
                  height={38}
                  alt="logout icon link"
                />
                <span className={styles.pageName}>Logout</span>
              </Link>
            </>
          ) : (
            <>
              <Link href={"/login"} className={styles.link}>
                <Image
                  className={styles.pageIcon}
                  src={LoginIcon}
                  width={38}
                  height={38}
                  alt="login icon link"
                />
                <span className={styles.pageName}>Login</span>
              </Link>
              <Link href={"/register"} className={styles.link}>
                <Image
                  className={styles.pageIcon}
                  src={RegisterIcon}
                  width={38}
                  height={38}
                  alt="register icon link"
                />
                <span className={styles.pageName}>Register</span>
              </Link>
            </>
          )}
        </div>
      </menu>
      <button
        className={styles.bergerMenu}
        onClick={(e) => {
          setShowNav(true);
          setOverlay(true);
        }}
      >
        <Image
          className={styles.bergerMenuIcon}
          src={MenuIcon}
          width={30}
          height={30}
          alt="menu button icon"
        />
      </button>
    </header>
  );
};

export default Navebar;
