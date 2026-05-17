import { useState } from "react";

export default function Content({content}){
    if (content === "Problems") return <h1>Problems</h1>;
    else if (content === "Contest") return <h1>Contest</h1>;
}