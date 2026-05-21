import { useState } from 'react'

import './App.css'

import Profile from './components/Profile'
import NavigatorBar from './components/NavigatorBar'
import LobbyContent from './components/LobbyContent'

function App() {
    const [main, setMain] = useState("Lobby");
    const [content, setContent] = useState("Problems");

    if (main === "Lobby") {
        return (
            <>
                <header>
                    Grader
                    <Profile />
                </header>

                <main>
                    <NavigatorBar setContent={setContent} />
                    <LobbyContent content={content} setMain = {setMain} />
                </main>
            </>
        )
    }

    return null;
}

export default App;