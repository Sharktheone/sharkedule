import {Button, CloseButton, Group, Input, Stack, Text, Title} from '@mantine/core'
import {DragDropContext, DragStart, DragUpdate, DropResult} from "react-beautiful-dnd"
import Column from "./column/column"
import {useEffect, useRef, useState} from "react"
import {kanbanBoardType} from "./types"
import {useLoaderData, useNavigate} from "react-router-dom"
import {IconPlus} from "@tabler/icons-react"
import styles from "./styles.module.scss"
import {useStyles} from "./styles"
import {api} from "../../../api/api"
import {notifications} from "@mantine/notifications"
import {ghostType} from "./ghost"

export default function Kanban() {
    const loaderData = useLoaderData()
    const [board, setBoard] = useState<kanbanBoardType>(loaderData as kanbanBoardType)
    const navigate = useNavigate()
    const [isAdding, setIsAdding] = useState(false)
    const newColRef = useRef<HTMLInputElement>(null)
    const [ghost, setGhost] = useState<ghostType | undefined>(undefined)
    const [removeTimeout, setRemoveTimeout] = useState<number | undefined>(undefined)

    const {classes, cx} = useStyles()

    useEffect(() => {
        setBoard(loaderData as kanbanBoardType)
    }, [loaderData])

    function getDraggedElement(draggableId: string) {
        return document.querySelector(`[data-rbd-drag-handle-draggable-id='${draggableId}'] > div > div`)
    }

    function dragStartHandler(event: DragStart) {

        let draggedElement = getDraggedElement(event.draggableId)
        if (!draggedElement) return
        console.log(draggedElement)

        let rect = draggedElement.getBoundingClientRect()

        setGhost({
            height: rect.height,
            index: event.source.index,
            hoveredColumnID: event.source.droppableId,
        })

        console.log("drag start")
    }

    function dragUpdateHandler(event: DragUpdate) {
        let draggedElement = getDraggedElement(event.draggableId)
        if (!draggedElement) return

        let rect = draggedElement.getBoundingClientRect()

        setGhost({
            height: rect.height,
            index: event.destination?.index ?? ghost?.index ?? 0,
            hoveredColumnID: event.destination?.droppableId ?? ghost?.hoveredColumnID ?? "",
        })

    }

    function dragEndHandler(result: DropResult) {
        let {destination, source, draggableId} = result
        console.log(result)
        if (!destination) return
        if (destination.droppableId === source.droppableId && destination.index === source.index) return
        reorderTask(source.droppableId, draggableId, destination.index, destination.droppableId)
        setGhost(undefined)
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
        setIsAdding(true)
    }

    function cancelAddColumn() {
        setRemoveTimeout(
            setTimeout(() => {
                setIsAdding(false)
            }, 100)
        )
    }

    function addColumn() {
        if (removeTimeout) clearTimeout(removeTimeout)

        const name = newColRef.current?.value
        if (!name) {
            notifications.show({title: "Error", message: "Column name cannot be empty", color: "red"})
            return
        }

        api.put(`/kanbanboard/${board.uuid}/column/new`, {name: name}).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: res.data, color: "red"})
                    console.log(res)
                } else {
                    notifications.show({title: "Success", message: "Column created", color: "green"})
                    if (newColRef.current) newColRef.current.value = ""

                    navigate("")
                }
            }).catch(
            (err) => {
                notifications.show({title: "Error", message: err.message, color: "red"})
                console.log(err)
            }
        )
    }

    useEffect(() => {
        newColRef?.current?.focus()
    }, [isAdding])


    return (
        <div className={styles.board}>
            <Title order={1} align="center">{board.name}</Title>
            <Text mb="sm" align="center" color="dimmed">Drag and drop tasks to reorder them</Text>
            <DragDropContext onDragStart={dragStartHandler} onDragEnd={dragEndHandler} onDragUpdate={dragUpdateHandler}>
                <Group className={styles.cols} position="center" align="start" noWrap={true}>
                    {board.columns?.map((column) => (
                        <Column key={column.uuid} column={column} renameColumn={renameColumn} renameTask={renameTask}
                                boardUUID={board.uuid} ghost={ghost}/>
                    ))}

                    {!isAdding ?
                        <>
                            <button onClick={handleNewColumn} className={`${cx(classes.addColumn)} ${styles.footer}`}>
                                <IconPlus size={24}/>
                                <Text align="center">Add a Column</Text>
                            </button>
                        </> :
                        <Stack className={styles.add}>
                            <Input ref={newColRef} onBlur={cancelAddColumn}
                                   placeholder="Column name"></Input>
                            <div className={styles.menu}>
                                <Button onClick={addColumn} gradient={{from: "#6dd6ed", to: "#586bed"}}
                                        variant="gradient">Create</Button>
                                <CloseButton onClick={() => setIsAdding(false)}/>
                            </div>


                        </Stack>

                    }

                </Group>
            </DragDropContext>
        </div>
    )
}

