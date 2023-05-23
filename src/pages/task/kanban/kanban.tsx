import {Button, CloseButton, Input, Stack, Text, Title} from '@mantine/core'
import {DragDropContext, DragStart, DragUpdate, Droppable, DropResult} from "react-beautiful-dnd"
import Column from "./column/column"
import {useEffect, useRef, useState} from "react"
import {kanbanBoardType} from "./types"
import {useLoaderData, useNavigate} from "react-router-dom"
import {IconPlus} from "@tabler/icons-react"
import styles from "./styles.module.scss"
import {useStyles} from "./styles"
import {api} from "../../../api/api"
import {notifications} from "@mantine/notifications"
import {dragHandlers} from "./dragHandlers"

export default function Kanban() {
    const loaderData = useLoaderData()
    const [board, setBoard] = useState<kanbanBoardType>(loaderData as kanbanBoardType)
    const navigate = useNavigate()
    const [isAdding, setIsAdding] = useState(false)
    const newColRef = useRef<HTMLInputElement>(null)
    const [removeTimeout, setRemoveTimeout] = useState<number | undefined>(undefined)

    const {classes, cx} = useStyles()

    const drag = new dragHandlers(board, setBoard)

    useEffect(() => {
        setBoard(loaderData as kanbanBoardType)
    }, [loaderData])

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
            <DragDropContext onDragStart={drag.Start} onDragEnd={drag.End} onDragUpdate={drag.Update}>
                <Droppable droppableId={board.uuid} type="column" direction="horizontal">
                    {(provided) => (
                        <div
                            ref={provided.innerRef}
                            {...provided.droppableProps}>
                            <div className={styles.cols}>
                                {board.columns?.map((column) => (
                                    <div key={column.uuid}>
                                        <Column column={column} renameColumn={renameColumn}
                                                index={board.columns?.indexOf(column) ?? 0}
                                                renameTask={renameTask}
                                                boardUUID={board.uuid} ghost={drag.ghost}/>
                                    </div>
                                ))}


                                {provided.placeholder}

                                {!isAdding ?
                                    <>
                                        <button onClick={handleNewColumn}
                                                className={`${cx(classes.addColumn)} ${styles.footer}`}>
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

                            </div>
                        </div>
                    )

                    }

                </Droppable>

            </DragDropContext>
        </div>
    )
}

