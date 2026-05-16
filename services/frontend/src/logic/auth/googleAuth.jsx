import {initializeApp} from "https://www.gstatic.com/firebasejs/11.6.1/firebase-app.js";
import {
    getAuth,
    GoogleAuthProvider,
    signInWithPopup,
    signOut
} from "https://www.gstatic.com/firebasejs/11.6.1/firebase-auth.js";
import firebaseConfig from "../../config/firebaseConfig.jsx";
import axios from "axios";

const app = initializeApp(firebaseConfig);
const auth = getAuth(app);
const provider = new GoogleAuthProvider();

export async function googleSignIn() {
    const result = await signInWithPopup(auth, provider);
    const user = result.user;
    const token = await user.getIdToken();
    const response = await axios.post("http://localhost:3000/authen/google", {
            token: token
        },
        {
            withCredentials: true
        }
    );
    console.log(response.data);
}