import { useState } from "react";
import { signIn } from "../logic/auth.jsx";

export default function Profile(){
    const [isLogedIn, setLogedIn] = useState(false);
    if(isLogedIn) return <><h1>Loged In</h1></>;
    else{
        return (
            <>
                <button onClick={signIn}>Sign In</button>
            </>
        );
    }
}