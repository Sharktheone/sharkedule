import {Button, CloseButton, Text, Textarea, Title} from "@mantine/core"
import {useStyles} from "./styles"
import Task from "./task/task"
import {Draggable, Droppable} from "react-beautiful-dnd"
import {Dispatch, SetStateAction, useContext, useMemo} from "react"
import styles from "./styles.module.scss"
import {IconPlus, IconTrash} from "@tabler/icons-react"
import {ghostType} from "../ghost"
import {handlers} from "./handlers"
import {Column, environment} from "@kanban/types2"
import {EnvironmentContext} from "@kanban/environment"

type ColumnProps = {
    column: string
    boardUUID: string
    ghost?: ghostType
}

export default function Column({column, ghost, boardUUID}: ColumnProps) {
    const {classes, cx} = useStyles()

    const {environment, setEnvironment} = useContext(EnvironmentContext)

    const h = new handlers(column, boardUUID, setEnvironment, environment, ghost)

    h.checkGhost()
    h.checkAdding()

    function getIndex() {
        const board = environment.boards?.find(b => b.uuid === boardUUID)
        return board?.columns?.findIndex(c => c === column) ?? 0
    }

    function getColumn() {
        return useMemo(() => {
            return environment.columns?.find(c => c.uuid === column) ?? {} as Column
        }, [environment, column])

    }

    return (
        <Draggable draggableId={column} index={getIndex()}>
            {(provided, snapshot) => (
                <div
                    className={snapshot.isDragging ? styles.dragging : ""}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}

                    ref={provided.innerRef}
                >
                    <Droppable type="task" droppableId={column} direction="vertical">
                        {(provided) => (
                            <div className={styles.colDrop} {...provided.droppableProps} ref={provided.innerRef}>
                                <div className={`${cx(classes.column)} ${styles.column}`}>
                                    <Title align="left" className={cx(classes.title)} order={3}>
                                        <div>
                                            <span onClick={() => h.editText()} contentEditable={h.editable}
                                                  onBlur={event => h.handleBlur(event)}>{getColumn().name}
                                            </span>
                                            <button onClick={() => h.handleDelete()}>
                                                <IconTrash/>
                                            </button>
                                        </div>
                                    </Title>
                                    <div ref={h.tasksRef}>
                                        {getColumn().tasks?.map((task, index) => (
                                            <Draggable key={task} draggableId={task} index={index}>
                                                {(provided, snapshot) => (
                                                    <div
                                                        className={snapshot.isDragging ? styles.dragging : ""}
                                                        {...provided.draggableProps}
                                                        {...provided.dragHandleProps}

                                                        ref={provided.innerRef}
                                                    >

                                                        <div style={{paddingBottom: "0.625rem"}}>
                                                            <Task key={task} task={task}
                                                                  renameTask={(uuid, name) => h.renameTask(uuid, name)}
                                                                  board={boardUUID} column={column}
                                                            />
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