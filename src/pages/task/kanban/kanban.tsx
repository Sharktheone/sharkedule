import {Group, Text, Title} from '@mantine/core'
import {DragDropContext, DropResult} from "react-beautiful-dnd"
import Column from "./column/column"
import {useEffect, useState} from "react"
import {kanbanBoardType} from "./types"
import {useLoaderData, useNavigate} from "react-router-dom"
import {IconPlus} from "@tabler/icons-react";
import styles from "./styles.module.scss"
import {useStyles} from "./styles";
import {api} from "../../../api/api";
import {notifications} from "@mantine/notifications";
import NewColumnModal from "./column/NewColumnModal";
import {useDisclosure} from "@mantine/hooks";

export default function Kanban() {
    const loaderData = useLoaderData()
    const [board, setBoard] = useState<kanbanBoardType>(loaderData as kanbanBoardType)
    const [newColumnOpened, {open, close}] = useDisclosure(false)
    const navigate = useNavigate()

    const {classes, cx} = useStyles()

    useEffect(() => {
        setBoard(loaderData as kanbanBoardType)
    }, [loaderData])

    function dragEndHandler(result: DropResult) {
        let {destination, source, draggableId} = result
        console.log(result)
        if (!destination) return
        if (destination.droppableId === source.droppableId && destination.index === source.index) return
        reorderTask(source.droppableId, draggableId, destination.index, destination.droppableId)
    }

    function reorderTask(fromColumn: string, uuid: string, to: number, toColumn: string,) {
        let newBoard = {...board}

        let fromColumnIndex = newBoard?.columns?.findIndex((column) => column.uuid === fromColumn)
        let toColumnIndex = newBoard?.columns?.findIndex((column) => column.uuid === toColumn)
        let taskIndex = newBoard?.columns[fromColumnIndex]?.tasks?.findIndex((task) => task.uuid === uuid)
        let [task] = newBoard?.columns[fromColumnIndex]?.tasks?.splice(taskIndex, 1)


        newBoard?.columns[toColumnIndex]?.tasks?.splice(to, 0, task)
        setBoard(newBoard)
    }

    function renameTask(uuid: string, name: string) {
        let newBoard = {...board}
        newBoard.columns?.forEach((column) => {
            column.tasks?.forEach((task) => {
                if (task.uuid === uuid) {
                    task.name = name
                    return
                }
            })
        })
        setBoard(newBoard)
    }

    function renameColumn(uuid: string, name: string) {
        let newBoard = {...board}
        newBoard.columns?.forEach((column) => {
            if (column.uuid === uuid) {
                column.name = name
                return
            }
        })
        setBoard(newBoard)
    }

    function handleNewColumn() {
        open()
    }

    function addColumn(name: string) {
        api.put(`/kanbanboard/${board.uuid}/column/new`, {name: name}).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: res.data, color: "red"})
                    console.log(res)
                } else {
                    notifications.show({title: "Success", message: "Column created", color: "green"})
                    navigate("")
                }
            }).catch(
            (err) => {
                notifications.show({title: "Error", message: err.message, color: "red"})
                console.log(err)
            }
        )

    }


    return (
        <div className={styles.board}>
            <Title order={1} align="center">{board.name}</Title>
            <Text mb="sm" align="center" color="dimmed">Drag and drop tasks to reorder them</Text>
            <DragDropContext onDragEnd={dragEndHandler}>
                <Group className={styles.cols} position="center" align="start" noWrap={true}>
                    {board.columns?.map((column) => (
                        <Column key={column.uuid} column={column} renameColumn={renameColumn} renameTask={renameTask}
                                boardUUID={board.uuid}/>
                    ))}
                    <button onClick={handleNewColumn} className={`${cx(classes.addColumn)} ${styles.footer}`}>
                        <IconPlus size={24}/>
                        <Text align="center">Add a Column</Text>
                    </button>
                </Group>
            </DragDropContext>
            <NewColumnModal close={close} opened={newColumnOpened} addColumn={addColumn}/>
        </div>
    )
}

