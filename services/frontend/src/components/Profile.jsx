import { cache, useEffect, useState } from "react";
import { googleSignIn } from "../logic/auth/googleAuth.jsx";
import getUserData from "../logic/auth/UserData.jsx";

export default function Profile(){
    const [isLogedIn, setLoggedIn] = useState(false);
    const [photo, setPhoto] = useState("");
    const [username, setUsername] = useState("");

    useEffect(() => {
        async function loadUser() {
            await getUserData(setLoggedIn, setUsername, setPhoto);
        }

        loadUser();
    }, []);

    if (isLogedIn) {
        return (
            <>
                <h3>{username}</h3>
                <img src={photo} alt="profile"/>
            </>
        );
    }
    else {
        return (
            <>
                <button onClick={() => {googleSignIn(setLoggedIn, setUsername, setPhoto)}}>Sign In</button>
            </>
        );
    }
}