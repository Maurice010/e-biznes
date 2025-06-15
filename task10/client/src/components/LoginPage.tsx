import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await fetch("task10-back-grfchafkdxbxd8bb.polandcentral-01.azurewebsites.net/api/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });
      const data = await res.json();
      if (res.ok) {
        localStorage.setItem("token", data.token);
        navigate("/");
      } else {
        alert(data.error || "Login failed");
      }
    } catch {
      alert("Network error");
    }
  };

  const handleGoogle = () => {
    window.location.href = "task10-back-grfchafkdxbxd8bb.polandcentral-01.azurewebsites.net/api/auth/google";
  };

  const handleGithub = () => {
    window.location.href = "task10-back-grfchafkdxbxd8bb.polandcentral-01.azurewebsites.net/api/auth/github";
  };

  return (
    <div style={container}>
      <h2>Logowanie</h2>
      <form onSubmit={handleLogin} style={formStyle}>
        <input type="email" placeholder="Email" value={email} onChange={e => setEmail(e.target.value)} required />
        <input type="password" placeholder="Hasło" value={password} onChange={e => setPassword(e.target.value)} required />
        <button type="submit">Zaloguj</button>
      </form>
      <button onClick={handleGoogle}>Zaloguj przez Google</button>
      <button onClick={handleGithub}>Zaloguj przez GitHub</button>
      <p>
        Nie masz konta? <a href="/register">Zarejestruj się</a>
      </p>
      
    </div>
  );
}

const container = {
  display: "flex",
  flexDirection: "column" as const,
  alignItems: "center",
  justifyContent: "center",
  height: "100vh",
  gap: "1rem",
};

const formStyle = {
  display: "flex",
  flexDirection: "column" as const,
  gap: "0.5rem",
  width: "200px",
};
