import {useEffect, useState} from "react"
import {Link, useLoaderData, useNavigate} from "react-router-dom"
import {Button, Container, Title} from "@mantine/core"
import {useColors} from "./styles"
import styles from "./styles.module.scss"
import CreateNewModal from "@/pages/task/createNewModal"
import {useDisclosure} from "@mantine/hooks"
import {api} from "@/api/api"
import {notifications} from "@mantine/notifications"
import {IconTrash, IconX} from "@tabler/icons-react"
import {NameList} from "@kanban/types2"


export default function KanbanBoards() {
    let loaderData = useLoaderData()
    let navigate = useNavigate()

    let {classes, cx} = useColors()

    const [nameList, setNameList] = useState(loaderData as NameList[])

    useEffect(() => {
        setNameList(loaderData as NameList[])
    }, [loaderData])

    const [newOpened, {open, close},] = useDisclosure(false)

    function openNewBoard() {
        open()
    }

    function createBoard(name: string, description: string) {
        api.put("/kanban/board/new", {name: name, description: description}).then(
            (res) => {
                if (res.status > 300) {
                    console.log(res)
                    notifications.show({
                        title: "Error",
                        message: res.data.message ?? "Unknown Error",
                        color: "red",
                        icon: <IconX/>
                    })
                } else {
                    notifications.show({title: "Success", message: "Board created", color: "green"})

                    navigate(`${res.data.uuid}`)
                }

            }).catch(e => {
            notifications.show({title: "Error", message: e.message, color: "red", icon: <IconX/>})
        })
    }

    function deleteBoard() {
        api.delete(`/kanban/board/${nameList[0].uuid}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({
                        title: "Error",
                        message: res.data.message ?? "Unknown Error",
                        color: "red",
                        icon: <IconX/>
                    })
                } else {
                    notifications.show({title: "Success", message: "Deleted Board", color: "green"})
                    navigate("")
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
                    nameList.length > 0 ?
                        <>
                            {nameList.map((board) => (
                                <li>
                                    <Link to={board.uuid}>
                                        {board.name}
                                    </Link>
                                    <div>
                                        <div>
                                            <button onClick={deleteBoard}>
                                                <IconTrash/>
                                            </button>
                                        </div>
                                    </div>

                                </li>
                            ))
                            }
                        </>
                        : <li className="no-boards">
                            <Title color="dimmed">No Boards</Title>
                        </li>
                }

            </ul>
        </Container>
    )
}