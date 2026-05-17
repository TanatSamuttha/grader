export default function getUserData(){
    const result = axios.get("http://localhost:3000/me",
        {
            withCredentials = true
        }
    );
}