import {useState} from "react"
import {kanbanBoardType} from "./kanban/types"
import {Link, useLoaderData} from "react-router-dom"
import {Container, Title} from "@mantine/core"
import {useColors} from "./styles"
import styles from "./styles.module.scss"


export default function KanbanBoards() {
    let loaderData = useLoaderData()

    let {classes, cx} = useColors()

    const [boardNames, setBoardNames] = useState(loaderData as kanbanBoardType[])


    return (
        <Container className={`${styles.boards} ${cx(classes.colors)}`}>
            <Title>Your Boards</Title>

            <ul>
                {boardNames.map((board) => (
                    <li>
                        <Link to={board.uuid}>{board.name}</Link>
                    </li>
                ))
                }
            </ul>
        </Container>
    )
}