import { Route, Routes } from "react-router-dom";
import "./App.css";
import Dashboard from "./components/pages/Dashboard";
import Product from "./components/pages/Product";
import Order from "./components/pages/Order";
import NewProduct from "./components/pages/NewProduct";
import UpdateProduct from "./components/pages/UpdateProduct";
import NotFound from "./components/pages/NotFound";
import Login from "./components/pages/Login";
import Register from "./components/pages/Register";

function App() {
  return (
    <Routes>
      <Route path="/">
        <Route index element={<Dashboard />} />
        <Route path="auth">
          <Route path="login" element={<Login />} />
          <Route path="register" element={<Register />} />
        </Route>
        <Route path="product">
          <Route index element={<Product />} />
          <Route path="new" element={<NewProduct />} />
          <Route path="update" element={<UpdateProduct />} />
        </Route>
        <Route path="order" element={<Order />} />
        <Route path="*" element={<NotFound />} />
      </Route>
    </Routes>
  );
}

export default App;
