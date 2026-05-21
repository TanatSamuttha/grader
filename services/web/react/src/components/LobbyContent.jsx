import Problems from "./Problems";

export default function LobbyContent({ content, setMain }) {
    if (content === "Problems") {
        return <Problems visibility="public" guild="" setMain={setMain} />;
    } 
    else if (content === "Contest") {
        return <h1>Contest</h1>;
    }

    return null;
}