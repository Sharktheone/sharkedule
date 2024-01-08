import {useEffect, useState} from "react"
import {Link, useLoaderData, useNavigate} from "react-router-dom"
import styles from "./styles.module.scss"
import {useDisclosure} from "@/hooks"
import {api} from "@/api/api"
import {IconArrowBadgeDown, IconPlus, IconTrash, IconX} from "@tabler/icons-react"
import {NameList, WorkspaceList} from "@kanban/types"
import {Button, Text, Title} from "@/components/ui"
import CreateNewModal from "@/pages/task/createNewModal"
import {toast} from "react-toastify"


export default function Kanban() {
    let loaderData = useLoaderData()
    let navigate = useNavigate()

    const [workspaces, setWorkspaces] = useState(loaderData as WorkspaceList[])
    const [collapsed, setCollapsed] = useState({} as { [key: string]: boolean })
    const [newBoardWorkspace, setNewBoardWorkspace] = useState<NameList | undefined>(undefined)

    useEffect(() => {

        console.log(loaderData)

        setWorkspaces(loaderData as WorkspaceList[])
    }, [loaderData])

    const [newOpened, {open, close}] = useDisclosure(false)

    function openNewBoard(ws: NameList) {
        setNewBoardWorkspace(ws)
        open()
    }

    function openNewWorkspace() {

    }

    function createBoard(workspace: string, name: string, description: string) {
        api.put(`/${workspace}/kanban/board/new`, {name: name, description: description}).then(
            (res) => {
                if (res.status > 300) {
                    console.log(res)

                    toast(`Error creating board: ${res.data.message}`, {icon: <IconX/>, type: "error"})
                } else {
                    toast("Board created", {type: "success"})

                    navigate(`${res.data.uuid}`)
                }

            }).catch(e => {
            toast(`Error creating board: ${e}`, {icon: <IconX/>, type: "error"})
        })
    }

    function deleteBoard(workspace: string, board: string) {
        api.delete(`/${workspace}/kanban/board/${workspace}/${board}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    toast(`Error deleting board: ${res.data.message}`, {icon: <IconX/>, type: "error"})
                } else {
                    toast("Deleted Board", {type: "success"})
                    navigate(".", {replace: true})
                }
            }).catch(e => {
            toast(`Error deleting board: ${e}`, {icon: <IconX/>, type: "error"})
        })
    }

    function deleteWorkspace(workspace: string) {
        api.delete(`/${workspace}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    toast(`Error deleting workspace: ${res.data.message}`, {icon: <IconX/>, type: "error"})
                } else {
                    toast("Deleted Workspace", {type: "success"})
                    navigate(".", {replace: true})
                }
            }).catch(e => {
            toast(`Error deleting workspace: ${e}`, {icon: <IconX/>, type: "error"})
        })
    }


    function toggleCollapsed(workspace: string) {
        setCollapsed(prev => {
            const newState = {...prev}
            newState[workspace] = !newState[workspace]
            return newState
        })
    }

    function Workspaces() {
        if (!workspaces) return null
        if (workspaces.length === 0) return (
            <ul className={styles.workspaces}>
                <li className={styles.noWorkspaces}>
                    <Text s={2} dimmed>No Boards</Text>
                </li>
            </ul>
        )

        return (
            <ul className={styles.workspaces}>
                {workspaces.map((workspace) => (
                    <li className={styles.workspace} key={workspace.uuid}>
                        <div>
                            <div className={styles.workspaceName}>
                                <button onClick={() => toggleCollapsed(workspace.uuid)}>
                                    <IconArrowBadgeDown
                                        className={collapsed[workspace.uuid] ? styles.collapsed : undefined}/>
                                </button>
                                <Text c="white" w="bold" a="left" s={4}>{workspace.name}</Text>
                            </div>
                            <div className={styles.workspaceHovermenu}>
                                {/*TODO: don't use a hovermenu but a button which opens a list of options*/}
                                <div>
                                    <button onClick={() => openNewBoard(workspace)}>
                                        <IconPlus/>
                                    </button>
                                </div>
                                <div className={styles.delete}>
                                    <button onClick={() => deleteWorkspace(workspace.uuid)}>
                                        <IconTrash/>
                                    </button>
                                </div>
                            </div>
                        </div>
                        <Boards collapsed={collapsed[workspace.uuid]} workspace={workspace}/>

                    </li>
                ))}
            </ul>
        )
    }

    function Boards({workspace, collapsed}: { workspace: WorkspaceList, collapsed: boolean }) {
        if (!workspace) return null

        if (!workspace.boards || workspace.boards.length === 0) return (
            <ul className={`${styles.boards} ${collapsed ? styles.hide : undefined}`}>
                <li className={styles.noBoards}>
                    <Text className={styles.dimmed}>No boards created yet... 😥</Text>
                </li>
            </ul>
        )

        return (
            <ul className={`${styles.boards} ${collapsed ? styles.hide : undefined}`}>
                {workspace.boards.map((board) => (
                    <li key={board.uuid} className={styles.board}>
                        <Link to={`../board/${workspace.uuid}/${board.uuid}`}>
                            {board.name}
                        </Link>
                        <div className={styles.boardHovermenu}>
                            <div className={styles.delete}>
                                <button onClick={() => deleteBoard(workspace.uuid, board.uuid)}>
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
                <Title>Your Boards</Title>
                <Button gradient>
                    New Workspace
                </Button>
            </div>
            <CreateNewModal close={close} opened={newOpened} workspace={newBoardWorkspace}
                            handleCreate={(workspace, name, description) => createBoard(workspace, name, description)}/> {/*TODO: make workspace selectable*/}

            <Workspaces/>
        </div>
    )
}