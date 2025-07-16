import LoginForm from "../components/LoginForm";

function LoginPage() {
  const handleLoginSuccess = () => {
    window.location.href = "/dashboard"; // chuyển sang trang chính sau khi login
  };

  return <LoginForm onLogin={handleLoginSuccess} />;
}

export default LoginPage;
