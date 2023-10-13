import {Link} from "react-router-dom"

export default function Home() {
    return (
        <div>
            <Link to="/kanbanboard">Tasks</Link>
        </div>
    )
}