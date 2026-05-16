import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import heroImg from './assets/hero.png'
import './App.css'
import Profile from './components/Profile'

function App() {
    return (
        <>
            <header>
                Grader
                <Profile></Profile>
            </header>
        </>
    )
}

export default App
