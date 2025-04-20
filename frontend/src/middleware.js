// middleware.js
import { NextResponse } from "next/server";

export function middleware(request) {
  const session = request.cookies.get("session");

  const isLoggedIn = !!session?.value;

  if (!isLoggedIn && request.nextUrl.pathname !== "/login") {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/", "/chat"], // Protect these routes
};
