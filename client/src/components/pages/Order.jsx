import React, { useEffect, useState } from "react";
import Sidebar from "../layouts/Sidebar";
import { jwtDecode } from "jwt-decode";
import Cookies from "js-cookie";
import useSWR from "swr";

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

  return (
    <>
      <aside>
        <Sidebar></Sidebar>
      </aside>
      <main className="ml-72 p-10">
        <h1 className="text-4xl font-bold mb-10">Order List</h1>
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
                <th scope="col" className="px-6 py-3 rounded-e-lg">
                  Status
                </th>
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
                orders.map((order) => (
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
                    <td className={`px-6 py-4 whitespace-nowrap text-zinc-800 text-center uppercase font-bold ${formatStatus(order.status)}`}>
                      {order.status}
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
  return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0
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
}