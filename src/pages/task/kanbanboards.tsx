import {useState} from "react"
import {kanbanBoardType} from "./kanban/types"
import {Link, useLoaderData, useNavigate} from "react-router-dom"
import {Button, Container, Title} from "@mantine/core"
import {useColors} from "./styles"
import styles from "./styles.module.scss"
import CreateNewModal from "./createNewModal"
import {useDisclosure} from "@mantine/hooks"
import {api} from "../../api/api"
import {notifications} from "@mantine/notifications"
import {IconX} from "@tabler/icons-react"


export default function KanbanBoards() {
    let loaderData = useLoaderData()
    let navigate = useNavigate()

    let {classes, cx} = useColors()

    const [boardNames, setBoardNames] = useState(loaderData as kanbanBoardType[])
    const [newOpened, {open, close},] = useDisclosure(false)

    function openNewBoard() {
        open()
    }

    function createBoard(name: string) {
        console.log(name)
        api.put("/kanbanboard/new", {name: name}).then(
            (res) => {
                if (res.status > 300) {
                    console.log(res)
                    console.log("hello1")
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                    console.log("hello")
                } else {
                    notifications.show({title: "Success", message: "Board created", color: "green"})

                    navigate(`${res.data.uuid}`)
                }

            }).catch(e => {
                notifications.show({title: "Error", message: e.message, color: "red", icon: <IconX/>})
        })
    }

    return (
        <Container className={`${styles.boards} ${cx(classes.colors)}`}>
            <div>
                <Title>Your Boards</Title>
                <Button variant="gradient" gradient={{from: "yellow", to: "red"}} onClick={openNewBoard}>New
                    Board</Button>
            </div>
            <CreateNewModal close={close} opened={newOpened} handleCreate={createBoard}/>

            <ul>
                {
                    boardNames.length > 0 ?
                        <>
                            {boardNames.map((board) => (
                                <li>
                                    <Link to={board.uuid}>{board.name}</Link>
                                </li>
                            ))
                            }
                        </>
                        : <li>
                            <Title color="dimmed">No Boards</Title>
                        </li>
                }

            </ul>
        </Container>
    )
}