import {Button, CloseButton, Text, Textarea, Title} from "@mantine/core"
import {useStyles} from "./styles"
import Task from "./task/task"
import {kanbanBoardType, kanbanColumnType} from "../types"
import {Draggable, Droppable} from "react-beautiful-dnd"
import {Dispatch, SetStateAction} from "react"
import styles from "./styles.module.scss"
import {IconPlus, IconTrash} from "@tabler/icons-react"
import {ghostType} from "../ghost"
import {handlers} from "./handlers"

type ColumnProps = {
    column: kanbanColumnType
    boardUUID: string
    index: number
    ghost?: ghostType
    setBoard: Dispatch<SetStateAction<kanbanBoardType>>
    board: kanbanBoardType
}

export default function Column({column, setBoard, board, ghost, index}: ColumnProps) {
    const {classes, cx} = useStyles()

    const h = new handlers(column.uuid, setBoard, board, ghost)

    h.checkGhost()
    h.checkAdding()


    return (
        <Draggable draggableId={column.uuid} index={index}>
            {(provided, snapshot) => (
                <div
                    className={snapshot.isDragging ? styles.dragging : ""}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}

                    ref={provided.innerRef}
                >
                    <Droppable type="task" droppableId={column.uuid} direction="vertical">
                        {(provided) => (
                            <div className={styles.colDrop} {...provided.droppableProps} ref={provided.innerRef}>
                                <div className={`${cx(classes.column)} ${styles.column}`}>
                                    <Title align="left" className={cx(classes.title)} order={3}>
                                        <div>
                                <span onClick={() => h.editText()} contentEditable={h.editable}
                                      onBlur={event => h.handleBlur(event)}>{column.name}</span>
                                            <button onClick={() => h.handleDelete()}>
                                                <IconTrash/>
                                            </button>
                                        </div>
                                    </Title>
                                    <div ref={h.tasksRef}>
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
                                                            <Task key={task.uuid} task={task}
                                                                  renameTask={(uuid, name) => h.renameTask(uuid, name)}
                                                                  boardUUID={board.uuid} columnUUID={column.uuid}/>
                                                        </div>
                                                    </div>
                                                )}
                                            </Draggable>
                                        ))}
                                        {h.ghostElement ?
                                            <div className={`${cx(classes.ghost)} ${styles.ghost}`}
                                                 style={{
                                                     height: h.ghostElement.height,
                                                     top: h.ghostElement.offsetTop
                                                 }}/>
                                            : null
                                        }
                                    </div>

                                    {provided.placeholder}

                                    {h.isAdding ?
                                        <>
                                            <Textarea onBlur={() => h.removeIsAdding()} ref={h.nameRef} autosize
                                                      className={`${cx(classes.add)} ${styles.add}`}
                                                      placeholder="Task name..."/>
                                        </>

                                        : null}

                                    <div className={styles.footer}>
                                        {!h.isAdding ?
                                            <button onClick={() => h.handleNewTask()}>
                                                <IconPlus/>
                                                <Text size="sm"> Add a Task </Text>
                                            </button> :

                                            <div>
                                                <Button variant="gradient" gradient={{from: "#6dd6ed", to: "#586bed"}}
                                                        onClick={() => h.addTask()}> Create </Button>
                                                <CloseButton onClick={() => h.closeIsAdding()}/>
                                            </div>

                                        }

                                    </div>
                                </div>
                            </div>
                        )}
                    </Droppable>
                </div>
            )}
        </Draggable>
    )
}