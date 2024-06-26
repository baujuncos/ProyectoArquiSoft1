import React, { useState } from 'react';
import { useCookies } from 'react-cookie';

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [cookies, setCookie] = useCookies(['token', 'userId']);
    const [loading, setLoading] = useState(false); // Estado para el loading
    const [error, setError] = useState(''); // Estado para manejar errores

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true); // Comienza el loading
        setError(''); // Resetea el error

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ username, password })
            });

            if (!response.ok) {
                throw new Error('Login failed');
            }

            const data = await response.json();
            const { id_usuario, token } = data;

            // Guardar en localStorage
            localStorage.setItem('userId', id_usuario);
            localStorage.setItem('token', token);

            // Guardar en cookies
            setCookie('token', token, { path: '/' });
            setCookie('userId', id_usuario, { path: '/' });

            // Redirigir después del login exitoso
            window.location.href = '/';
        } catch (error) {
            console.error('Error:', error);
            setError('Login failed. Please check your credentials and try again.');
        } finally {
            setLoading(false); // Termina el loading
        }
    };

    return (
        <div className="login-container">
            <h2>Login</h2>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="username">Username:</label>
                    <input
                        type="text"
                        id="username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="password">Contraseña:</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                {error && <p style={{ color: 'red' }}>{error}</p>}
                <button type="submit" disabled={loading}>
                    {loading ? 'Logging in...' : 'Login'}
                </button>
            </form>
        </div>
    );
};

export default Login;
