import {Text, Title} from "@mantine/core"
import {useStyles} from "./styles"
import Task from "./task/task"
import {kanbanColumnType} from "../types"
import {Draggable, Droppable} from "react-beautiful-dnd"
import {useState} from "react"
import styles from "./styles.module.scss"
import {IconPlus, IconTrash, IconX} from "@tabler/icons-react";
import {api} from "../../../../api/api";
import {notifications} from "@mantine/notifications";
import {useNavigate} from "react-router-dom";
import NewTaskModal from "./task/NewTaskModal";
import {useDisclosure} from "@mantine/hooks";

type ColumnProps = {
    column: kanbanColumnType
    renameColumn: (uuid: string, name: string) => void
    renameTask: (uuid: string, name: string) => void
    boardUUID: string
}

export default function Column({column, renameColumn, renameTask, boardUUID}: ColumnProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)
    const [newTaskOpened, {open, close}] = useDisclosure(false)
    const navigate = useNavigate()

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

    function handleNewTask() {
        open()
    }

    function addTask(name: string) {
        api.put(`/kanbanboard/${boardUUID}/column/${column.uuid}/task/new`, {name: name}).then(
            (res) => {
                if (res.status > 300) {
                    notifications.show({title: "Error", message: "res.data", color: "red", icon: <IconX/>})
                } else {
                    renameTask(res.data.uuid, name)
                    close()
                    navigate("")
                }

            }).catch(e => {
            notifications.show({title: "Error", message: e.message, color: "red", icon: <IconX/>})
        })
    }


    return (
        <div className={`${cx(classes.column)} ${styles.column}`}>
            <Title align="left" className={cx(classes.title)} order={3}>
                <div>
                    <span onClick={editText} contentEditable={editable} onBlur={handleBlur}>{column.name}</span>
                    <button onClick={handleDelete}>
                        <IconTrash/>
                    </button>
                </div>
            </Title>
            <Droppable droppableId={column.uuid} direction="vertical">
                {(provided) => (
                    <div {...provided.droppableProps} ref={provided.innerRef}>
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
                                            <Task key={task.uuid} task={task} renameTask={renameTask} boardUUID={boardUUID} columnUUID={column.uuid}/>
                                        </div>
                                    </div>
                                )}
                            </Draggable>
                        ))}
                        {provided.placeholder}
                    </div>
                )}
            </Droppable>

            <div className={styles.footer}>
                <button onClick={handleNewTask}>
                    <IconPlus/>
                    <Text size="sm"> Add a Task </Text>
                </button>
            </div>
            <NewTaskModal close={close} opened={newTaskOpened} addTask={addTask}/>
        </div>
    )
}