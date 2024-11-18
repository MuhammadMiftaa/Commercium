import React, { useEffect, useState } from "react";
import Sidebar from "../layouts/Sidebar";
import { jwtDecode } from "jwt-decode";
import Cookies from "js-cookie";
import useSWR from "swr";
import { FaCheck } from "react-icons/fa6";

export default function Order() {
  // Check if user is admin🐳
  const [isAdmin, setIsAdmin] = useState(() => {
    const { role } = Cookies.get("token")
      ? jwtDecode(Cookies.get("token"))
      : { role: "" };
    return role === "admin";
  });
  // Check if user is admin🐳

  // GET request to fetch all products🐳
  const [orders, setOrders] = useState([]);
  const fetcher = (url, init) => fetch(url, init).then((res) => res.json());
  const { data, isLoading, error } = useSWR(
    "http://localhost:8080/v1/orders",
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
    if (data?.status) setOrders(data.data);
  }, [data]);
  // GET request to fetch all products🐳

  // PUT request to update order status🐳
  const [errorMessage, setErrorMessage] = useState(null);
  const handlePaidOrder = (id) => {
    fetch(`http://localhost:8080/v1/orders/${id}/paid`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
    })
      .then((res) => res.json())
      .then((data) => {
        if (data.status) {
          const updatedOrders = orders.map((order) =>
            order.id === id ? { ...order, status: "paid" } : order
          );
          setOrders(updatedOrders);
          alert("Success paid order");
        } else {
          setErrorMessage(data.message);
        }
      })
      .catch((error) => {
        setErrorMessage(error.message);
      });
  };

  useEffect(() => {
    if (errorMessage) alert(errorMessage);
    setErrorMessage(null);
  }, [errorMessage]);
  // PUT request to update order status🐳

  // Search order by customer name🐳
  const [search, setSearch] = useState("");
  const [filteredOrders, setFilteredOrders] = useState([]);
  useEffect(() => {
    setFilteredOrders(orders);
    const filteredOrders = orders.filter((order) =>
      order.customer_name.toLowerCase().includes(search.toLowerCase())
    );
    setFilteredOrders(filteredOrders);
  }, [search, orders]);
  // Search order by customer name🐳

  return (
    <>
      <aside>
        <Sidebar></Sidebar>
      </aside>
      <main className="ml-72 p-10">
        <div className="flex justify-between items-center mb-10">
          <h1 className="text-4xl font-bold">Order List</h1>

          <div className="w-64">
            <label
              htmlFor="default-search"
              className="mb-2 text-sm font-medium text-gray-900 sr-only"
            >
              Search
            </label>
            <div className="relative">
              <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                <svg
                  className="w-4 h-4 text-gray-500"
                  aria-hidden="true"
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 20 20"
                >
                  <path
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                  />
                </svg>
              </div>
              <input
                onChange={(e) => setSearch(e.target.value)}
                type="search"
                id="default-search"
                className="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-2xl bg-gray-50 focus:ring-blue-500 focus:border-blue-500"
                placeholder="Search by Customer Name"
                required
              />
            </div>
          </div>
        </div>
        <div className="relative">
          <table className="w-full text-sm text-left rtl:text-right text-gray-500">
            <thead className="text-xs text-gray-700 uppercase bg-gray-100">
              <tr>
                <th scope="col" className="px-6 py-3 rounded-l-lg">
                  ID
                </th>
                <th scope="col" className="px-6 py-3">
                  Customer Name
                </th>
                <th scope="col" className="px-6 py-3">
                  Product name
                </th>
                <th scope="col" className="px-6 py-3">
                  Qty
                </th>
                <th scope="col" className="px-6 py-3 text-center">
                  Price
                </th>
                <th scope="col" className="px-6 py-3 text-center">
                  Total Price
                </th>
                <th scope="col" className="px-6 py-3 text-center">
                  Status
                </th>
                <th
                  scope="col"
                  className="px-6 py-3 text-center rounded-e-lg"
                ></th>
              </tr>
            </thead>
            <tbody>
              {isLoading ? (
                <tr>
                  <td colSpan={7} className="px-6 py-4 text-center">
                    Loading...
                  </td>
                </tr>
              ) : (
                filteredOrders.map((order) => (
                  <tr className="bg-white" key={order.id}>
                    <th
                      scope="row"
                      className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                    >
                      {order.id}
                    </th>
                    <td className="px-6 py-4 whitespace-nowrap text-zinc-800 font-semibold">
                      {order.customer_name}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-zinc-800">
                      {order.product_name}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-zinc-800">
                      {order.quantity}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-zinc-800">
                      {formatRupiah(order.product_price)}
                    </td>
                    <td className="px-6 py-4 whitespace-nowrap text-zinc-800">
                      {formatRupiah(order.total_price)}
                    </td>
                    <td
                      className={`px-6 py-4 whitespace-nowrap text-zinc-800 text-center uppercase font-bold ${formatStatus(
                        order.status
                      )}`}
                    >
                      {order.status}
                    </td>
                    <td>
                      {isAdmin && (
                        <button
                          disabled={order.status === "paid"}
                          onClick={() => handlePaidOrder(order.id)}
                          type="button"
                          className={`px-4 py-2 ${
                            order.status === "pending"
                              ? "bg-blue-500 active:bg-blue-800"
                              : "bg-zinc-500"
                          } text-white rounded-md`}
                        >
                          <FaCheck />
                        </button>
                      )}
                    </td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>
      </main>
    </>
  );
}

const formatRupiah = (value) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(value);
};

const formatStatus = (status) => {
  switch (status) {
    case "pending":
      return "text-yellow-400";
    case "paid":
      return "text-green-600";
    case "cancelled":
      return "text-red-500";
    default:
      return "text-gray-500";
  }
};
