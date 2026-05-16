import { useEffect, useState } from "react";
import { googleSignIn } from "../logic/auth/googleAuth.jsx";

export default function Profile(){
    const [isLogedIn, setLoggedIn] = useState(false);

    useEffect(() => {
        const cookies = document.cookie.split("; ");

        for (let cookie of cookies) {
            const [key, value] = cookie.split("=");

            if (key === "IsAuthenticated" && value === "true") {
                setLoggedIn(true);
                break;
            }
        }
    }, []);

    if(isLogedIn) return <><h1>Loged In</h1></>;
    else{
        return (
            <>
                <button onClick={() => {googleSignIn(setLogedIn)}}>Sign In</button>
            </>
        );
    }
}