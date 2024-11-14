import React from 'react'
import { Link } from 'react-router-dom'
import { FaHome } from "react-icons/fa";
import { AiFillProduct } from "react-icons/ai";
import { HiNewspaper } from "react-icons/hi";
import { FaSignOutAlt } from "react-icons/fa";
import Cookies from "js-cookie";

export default function Sidebar() {
  const handleLogout = () => {
    Cookies.remove("token");
    window.location.href = "/auth/login";
  }

  return (
    <nav className='bg-[#0e1111] text-[#FEFFEC] fixed top-0 bottom-0 left-0 w-72 p-10 font-serif'>
        <div className='flex justify-between items-center'>
            {/* <img className='w-10 rounded-full' src="https://pbs.twimg.com/profile_images/1830068480554516480/8cKBeoiA_400x400.jpg" alt="Logo" /> */}
            <h1 className='font-bold text-3xl'>Commercium</h1>
        </div>
        <ul className='flex flex-col text-xl mt-10 mb-4 gap-2'>
            <li className='flex gap-1 items-center'><FaHome /><Link to={"/"}>Home</Link></li>
            <li className='flex gap-1 items-center'><AiFillProduct /><Link to={"/product"}>Product</Link></li>
            <li className='flex gap-1 items-center'><HiNewspaper /><Link to={"/order"}>Order</Link></li>
        </ul>
        <hr />
        <ul className='flex flex-col text-xl mt-4 gap-2'>
            <li className='flex gap-1 items-center text-red-400 ml-3'><FaSignOutAlt /><button onClick={handleLogout} type='button'>Logout</button></li>
        </ul>
    </nav>
  )
}
