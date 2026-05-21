import { useState } from 'react'

import './App.css'

import Profile from './components/Profile'
import NavigatorBar from './components/NavigatorBar'
import LobbyContent from './components/LobbyContent'

const useStore = create((set) => ({
    main: (
        <NavigatorBar setContent={setContent} />,
        <LobbyContent content={content} setMain = {setMain} />
    ),
    setMain: (main) => set({main})
}));

function App() {
    const [content]
    return (
        <>
            <header>
                Grader
                <Profile />
            </header>
            <main>
                {main}
            </main>
        </>
    );

    return null;
}

export default App;