import { useEffect, useState } from "react";
import Sidebar from "../layouts/Sidebar";
import useSWR from "swr";
import { useNavigate } from "react-router-dom";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";

export default function Product() {
  const navigate = useNavigate();

  // Check if user is admin🐳
  const [isAdmin, setIsAdmin] = useState(() => {
    const { role } = Cookies.get("token")
      ? jwtDecode(Cookies.get("token"))
      : { role: "" };
    return role === "admin" ? true : false;
  });
  // Check if user is admin🐳

  // GET request to fetch all products🐳
  const [products, setProducts] = useState([]);
  const fetcher = (url, init) => fetch(url, init).then((res) => res.json());
  const { data, isLoading, error } = useSWR(
    "http://localhost:8080/v1/products",
    (url) =>
      fetcher(url, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
      })
  );

  useEffect(() => {
    if (data?.status) {
      setProducts(data.data);
    } else {
      console.log("Error fetching products");
    }
  }, [data]);
  // GET request to fetch all products🐳

  // DELETE request to delete a product🐳
  const handleDelete = (id) => {
    if (!isAdmin) {
      alert("You are not authorized to delete this product");
      return;
    }
    const isConfirm = confirm("Are you sure you want to delete this product?");
    if (!isConfirm) return;
    fetch(`http://localhost:8080/v1/products/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.status) {
          setProducts(products.filter((product) => product.id !== id));
        } else {
          console.log("Error deleting product");
        }
      });
  };
  // DELETE request to delete a product🐳

  return (
    <>
      <aside>
        <Sidebar></Sidebar>
      </aside>
      <main className="ml-72 p-10">
        <div className="flex justify-between items-center">
          <h1 className="text-2xl font-bold my-5">Product Table</h1>
          <button
            type="button"
            onClick={() => navigate("/product/new")}
            class="text-white bg-gradient-to-br from-green-400 to-blue-600 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-green-200 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
          >
            Add Product
          </button>
        </div>
        <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table className="w-full text-sm text-left rtl:text-right text-gray-500 ">
            <thead className="text-xs text-gray-700 uppercase bg-gray-50 ">
              <tr>
                <th scope="col" className="px-6 py-3">
                  Product name
                </th>
                <th scope="col" className="px-6 py-3">
                  Price
                </th>
                <th scope="col" className="px-6 py-3">
                  Stock
                </th>
              </tr>
            </thead>
            <tbody>
              {isLoading ? (
                <tr>
                  <td colSpan={4} className="px-6 py-4 text-center">
                    Loading...
                  </td>
                </tr>
              ) : (
                products.map((product) => (
                  <tr
                    key={product.id}
                    className="bg-white border-b hover:bg-gray-50"
                  >
                    <th
                      scope="row"
                      className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                    >
                      {product.name}
                    </th>
                    <td className="px-6 py-4">{product.price}</td>
                    <td className="px-6 py-4">{product.stock}</td>
                    <td className="flex items-center px-6 py-4">
                      <button
                        type="button"
                        onClick={() => navigate(`/product/edit/${product.id}`)}
                        className="font-medium text-blue-600 hover:underline"
                      >
                        Edit
                      </button>
                      <button
                        type="button"
                        onClick={() => handleDelete(product.id)}
                        className="font-medium text-red-600 hover:underline ms-3"
                      >
                        Remove
                      </button>
                    </td>
                  </tr>
                ))
              )}
            </tbody>
            <tfoot>
              <tr class="font-semibold text-gray-900">
                <th scope="row" class="px-6 py-3 text-base">
                  Total Product :
                </th>
                <td class="px-6 py-3"></td>
                <td class="px-6 py-3"></td>
                <td class="px-6 py-3">{products.length} Products</td>
              </tr>
            </tfoot>
          </table>
        </div>
      </main>
    </>
  );
}
