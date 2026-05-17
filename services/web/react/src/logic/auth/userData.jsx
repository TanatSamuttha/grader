import axios from "axios";

export default async function getUserData(setLoggedIn, setUsername, setPhoto){
    try{
        const result = await axios.get("http://localhost:3000/me",
            {
                withCredentials: true
            }
        );
        setLoggedIn(true);
        setUsername(result.data.username);
        const photoURL = result.data.photoURL;
        setPhoto(photoURL);
    }
    catch(err){
        setLoggedIn(false);
    }
}