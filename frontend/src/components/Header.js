import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import './Header.css';

const Header = () => {
  const location = useLocation();

  return (
    <header className="header">
      <nav className="nav">
        <Link to="/" className={location.pathname === '/' ? 'nav-link active' : 'nav-link'}>Главная</Link>
        <Link to="/active" className={location.pathname === '/active' ? 'nav-link active' : 'nav-link'}>Активные проекты</Link>
        <Link to="/archive" className={location.pathname === '/archive' ? 'nav-link active' : 'nav-link'}>Архив проектов</Link>
        <Link to="/profile" className={location.pathname === '/profile' ? 'nav-link active' : 'nav-link'}>Личный аккаунт</Link>
        <Link to="/login" className={location.pathname === '/login' ? 'nav-link active' : 'nav-link'}>Вход/Регистрация</Link>
      </nav>
    </header>
  );
};

export default Header;