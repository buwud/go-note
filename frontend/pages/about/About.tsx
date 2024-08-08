import React from 'react';
import './About.css';
import { Link } from 'react-router-dom'

export default function About() {
    return (
        <div className="about-page">
            <header className="navbar">
                <h1 className="home-title">Go Note</h1>
                <nav>
                    <ul className="navbar-menu">
                        <li><a href="/" className="navbar-link">Home</a></li>
                        <li><a href="/about" className="navbar-link">About</a></li>
                        <li><a href="/help" className="navbar-link">Help</a></li>
                    </ul>
                </nav>
            </header>
            <div className="about-container">
                <h1>About Us</h1>
                <p>This is the About page of our website. Here we describe the purpose of our application and how it can benefit users.</p>
                <footer>
                    <p>Created by buwu | Hacked by Tahinli</p>
                </footer>
            </div>
        </div>
    );
}
