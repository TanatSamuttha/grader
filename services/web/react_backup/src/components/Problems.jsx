import axios from "axios";
import { useEffect, useState } from "react";

export default function Problems({ visibility, guild, setMain }) {
    const [problems, setProblems] = useState([]);

    useEffect(() => {
        async function fetchProblems() {
            try {
                if (visibility === "public") {
                    const response = await axios.get("/problem/public");
                    setProblems(response.data.problems);
                    console.log(response);
                }
            } catch (err) {
                console.error(err);
            }
        }

        fetchProblems();
    }, [visibility]);

    return (
        <ul>
            {problems.map((problem) => (
                <li key={problem.problem_id} onClick={() => openProblem(problem.problem_id, setMain)}>
                    {problem.name}
                </li>
            ))}
        </ul>
    );
}

function openProblem(problem_id, setMain){
    setMain("");
}