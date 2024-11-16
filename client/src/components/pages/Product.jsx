import { useEffect, useState } from "react";
import Sidebar from "../layouts/Sidebar";
import useSWR from "swr";
import { useNavigate, useSearchParams } from "react-router-dom";
import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";
import { FaCircleChevronLeft } from "react-icons/fa6";
import { FaCircleChevronRight } from "react-icons/fa6";

export default function Product() {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const [page, setPage] = useState(Number(searchParams.get("page")) || 1);

  // Check if user is admin🐳
  const [isAdmin, setIsAdmin] = useState(() => {
    const { role } = Cookies.get("token")
      ? jwtDecode(Cookies.get("token"))
      : { role: "" };
    return role === "admin";
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
    if (data?.status) setProducts(data.data);
  }, [data]);
  // GET request to fetch all products🐳

  // Sort and Order🐳
  const [sort, setSort] = useState(searchParams.get("sort") || "");
  const [order, setOrder] = useState(searchParams.get("order") || "");
  useEffect(() => {
    if (sort !== "" && order !== "") {
      const paramsObject = Object.fromEntries(searchParams.entries());
      setSearchParams({ ...paramsObject, sort, order });
      const productSorted = [...products].sort((a, b) => {
        if (typeof a[sort] === "string") {
          return a[sort].localeCompare(b[sort]);
        }
        return a[sort] - b[sort];
      });
      setProducts(order === "asc" ? productSorted : productSorted.reverse());
    }
  }, [sort, order]);
  // Sort and Order🐳

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
            className="text-white bg-gradient-to-br from-green-400 to-blue-600 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-green-200 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2"
          >
            Add Product
          </button>
        </div>
        <div className="relative overflow-x-auto shadow-md sm:rounded-lg">
          <table className="w-full text-sm text-left rtl:text-right text-gray-500 ">
            <thead className="text-xs text-gray-700 uppercase bg-gray-50">
              <tr>
                <th scope="col" className="px-6 py-3">
                  <div className="flex items-center">
                    Product name
                    <button
                      type="button"
                      onClick={() => {
                        setSort("name");
                        setOrder(order === "asc" ? "desc" : "asc");
                      }}
                    >
                      <svg
                        className="w-3 h-3 ms-1.5"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M8.574 11.024h6.852a2.075 2.075 0 0 0 1.847-1.086 1.9 1.9 0 0 0-.11-1.986L13.736 2.9a2.122 2.122 0 0 0-3.472 0L6.837 7.952a1.9 1.9 0 0 0-.11 1.986 2.074 2.074 0 0 0 1.847 1.086Zm6.852 1.952H8.574a2.072 2.072 0 0 0-1.847 1.087 1.9 1.9 0 0 0 .11 1.985l3.426 5.05a2.123 2.123 0 0 0 3.472 0l3.427-5.05a1.9 1.9 0 0 0 .11-1.985 2.074 2.074 0 0 0-1.846-1.087Z" />
                      </svg>
                    </button>
                  </div>
                </th>
                <th scope="col" className="px-6 py-3">
                  <div className="flex items-center">
                    Category
                    <button
                      type="button"
                      onClick={() => {
                        setSort("category");
                        setOrder(order === "asc" ? "desc" : "asc");
                      }}
                    >
                      <svg
                        className="w-3 h-3 ms-1.5"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M8.574 11.024h6.852a2.075 2.075 0 0 0 1.847-1.086 1.9 1.9 0 0 0-.11-1.986L13.736 2.9a2.122 2.122 0 0 0-3.472 0L6.837 7.952a1.9 1.9 0 0 0-.11 1.986 2.074 2.074 0 0 0 1.847 1.086Zm6.852 1.952H8.574a2.072 2.072 0 0 0-1.847 1.087 1.9 1.9 0 0 0 .11 1.985l3.426 5.05a2.123 2.123 0 0 0 3.472 0l3.427-5.05a1.9 1.9 0 0 0 .11-1.985 2.074 2.074 0 0 0-1.846-1.087Z" />
                      </svg>
                    </button>
                  </div>
                </th>
                <th scope="col" className="px-6 py-3">
                  <div className="flex items-center">
                    Price
                    <button
                      type="button"
                      onClick={() => {
                        setSort("price");
                        setOrder(order === "asc" ? "desc" : "asc");
                      }}
                    >
                      <svg
                        className="w-3 h-3 ms-1.5"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M8.574 11.024h6.852a2.075 2.075 0 0 0 1.847-1.086 1.9 1.9 0 0 0-.11-1.986L13.736 2.9a2.122 2.122 0 0 0-3.472 0L6.837 7.952a1.9 1.9 0 0 0-.11 1.986 2.074 2.074 0 0 0 1.847 1.086Zm6.852 1.952H8.574a2.072 2.072 0 0 0-1.847 1.087 1.9 1.9 0 0 0 .11 1.985l3.426 5.05a2.123 2.123 0 0 0 3.472 0l3.427-5.05a1.9 1.9 0 0 0 .11-1.985 2.074 2.074 0 0 0-1.846-1.087Z" />
                      </svg>
                    </button>
                  </div>
                </th>
                <th scope="col" className="px-6 py-3">
                  <div className="flex items-center">
                    Stock
                    <button
                      type="button"
                      onClick={() => {
                        setSort("stock");
                        setOrder(order === "asc" ? "desc" : "asc");
                      }}
                    >
                      <svg
                        className="w-3 h-3 ms-1.5"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M8.574 11.024h6.852a2.075 2.075 0 0 0 1.847-1.086 1.9 1.9 0 0 0-.11-1.986L13.736 2.9a2.122 2.122 0 0 0-3.472 0L6.837 7.952a1.9 1.9 0 0 0-.11 1.986 2.074 2.074 0 0 0 1.847 1.086Zm6.852 1.952H8.574a2.072 2.072 0 0 0-1.847 1.087 1.9 1.9 0 0 0 .11 1.985l3.426 5.05a2.123 2.123 0 0 0 3.472 0l3.427-5.05a1.9 1.9 0 0 0 .11-1.985 2.074 2.074 0 0 0-1.846-1.087Z" />
                      </svg>
                    </button>
                  </div>
                </th>
                <th scope="col" className="px-6 py-3">
                  <span className="sr-only">Edit</span>
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
                products.map((product, index) => {
                  return (
                    index >= (page - 1) * 10 &&
                    index <= page * 10 - 1 && (
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
                        <td className="px-6 py-4 capitalize">
                          {product.category}
                        </td>
                        <td className="px-6 py-4">{product.price}</td>
                        <td className="px-6 py-4">{product.stock}</td>
                        <td className="flex items-center px-6 py-4">
                          <button
                            type="button"
                            onClick={() =>
                              navigate(`/product/edit/${product.id}`)
                            }
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
                    )
                  );
                })
              )}
            </tbody>
            <tfoot>
              <tr className="font-semibold text-gray-900">
                <th scope="row" className="px-6 py-3 text-base">
                  Total Product :
                </th>
                <td className="px-6 py-3">{products.length} Products</td>
                <td className="px-6 py-3"></td>
                <td className="px-6 py-3"></td>
                {products.length > 10 && (
                  <td className="px-6 py-3 flex text-xl justify-between">
                    <button
                      type="button"
                      // SETTING THE PAGE NUMBER🦺
                      onClick={() => {
                        const prevPage =
                          page === 1
                            ? Math.ceil(products.length / 10)
                            : page - 1;
                        setPage(prevPage);
                        const paramsObject = Object.fromEntries(
                          searchParams.entries()
                        );
                        setSearchParams({ ...paramsObject, page: prevPage });
                      }}
                      // SETTING THE PAGE NUMBER🦺
                    >
                      <FaCircleChevronLeft />
                    </button>
                    <button
                      type="button"
                      // SETTING THE PAGE NUMBER🦺
                      onClick={() => {
                        const nextPage =
                          page === Math.ceil(products.length / 10)
                            ? 1
                            : page + 1;
                        setPage(nextPage);
                        const paramsObject = Object.fromEntries(
                          searchParams.entries()
                        );
                        setSearchParams({ ...paramsObject, page: nextPage });
                        // SETTING THE PAGE NUMBER🦺
                      }}
                    >
                      <FaCircleChevronRight />
                    </button>
                  </td>
                )}
              </tr>
            </tfoot>
          </table>
        </div>
      </main>
    </>
  );
}
