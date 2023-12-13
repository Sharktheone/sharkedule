import {useEffect, useState} from "react"
import {Link, useLoaderData, useNavigate} from "react-router-dom"
import styles from "./styles.module.scss"
import {useDisclosure} from "@/hooks"
import {api} from "@/api/api"
import {notifications} from "@mantine/notifications"
import {IconTrash, IconX} from "@tabler/icons-react"
import {WorkspaceList} from "@kanban/types"
import {Button} from "@/components/ui"
import CreateNewModal from "@/pages/task/createNewModal"


export default function Kanban() {
    let loaderData = useLoaderData()
    let navigate = useNavigate()

    const [workspaces, setWorkspaces] = useState(loaderData as WorkspaceList[])

    useEffect(() => {

        console.log(loaderData)

        setWorkspaces(loaderData as WorkspaceList[])
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
        api.delete(`/kanban/board/${workspaces[0].uuid}/delete`).then(
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

    function Workspaces() {

        if (!workspaces) return null
        if (workspaces.length === 0) return (
            <ul className={styles.workspaces}>
                <li className={styles.noWorkspaces}>
                    <h1 className={styles.dimmed}>No Boards</h1>
                </li>
            </ul>
        )

        return (
            <ul className={styles.workspaces}>
                {workspaces.map((workspace) => (
                    <li key={workspace.uuid}>
                        <h1>{workspace.name}</h1>
                        <Boards workspace={workspace}/>
                    </li>
                ))}
            </ul>
        )
    }

    function Boards({workspace}: { workspace: WorkspaceList }) {
        if (!workspace) return null
        if (!workspace.boards) return null

        if (workspace.boards.length === 0) return (
            <ul className={styles.boards}>
                <li className={styles.noBoards}>
                    <h1 className={styles.dimmed}>No Boards</h1>
                </li>
            </ul>
        )

        return (
            <ul className={styles.boards}>
                {workspace.boards.map((board) => (
                    <li key={board.uuid} className={styles.board}>
                        <Link to={`${workspace.uuid}/${board.uuid}`}>
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
                ))}
            </ul>
        )
    }

    return (
        <div className={styles.dashboard}>
            <div className={styles.boardsHeader}>
                <h1>Your Boards</h1>
                <Button gradient onClick={openNewBoard}>
                    New Board
                </Button>
            </div>
            <CreateNewModal close={close} opened={newOpened} handleCreate={createBoard}/>

            <ul className={styles.workspaces}>
                <Workspaces/>
            </ul>
        </div>
    )
}