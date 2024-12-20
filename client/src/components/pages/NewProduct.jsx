import React from "react";
import Sidebar from "../layouts/Sidebar";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import { jwtDecode } from "jwt-decode";
import Cookies from "js-cookie";
import Unauthorized from "./Unauthorized";
import { MdOutlineKeyboardBackspace } from "react-icons/md";

const postFormSchema = z.object({
  name: z.string(),
  category: z.string(),
  price: z.string(),
  stock: z.string(),
  description: z.string(),
});

export default function NewProduct() {
  // Check if user is admin🐳
  const [isAdmin, setIsAdmin] = useState(() => {
    const { role } = Cookies.get("token")
      ? jwtDecode(Cookies.get("token"))
      : { role: "" };
    return role === "admin" ? true : false;
  });

  if (!isAdmin) return <Unauthorized />;

  // Check if user is admin🐳

  // CREATE request to create a new product🐳
  const navigate = useNavigate();
  const [error, setError] = useState("");
  const { register, handleSubmit, formState } = useForm({
    resolver: zodResolver(postFormSchema),
  });

  const onSubmit = handleSubmit(async (data) => {
    const numericPrice = parseFloat(data.price);
    const numericStock = parseFloat(data.stock);
    const res = await fetch("http://localhost:8080/v1/products", {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        ...data,
        price: numericPrice,
        stock: numericStock,
      }),
    }).then((res) => res.json());

    if (res.status) {
      navigate("/product");
    } else {
      setError(res.message);
    }
  });
  // CREATE request to create a new product🐳

  return (
    <>
      <aside>
        <Sidebar></Sidebar>
      </aside>
      <main className="ml-72 p-10">
        <div className="flex justify-between items-center">
          <h1 className="text-2xl font-bold pb-2 mb-4 border-b-2">
            New Product
          </h1>
          <button
            onClick={() => navigate("/product")}
            type="button"
            className="text-red-500 hover:text-white border border-red-500 hover:bg-red-500 focus:ring-4 focus:outline-none focus:ring-red-300 font-medium rounded-lg px-3 py-1.5 text-center me-2 mb-2 text-xl"
          >
            <MdOutlineKeyboardBackspace />
          </button>
        </div>
        <form onSubmit={onSubmit} className="max-w-md">
          <div className="mb-4">
            <label
              htmlFor="name"
              className="block mb-2 text-sm font-medium text-gray-900"
            >
              Product name
            </label>
            <input
              type="text"
              id="name"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
              placeholder="Lenovo Thinkpad X1 Carbon"
              required
              {...register("name")}
            />
            {formState.errors.name && (
              <span className="text-red-400 text-sm">
                {formState.errors.name.message}
              </span>
            )}
          </div>
          <div className="w-full mb-4">
            <label
              htmlFor="category"
              className="block mb-2 text-sm font-medium text-gray-900"
            >
              Category
            </label>
            <input
              type="text"
              id="category"
              aria-describedby="helper-text-explanation"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
              placeholder="Laptop"
              required
              {...register("category")}
            />
            {formState.errors.category && (
              <span className="text-red-400 text-sm">
                {formState.errors.category.message}
              </span>
            )}
          </div>
          <div className="mb-4 flex gap-4">
            <div className="w-full">
              <label
                htmlFor="price"
                className="block mb-2 text-sm font-medium text-gray-900"
              >
                Price
              </label>
              <input
                type="number"
                id="price"
                aria-describedby="helper-text-explanation"
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
                placeholder="4800000"
                required
                {...register("price")}
              />
              {formState.errors.price && (
                <span className="text-red-400 text-sm">
                  {formState.errors.price.message}
                </span>
              )}
            </div>
            <div className="w-full">
              <label
                htmlFor="stock"
                className="block mb-2 text-sm font-medium text-gray-900"
              >
                Stock
              </label>
              <input
                type="number"
                id="stock"
                aria-describedby="helper-text-explanation"
                className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
                placeholder="4800"
                required
                {...register("stock")}
              />
              {formState.errors.stock && (
                <span className="text-red-400 text-sm">
                  {formState.errors.stock.message}
                </span>
              )}
            </div>
          </div>
          <div>
            <label
              htmlFor="description"
              className="block mb-2 text-sm font-medium text-gray-900"
            >
              Product description
            </label>
            <textarea
              id="description"
              rows="4"
              className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"
              placeholder="Product description here..."
              required
              {...register("description")}
            ></textarea>
            {formState.errors.description && (
              <span className="text-red-400 text-sm">
                {formState.errors.description.message}
              </span>
            )}
          </div>
          {error && (
            <h1 className="text-sm text-red-400 font-semibold">{error}</h1>
          )}
          <button
            type="submit"
            className="mt-8 w-full text-white bg-gradient-to-r from-blue-500 via-blue-600 to-blue-700 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 shadow-lg shadow-blue-500/50 dark:shadow-lg dark:shadow-blue-800/80 font-medium rounded-lg text-sm px-5 py-2.5 text-center me-2 mb-2 "
          >
            Add Product
          </button>
        </form>
      </main>
    </>
  );
}
