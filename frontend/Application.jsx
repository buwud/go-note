import React from 'react'
import { createRoot } from 'react-dom/client';

function Application() {
    return (
        <div>Hello World!</div>
    )
}

// Clear the existing HTML content
document.body.innerHTML = '<div id="app"></div>';

const root = document.getElementById('app');
if (root) {
    createRoot(root).render(<Application />);
}
else {
    throw new Error('Root element not found');
}
