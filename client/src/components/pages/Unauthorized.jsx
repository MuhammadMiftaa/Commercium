import React from "react";

export default function AdminRoutes({ children, isAdmin }) {
  return isAdmin ? children : <NotAdmin />;
}

function NotAdmin() {
  return <div>NotAdmin</div>;
}
