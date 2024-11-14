import React from "react";
import { useNavigate } from "react-router-dom";

export default function AdminRoutes({ children, isAdmin }) {
  return isAdmin ? children : <NotAdmin />;
}

function NotAdmin() {
  const navigate = useNavigate();
  return (
    <div
      className="bg-gradient-to-br from-blue-800 to-cyan-300 w-screen h-screen flex"
      style={{ backgroundImage: "/401.jpg" }}
    >
      <div className="mx-auto my-auto flex flex-col items-center">
        <h1 className="text-9xl font-bold">401</h1>
        <h2 className="text-6xl font-bold ">Unauthorized</h2>
        <br />
        <button
          onClick={() => {
            navigate(-1);
          }}
          className="text-4xl font-light hover:border-b-[1px] border-black"
        >
          Back to previous page
        </button>
      </div>
    </div>
  );
}
