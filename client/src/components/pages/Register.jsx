import { MdLock } from "react-icons/md";
import { SlLogin } from "react-icons/sl";
import { IoClose } from "react-icons/io5";
import { Link } from "react-router-dom";

export default function Register() {
  return (
    <div className="min-h-screen bg-[url(/login-bg.jpg)] flex">
      <div
        className="relative p-8 items-center flex flex-col rounded-3xl mx-auto my-auto w-1/2 bg-gradient-to-b from-cyan-200 to-white via-white to-80% "
        style={{ boxShadow: "0 0 8px #fff" }}
      >
        <Link
          to={"/"}
          className="absolute top-4 right-4 text-gray-500 text-xl"
        >
          <IoClose />
        </Link>
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

        <form className="w-full mx-auto mt-7 flex flex-col gap-4">
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
              id="username"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Username"
              name="username"
            />
          </div>
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
              id="fullname"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Fullname"
              name="fullname"
            />
          </div>
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
              name="email-address"
            />
          </div>
          <div className="relative w-full font-lora">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none text-gray-500 text-xl -ml-0.5">
              <MdLock />
            </div>
            <input
              type="password"
              id="password"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Password"
              name="password"
            />
          </div>
          <div className="relative w-full font-lora">
            <div className="absolute inset-y-0 start-0 flex items-center ps-3.5 pointer-events-none text-gray-500 text-xl -ml-0.5">
              <MdLock />
            </div>
            <input
              type="password"
              id="password-confirm"
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-cust-red focus:border-cust-red active:ring-cust-red active:border-cust-red block w-full ps-10 p-2.5"
              placeholder="Konfirmasi Password"
              name="password-confirm"
            />
          </div>
          <button
            type="button"
            className="text-white font-poppins mt-5 bg-gradient-to-br from-cyan-500 via-blue-500 to-cyan-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-cyan-200 dark:focus:ring-cyan-800 font-medium rounded-lg px-5 py-2.5 text-center"
          >
            Masuk Sekarang
          </button>
          <p className="font-poppins text-[0.8rem] font-light text-zinc-400 text-center">
            Sudah punya akun?{" "}
            <Link to={"/auth/login"} className="font-semibold text-black">Masuk</Link> di sini.
          </p>
          {/* 
          <p className="text-[0.8rem] font-poppins text-center text-zinc-400 before:content-['—————'] before:tracking-[-0.15em] before:mr-5 after:content-['—————'] after:tracking-[-0.15em] after:ml-5">
            Or sign in with it.
          </p>
          <div className="flex gap-3 justify-stretch">
            <button
              type="button"
              onClick={() => signIn("google")}
              className="flex justify-center w-full py-2 text-2xl border border-zinc-200 rounded-xl cursor-pointer active:translate-y-0.5 active:shadow-none"
              style={{ boxShadow: "0 3px 3px #ddd" }}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                width="26"
                height="26"
                viewBox="0 0 48 48"
              >
                <path
                  fill="#FFC107"
                  d="M43.611,20.083H42V20H24v8h11.303c-1.649,4.657-6.08,8-11.303,8c-6.627,0-12-5.373-12-12c0-6.627,5.373-12,12-12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C12.955,4,4,12.955,4,24c0,11.045,8.955,20,20,20c11.045,0,20-8.955,20-20C44,22.659,43.862,21.35,43.611,20.083z"
                ></path>
                <path
                  fill="#FF3D00"
                  d="M6.306,14.691l6.571,4.819C14.655,15.108,18.961,12,24,12c3.059,0,5.842,1.154,7.961,3.039l5.657-5.657C34.046,6.053,29.268,4,24,4C16.318,4,9.656,8.337,6.306,14.691z"
                ></path>
                <path
                  fill="#4CAF50"
                  d="M24,44c5.166,0,9.86-1.977,13.409-5.192l-6.19-5.238C29.211,35.091,26.715,36,24,36c-5.202,0-9.619-3.317-11.283-7.946l-6.522,5.025C9.505,39.556,16.227,44,24,44z"
                ></path>
                <path
                  fill="#1976D2"
                  d="M43.611,20.083H42V20H24v8h11.303c-0.792,2.237-2.231,4.166-4.087,5.571c0.001-0.001,0.002-0.001,0.003-0.002l6.19,5.238C36.971,39.205,44,34,44,24C44,22.659,43.862,21.35,43.611,20.083z"
                ></path>
              </svg>
            </button>
            <div
              className="flex justify-center w-full py-2 text-2xl border border-zinc-200 rounded-xl cursor-pointer active:translate-y-0.5 active:shadow-none"
              style={{ boxShadow: "0 3px 3px #ddd" }}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                x="0px"
                y="0px"
                width="26"
                height="26"
                viewBox="0 0 48 48"
              >
                <path
                  fill="#039be5"
                  d="M24 5A19 19 0 1 0 24 43A19 19 0 1 0 24 5Z"
                ></path>
                <path
                  fill="#fff"
                  d="M26.572,29.036h4.917l0.772-4.995h-5.69v-2.73c0-2.075,0.678-3.915,2.619-3.915h3.119v-4.359c-0.548-0.074-1.707-0.236-3.897-0.236c-4.573,0-7.254,2.415-7.254,7.917v3.323h-4.701v4.995h4.701v13.729C22.089,42.905,23.032,43,24,43c0.875,0,1.729-0.08,2.572-0.194V29.036z"
                ></path>
              </svg>
            </div>
            <div
              className="flex justify-center w-full py-2 text-2xl border border-zinc-200 rounded-xl cursor-pointer active:translate-y-0.5 active:shadow-none"
              style={{ boxShadow: "0 3px 3px #ddd" }}
            >
              <FaApple />
            </div>
          </div> */}
        </form>
      </div>
    </div>
  )
}
