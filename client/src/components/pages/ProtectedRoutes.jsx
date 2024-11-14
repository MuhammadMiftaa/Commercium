import React from "react";
import { Navigate } from "react-router-dom";

export default function ProtectedRoutes({ children, isAuthenticated }) {
  return isAuthenticated ? children : <Navigate to="/auth/login" />;
}
