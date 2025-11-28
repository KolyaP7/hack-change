import React from 'react';

const Login = () => {
  return (
    <div>
      <h1>Регистрация и вход</h1>
      <form>
        <input type="email" placeholder="Email" />
        <input type="password" placeholder="Пароль" />
        <button type="submit">Войти</button>
        <button type="button">Регистрация</button>
      </form>
    </div>
  );
};

export default Login;