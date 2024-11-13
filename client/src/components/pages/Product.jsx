import { useEffect, useState } from "react";
import Sidebar from "../layouts/Sidebar";
import useSWR from "swr";

export default function Product() {
  const [products, setProducts] = useState([]);
  const fetcher = (url) => fetch(url).then((res) => res.json());
  const { data, isLoading, error } = useSWR(
    "http://localhost:8080/v1/products",
    fetcher
  );
  useEffect(() => {
    if (data) {
      setProducts(data);
    }
    console.log(products);
  }, [data]);

  return (
    <>
      <aside>
        <Sidebar></Sidebar>
      </aside>
      <main className="ml-72 p-10">
        <h1 className="text-2xl font-bold my-5">Product Table</h1>
        <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table className="w-full text-sm text-left rtl:text-right text-gray-500 ">
            <thead className="text-xs text-gray-700 uppercase bg-gray-50 ">
              <tr>
                <th scope="col" className="px-6 py-3">
                  Product name
                </th>
                {/* <th scope="col" className="px-6 py-3">
                  Color
                </th>
                <th scope="col" className="px-6 py-3">
                  Category
                </th>
                <th scope="col" className="px-6 py-3">
                  Accessories
                </th>
                <th scope="col" className="px-6 py-3">
                  Available
                </th> */}
                <th scope="col" className="px-6 py-3">
                  Price
                </th>
                <th scope="col" className="px-6 py-3">
                  Weight
                </th>
                <th scope="col" className="px-6 py-3">
                  Stock
                </th>
              </tr>
            </thead>
            <tbody>
              <tr className="bg-white border-b hover:bg-gray-50">
                <th
                  scope="row"
                  className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
                >
                  Apple MacBook Pro 17"
                </th>
                {/* <td className="px-6 py-4">Silver</td>
                <td className="px-6 py-4">Laptop</td>
                <td className="px-6 py-4">Yes</td>
                <td className="px-6 py-4">Yes</td> */}
                <td className="px-6 py-4">$2999</td>
                <td className="px-6 py-4">1000</td>
                <td className="flex items-center px-6 py-4">
                  <a href="#" className="font-medium text-blue-600 hover:underline">
                    Edit
                  </a>
                  <a
                    href="#"
                    className="font-medium text-red-600 hover:underline ms-3"
                  >
                    Remove
                  </a>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </main>
    </>
  );
}
