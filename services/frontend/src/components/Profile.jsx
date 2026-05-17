import { useEffect, useState } from "react";
import { googleSignIn } from "../logic/auth/googleAuth.jsx";

export default function Profile(){
    const [isLogedIn, setLoggedIn] = useState(false);
    const [photoURL, setPhotoURL] = useState("");
    const [username, setUsername] = useState("");

    useEffect(() => {
        const cookies = document.cookie.split("; ");

        let photoURL = "";
        let username = "";

        for (let cookie of cookies) {
            const [key, value] = cookie.split("=");

            if (key === "Username" && value != "") {
                username = value;
            }
            else if (key === "PhotoURL" && value != "") {
                photoURL = value;
            }
        }

        if (username != "") {
            setUsername(username);
            setLoggedIn(true);
        }
        if (photoURL != "") {
            setPhotoURL(photoURL);
        }
    }, []);

    console.log(photoURL);

    if (isLogedIn) {
        return (
            <>
                <h3>{username}</h3>
                <img src={photoURL} alt="profile"/>
            </>
        );
    }
    else {
        return (
            <>
                <button onClick={() => {googleSignIn(setLoggedIn)}}>Sign In</button>
            </>
        );
    }
}