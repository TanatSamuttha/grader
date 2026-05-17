import axios from "axios";

export default async function getUserData(setLoggedIn, setUsername, setPhoto){
    try{
        const result = await axios.get("/auth/me");
        setLoggedIn(true);
        setUsername(result.data.username);
        const photoURL = result.data.photoURL;
        setPhoto(photoURL);
        console.log(result.data);
    }
    catch(err){
        setLoggedIn(false);
    }
}