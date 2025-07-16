import { useState } from "react";

function LoginForm({ onLogin }) {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();
      if (data.token) {
        localStorage.setItem("token", data.token); // lưu JWT
        onLogin(); // chuyển trang
      } else {
        alert("Sai tài khoản hoặc mật khẩu");
      }
    } catch (err) {
      alert("Lỗi kết nối máy chủ");
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="bg-white p-6 rounded shadow-md max-w-sm mx-auto mt-10"
    >
      <h2 className="text-xl font-semibold mb-4 text-center">Đăng nhập</h2>
      <input
        className="w-full mb-3 border p-2 rounded"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="password"
        className="w-full mb-3 border p-2 rounded"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button className="bg-blue-500 text-white px-4 py-2 rounded w-full">
        Đăng nhập
      </button>

      <p className="text-center mt-4">I haven't account?{" "}
        <a href="/register" className="text-blue-600 hover:underline">Register</a>
      </p>
      
    </form>
  );
}

export default LoginForm;
