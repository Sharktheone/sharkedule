import {Button, CloseButton, Text, Textarea, Title} from "@mantine/core"
import {useStyles} from "./styles"
import Task from "./task/task"
import {kanbanColumnType} from "../types"
import {Draggable, Droppable} from "react-beautiful-dnd"
import {useEffect, useRef, useState} from "react"
import styles from "./styles.module.scss"
import {IconPlus, IconTrash, IconX} from "@tabler/icons-react"
import {api} from "../../../../api/api"
import {notifications} from "@mantine/notifications"
import {useNavigate} from "react-router-dom"
import {ghostElementType, ghostType} from "../ghost"

type ColumnProps = {
    column: kanbanColumnType
    renameColumn: (uuid: string, name: string) => void
    renameTask: (uuid: string, name: string) => void
    boardUUID: string
    ghost?: ghostType
}

export default function Column({column, renameColumn, renameTask, boardUUID, ghost}: ColumnProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)
    const [isAdding, setIsAdding] = useState(false)
    const navigate = useNavigate()
    const nameRef = useRef<HTMLTextAreaElement>(null)
    const tasksRef = useRef<HTMLDivElement>(null)
    const [ghostElement, setGhostElement] = useState<ghostElementType | undefined>()
    const [removeTimeout, setRemoveTimeout] = useState<number | undefined>(undefined)

    function editText() {
        setEditable(true)
    }

    function handleBlur(e: any) {
        setEditable(false)
        renameColumn(column.uuid, e.target.innerText)
    }

    function handleDelete() {
        api.delete(`/kanbanboard/${boardUUID}/column/${column.uuid}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                } else {
                    notifications.show({title: "Success", message: "Deleted Column", color: "green"})
                    navigate("")

                }
            }
        )
    }

    useEffect(() => {
        if (!ghost) {
            setGhostElement(undefined)
            return
        }
        if (ghost.hoveredColumnID !== column.uuid) {
            setGhostElement(undefined)
            return
        }

        let offset = 0

        let tasks = [].slice.call(tasksRef.current?.children) as HTMLDivElement[]

        console.log(tasks.filter(task => !task.className.includes(styles.ghost)))

        tasks.filter(task => !task.className.includes(styles.ghost)).forEach((task, index) => {
            if (index < ghost.index) {
                offset += task.getBoundingClientRect().height
            }
        })

        const ghostElement = {
            height: ghost.height + "px",
            offsetTop: offset + "px",
        }

        console.log(ghostElement)

        setGhostElement(ghostElement)


    }, [ghost])


    function handleNewTask() {
        setIsAdding(true)
    }

    useEffect(() => {
        if (isAdding) {
            nameRef.current?.focus()
        }
    }, [isAdding])


    function addTask() {
        if (removeTimeout) clearTimeout(removeTimeout)

        let name: string
        if (nameRef.current?.value) {
            name = nameRef.current?.value
        } else {
            notifications.show({title: "Error", message: "Task name cannot be empty", color: "red", icon: <IconX/>})
            return
        }

        api.put(`/kanbanboard/${boardUUID}/column/${column.uuid}/task/new`, {name: name}).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                } else {
                    renameTask(res.data.uuid, name)
                    if (nameRef.current) {
                        nameRef.current.value = ""
                        nameRef.current.focus()
                    }
                    navigate("")
                }

            }).catch(e => {
            notifications.show({title: "Error", message: e.message, color: "red", icon: <IconX/>})
        })
    }


    function removeIsAdding() {
        if (removeTimeout) clearTimeout(removeTimeout)
        setRemoveTimeout(setTimeout(() => {
            setIsAdding(false)
        }, 100))
    }

    return (
        <Droppable droppableId={column.uuid} direction="vertical">
            {(provided) => (
                <div className={styles.colDrop} {...provided.droppableProps} ref={provided.innerRef}>
                    <div className={`${cx(classes.column)} ${styles.column}`}>
                        <Title align="left" className={cx(classes.title)} order={3}>
                            <div>
                                <span onClick={editText} contentEditable={editable}
                                      onBlur={handleBlur}>{column.name}</span>
                                <button onClick={handleDelete}>
                                    <IconTrash/>
                                </button>
                            </div>
                        </Title>
                        <div ref={tasksRef}>
                            {column.tasks?.map((task, index) => (
                                <Draggable key={task.uuid} draggableId={task.uuid} index={index}>
                                    {(provided, snapshot) => (
                                        <div
                                            className={snapshot.isDragging ? styles.dragging : ""}
                                            {...provided.draggableProps}
                                            {...provided.dragHandleProps}

                                            ref={provided.innerRef}
                                        >

                                            <div style={{paddingBottom: "0.625rem"}}>
                                                <Task key={task.uuid} task={task} renameTask={renameTask}
                                                      boardUUID={boardUUID} columnUUID={column.uuid}/>
                                            </div>
                                        </div>
                                    )}
                                </Draggable>
                            ))}
                            {ghostElement ?
                                <div className={`${cx(classes.ghost)} ${styles.ghost}`}
                                     style={{height: ghostElement.height, top: ghostElement.offsetTop}}/>
                                : null
                            }
                        </div>

                        {provided.placeholder}

                        {isAdding ?
                            <>
                                <Textarea onBlur={removeIsAdding} ref={nameRef} autosize
                                          className={`${cx(classes.add)} ${styles.add}`}
                                          placeholder="Task name..."/>
                            </>

                            : null}

                        <div className={styles.footer}>
                            {!isAdding ?
                                <button onClick={handleNewTask}>
                                    <IconPlus/>
                                    <Text size="sm"> Add a Task </Text>
                                </button> :

                                <div>
                                    <Button variant="gradient" gradient={{from: "#6dd6ed", to: "#586bed"}}
                                            onClick={addTask}> Create </Button>
                                    <CloseButton onClick={() => setIsAdding(false)}/>
                                </div>

                            }

                        </div>
                    </div>
                </div>
            )}
        </Droppable>
    )
}