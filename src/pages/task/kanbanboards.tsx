import {useState} from "react"
import {kanbanBoardType} from "./kanban/types"
import {Link, useLoaderData} from "react-router-dom"
import {Button, Container, Title} from "@mantine/core"
import {useColors} from "./styles"
import styles from "./styles.module.scss"
import CreateNewModal from "./createNewModal"
import {useDisclosure} from "@mantine/hooks"


export default function KanbanBoards() {
    let loaderData = useLoaderData()

    let {classes, cx} = useColors()

    const [boardNames, setBoardNames] = useState(loaderData as kanbanBoardType[])
    const [newOpened, {open, close}, ] = useDisclosure(false)

    function newBoard() {
        open()
    }

    return (
        <Container className={`${styles.boards} ${cx(classes.colors)}`}>
            <Title>Your Boards</Title>

            <Button variant="gradient" gradient={{from: "yellow", to: "red"}} onClick={newBoard}>New Board</Button>
            <CreateNewModal close={close} opened={newOpened}/>

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