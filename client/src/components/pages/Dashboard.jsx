import React from "react";
import Sidebar from "../layouts/Sidebar";
import { IoPerson } from "react-icons/io5";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";
import { useState } from "react";

export default function Dashboard() {
  // GET user role and username from cookies🐳

  const token = Cookies.get("token");
  const decodedToken = token ? jwtDecode(token) : {};
  const [role, setRole] = useState(decodedToken.role || "");
  const [username, setUsername] = useState(decodedToken.username || "");

  // GET user role and username from cookies🐳
  return (
    <>
      <aside>
        <Sidebar></Sidebar>
      </aside>
      <main className="ml-72 p-10">
        <div className="w-full h-full">
          <div className="mx-auto my-auto flex flex-col items-center">
            <h1 className="text-4xl font-bold ">Welcome to</h1>
            <h2 className="text-6xl font-bold ">Commercium Dashboard</h2>
            <div className="text-9xl bg-gradient-to-br from-zinc-300 to-zinc-100 rounded-full p-7 mt-16 border border-black">
              <IoPerson />
            </div>
            <div className="flex flex-col items-center mt-6">
              <h2 className="text-lg font-light -mb-2">Name</h2>
              <h3 className="capitalize text-4xl font-bold">{username}</h3>
            </div>
            <div className="flex flex-col items-center mt-6">
              <h2 className="text-lg font-light -mb-2">Role</h2>
              <h3 className="capitalize text-4xl font-bold">{role}</h3>
            </div>
          </div>
        </div>
      </main>
    </>
  );
}
