import { useState, useEffect } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import './SignupForm.css'; // Link to the CSS file

const SignupForm = () => {
    const [form, setForm] = useState({ username: '', password: '', role: 'user' });
    const navigate = useNavigate();

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
        const res = await fetch(`${API_BASE}/api/auth/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(form)
        });

        if (res.ok) {
            alert('Signup successful!');
            navigate('/login');
        } else {
            const data = await res.json();
            alert(data.error || 'Signup failed');
        }
    };

    return (
        <div className="signup-container">
            <div className="signup-box">
                <h2 className="signup-title">Create Your Account</h2>
                <form onSubmit={handleSubmit} className="signup-form">
                    <input
                        name="username"
                        placeholder="Username"
                        onChange={handleChange}
                        required
                        className="signup-input"
                    />
                    <input
                        name="password"
                        type="password"
                        placeholder="Password"
                        onChange={handleChange}
                        required
                        className="signup-input"
                    />
                    <select
                        name="role"
                        onChange={handleChange}
                        className="signup-select"
                    >
                        <option value="user">User</option>
                        <option value="admin">Admin</option>
                    </select>
                    <button type="submit" className="signup-button">Register</button>
                </form>
                <p className="login-link">
                    Already have an account? <Link to="/login">Login here</Link>
                </p>
            </div>
        </div>
    );
};

export default SignupForm;
