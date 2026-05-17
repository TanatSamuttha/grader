import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import heroImg from './assets/hero.png'
import './App.css'
import Profile from './components/Profile'
import NavigatorBar from './components/NavigatorBar'
import Content from './components/Content'

function App() {
    const [content, setContent] = useState("Problems");

    return (
        <>
            <header>
                Grader
                <Profile />
            </header>
            <main>
                <NavigatorBar setContent = {setContent} />
                <Content content = {content} />
            </main>
        </>
    )
}

export default App
