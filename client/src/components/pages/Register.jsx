import { MdLock } from "react-icons/md";
import { SlLogin } from "react-icons/sl";
import { Link } from "react-router-dom";
import { z } from "zod";
import { useNavigate } from "react-router-dom";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";
import { IoMdPerson } from "react-icons/io";
import { BsFillPersonVcardFill } from "react-icons/bs";

const postFormSchema = z.object({
  username: z.string(),
  fullname: z.string(),
  email: z.string().email(),
  password: z.string().min(8, "Password must be at least 8 characters"),
});

export default function Register() {
  const navigate = useNavigate();

  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");

  const { register, handleSubmit, formState, watch } = useForm({
    resolver: zodResolver(postFormSchema),
  });

  const onSubmit = handleSubmit(async (data) => {
    if (watch("password") !== confirmPassword) {
      setError("Password and confirm password must be the same");
      return;
    }
    const res = await fetch("http://localhost:8080/v1/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    }).then((res) => res.json());

    if (res.status) {
      navigate("/auth/login");
    } else {
      setError(res.message);
    }
  });

  return (
    <div className="min-h-screen bg-[url(/login-bg.jpg)] flex p-8">
      <div
        className="relative p-8 items-center flex flex-col rounded-3xl mx-auto my-auto w-1/2 bg-gradient-to-b from-cyan-200 to-white via-white to-80% "
        style={{ boxShadow: "0 0 8px #fff" }}
      >
        <div className="text-2xl p-4 rounded-2xl bg-zinc-100 aspect-square shadow-xl">
          <SlLogin />
        </div>
        <div className="font-poppins flex flex-col items-center mt-4">
          <h1 className="text-2xl font-semibold text-center">
            Daftar akun baru
          </h1>
          <h2 className="text-center font-light text-[0.8rem] w-80 text-zinc-700">
            Daftar akun baru agar bisa mendapatkan akses ke fitur-fitur kami.
          </h2>
        </div>

        <form
          onSubmit={onSubmit}
          className="w-full mx-auto mt-7 flex flex-col gap-4"
        >
          <div className="relative w-full font-lora">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none">
              <div className="w-4 h-4 text-gray-500 dark:text-gray-400">
                <BsFillPersonVcardFill />
              </div>
            </div>
            <input
              type="text"
              id="username"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Username"
              {...register("username")}
            />
          </div>
          {formState.errors.username && (
            <p className="text-sm text-red-400 font-semibold -mt-4">
              {formState.errors.username.message}
            </p>
          )}
          <div className="relative w-full font-lora">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none">
              <div className="w-4 h-4 text-gray-500 dark:text-gray-400">
                <IoMdPerson />
              </div>
            </div>
            <input
              type="text"
              id="fullname"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Fullname"
              {...register("fullname")}
            />
          </div>
          {formState.errors.fullname && (
            <p className="text-sm text-red-400 font-semibold -mt-4">
              {formState.errors.fullname.message}
            </p>
          )}
          <div className="relative w-full font-lora">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none">
              <svg
                className="w-4 h-4 text-gray-500 dark:text-gray-400"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="currentColor"
                viewBox="0 0 20 16"
              >
                <path d="m10.036 8.278 9.258-7.79A1.979 1.979 0 0 0 18 0H2A1.987 1.987 0 0 0 .641.541l9.395 7.737Z" />
                <path d="M11.241 9.817c-.36.275-.801.425-1.255.427-.428 0-.845-.138-1.187-.395L0 2.6V14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V2.5l-8.759 7.317Z" />
              </svg>
            </div>
            <input
              type="text"
              id="email-address"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Email"
              {...register("email")}
            />
          </div>
          {formState.errors.email && (
            <p className="text-sm text-red-400 font-semibold -mt-4">
              {formState.errors.email.message}
            </p>
          )}
          <div className="relative w-full font-lora">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none text-gray-500 text-xl -ml-0.5">
              <MdLock />
            </div>
            <input
              type="password"
              id="password"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Password"
              {...register("password")}
            />
          </div>
          {formState.errors.password && (
            <p className="text-sm text-red-400 font-semibold -mt-4">
              {formState.errors.password.message}
            </p>
          )}
          <div className="relative w-full font-lora">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none text-gray-500 text-xl -ml-0.5">
              <MdLock />
            </div>
            <input
              onChange={(e) => setConfirmPassword(e.target.value)}
              type="password"
              id="password-confirm"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Konfirmasi Password"
            />
          </div>
          {error && (
            <h1 className="text-sm text-red-400 font-semibold">{error}</h1>
          )}
          <button
            type="submit"
            className="text-white font-poppins mt-5 bg-gradient-to-br from-cyan-500 via-blue-500 to-cyan-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-cyan-200 dark:focus:ring-cyan-800 font-medium rounded-lg px-5 py-2.5 text-center"
          >
            Daftar Sekarang
          </button>
          <p className="font-poppins text-[0.8rem] font-light text-zinc-400 text-center">
            Sudah punya akun?{" "}
            <Link to={"/auth/login"} className="font-semibold text-black">
              Masuk
            </Link>{" "}
            di sini.
          </p>
        </form>
      </div>
    </div>
  );
}
