import React from 'react'
import { Link } from 'react-router-dom'


export default function Home() {
    return (
        <div className="home-container">
            <header className="home-header">
                <h1 className="home-title">Go Note</h1>
                <nav>
                    <Link to="/about" className="home-link">About</Link>
                    <Link to="/login" className="home-link">Login</Link>
                </nav>
            </header>
            <main className="home-main">
                <h2 className="home-subtitle">Your Notes, Organized</h2>
                <p className="home-description">
                    Welcome to Go Note! Keep all your notes in one place, easily accessible and beautifully organized.
                </p>
                <Link to="/login" className="home-button">Get Started</Link>
            </main>
        </div>
    )
}
