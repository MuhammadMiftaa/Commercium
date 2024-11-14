import React from 'react'
import { useNavigate } from 'react-router-dom';

export default function NotFound() {
  const navigate = useNavigate();
  return (
    <div
      className="bg-gradient-to-br from-rose-800 to-pink-300 w-screen h-screen flex text"
      style={{ backgroundImage: "/401.jpg" }}
    >
      <div className="mx-auto my-auto flex flex-col items-center">
        <h1 className="text-9xl font-bold">404</h1>
        <h2 className="text-6xl font-bold ">Not Found</h2>
        <br />
        <button
          onClick={() => {
            navigate("/");
          }}
          className="text-4xl font-light hover:border-b-[1px] border-black"
        >
          Back to dashboard
        </button>
      </div>
    </div>
  );
}
