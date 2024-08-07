import React from 'react'
import { createRoot } from 'react-dom/client';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Home from './pages/home/Home';
import About from './pages/about/About';
import './pages/home/Home.css'
import Login from './pages/login/Login';

function Application() {
    return (
        <div>Hello World!</div>
    )
}

// Clear the existing HTML content
document.body.innerHTML = '<div id="app"></div>';

const root = document.getElementById('app');
if (root) {
    createRoot(root).render(
        <BrowserRouter>
            <Routes>
                <Route index element={<Home />} />
                <Route path="/about" element={<About />} />
                <Route path="/login" element={<Login />} />
                <Route path="/*" element={<Home />} />


            </Routes>
        </BrowserRouter>
    );
}
else {
    throw new Error('Root element not found');
}
