import React from "react";
import Sidebar from "../layouts/Sidebar";
export default function Dashboard() {
  return (
    <>
      <aside>
        <Sidebar></Sidebar>
      </aside>
      <main className="ml-80">Dashboard</main>
    </>
  );
}
