import { useEffect, useState } from "react";
import { googleSignIn } from "../logic/auth/googleAuth.jsx";

export default function Profile(){
    const [isLogedIn, setLoggedIn] = useState(false);
    const [photoURL, setPhotoURL] = useState("");
    const [username, setUsername] = useState("");

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