import { Route, Routes } from "react-router-dom";
import "./App.css";
import Dashboard from "./components/pages/Dashboard";
import Product from "./components/pages/Product";
import Order from "./components/pages/Order";
import NewProduct from "./components/pages/NewProduct";
// import NewProduct from "./components/pages/NewProductWithImage";
import UpdateProduct from "./components/pages/UpdateProduct";
import NotFound from "./components/pages/NotFound";
import Login from "./components/pages/Login";
import Register from "./components/pages/Register";
import { useState } from "react";
import Cookies from "js-cookie";
import ProtectedRoutes from "./components/pages/ProtectedRoutes";

function App() {

  const [isAuthenticated, setIsAuthenticated] = useState(() => {
    return Cookies.get("token") !== undefined ? true : false;
  })

  const handleLogin = () => {
    setIsAuthenticated(true);
  }

  return (
    <Routes>
      <Route path="/">
        <Route index element={<ProtectedRoutes isAuthenticated={isAuthenticated}><Dashboard /></ProtectedRoutes>} />
        <Route path="auth">
          <Route path="login" element={<Login handleLogin={handleLogin} />} />
          <Route path="register" element={<Register />} />
        </Route>
        <Route path="product">
          <Route index element={<ProtectedRoutes isAuthenticated={isAuthenticated}><Product /></ProtectedRoutes>} />
          <Route path="new" element={<ProtectedRoutes isAuthenticated={isAuthenticated}><NewProduct /></ProtectedRoutes>} />
          <Route path="edit/:id" element={<ProtectedRoutes isAuthenticated={isAuthenticated}><UpdateProduct /></ProtectedRoutes>} />
        </Route>
        <Route path="order" element={<Order />} />
        <Route path="*" element={<NotFound />} />
      </Route>
    </Routes>
  );
}

export default App;
