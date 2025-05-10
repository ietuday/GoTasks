import { useState, useEffect } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import './LoginForm.css';

const LoginForm = () => {
  const [form, setForm] = useState({ username: '', password: '' });
  const navigate = useNavigate();

  // 🧠 Redirect if token already exists
  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      navigate('/tasks');
    }
  }, [navigate]);

  const handleChange = e => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async e => {
    e.preventDefault();
    const API_BASE = 'http://localhost:8080';
    const res = await fetch(`${API_BASE}/api/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(form)
    });

    const data = await res.json();

    if (res.ok && data.token) {
      localStorage.setItem('token', data.token);
      alert('Login successful!');
      navigate('/tasks');
    } else {
      alert(data.error || 'Login failed');
    }
  };

  return (
    <div className="login-container">
      <div className="login-box">
        <h2 className="login-title">Login</h2>
        <form onSubmit={handleSubmit} className="login-form">
          <input name="username" onChange={handleChange} placeholder="Username" required className="login-input" />
          <input name="password" type="password" onChange={handleChange} placeholder="Password" required className="login-input" />
          <button type="submit" className="login-button">Login</button>
        </form>
        <p className="signup-link">
          Don’t have an account? <Link to="/signup">Sign up here</Link>
        </p>
      </div>
    </div>
  );
};

export default LoginForm;
