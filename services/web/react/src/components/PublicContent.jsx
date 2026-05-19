import Problems from "./Problems";

export default function PublicContent({ content }) {
    if (content === "Problems") {
        return <Problems visibility="public" />;
    } 
    else if (content === "Contest") {
        return <h1>Contest</h1>;
    }

    return null;
}