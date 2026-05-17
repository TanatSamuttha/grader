export default function NavigatorBar({setContent}) {
    return (
        <ul>
            <li onClick={() => {setContent("Problems")}}> Problems </li>
            <li onClick={() => {setContent("Contest")}}> Contest </li>
        </ul>
    )
}