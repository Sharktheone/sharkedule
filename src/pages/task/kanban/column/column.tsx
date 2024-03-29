import {Button, CloseButton, Text, Textarea, Title} from "@/components/ui"
import Task from "@/pages/task/kanban/column/task/task"
import {Draggable, Droppable} from "@hello-pangea/dnd"
import {useContext, useMemo} from "react"
import styles from "./styles.module.scss"
import {IconPlus, IconTrash} from "@tabler/icons-react"
import {ghostType} from "@/pages/task/kanban/ghost"
import {handlers} from "@/pages/task/kanban/column/handlers"
import {Column as Col} from "@kanban/types"
import {EnvironmentContext} from "@kanban/environment"


type ColumnProps = {
    column: string
    boardUUID: string
    ghost?: ghostType
}

export default function Column({column, ghost, boardUUID}: ColumnProps) {
    const {environment, setEnvironment} = useContext(EnvironmentContext)
    const h = new handlers(column, boardUUID, setEnvironment, environment, ghost)

    h.checkGhost()
    h.checkAdding()

    function getIndex() {
        const board = environment?.boards?.find(b => b.uuid === boardUUID)
        return board?.columns?.findIndex(c => c === column) ?? -1
    }

    function getColumn() {
        return useMemo(() => {
            return environment?.columns?.find(c => c.uuid === column) ?? {} as Col
        }, [environment, column])
    }

    return (
        <Draggable draggableId={column} index={getIndex()}>
            {(provided, snapshot) => (
                <div
                    className={snapshot.isDragging ? styles.dragging : undefined}
                    {...provided.draggableProps}
                    {...provided.dragHandleProps}

                    ref={provided.innerRef}
                >
                    <div className={styles.colDrop}>
                        <div className={styles.column}>
                            <Title s={3}>
                                <div>
                                    <span onClick={() => h.editText()}
                                          contentEditable={h.editable} // TODO: dont use contentEditable
                                          onBlur={event => h.handleBlur(event)}>
                                        {getColumn().name as string}
                                    </span>
                                    <button onClick={() => h.handleDelete()}>
                                        <IconTrash/>
                                    </button>
                                </div>
                            </Title>
                            <Droppable type="task" droppableId={column} direction="vertical">
                                {(provided, snapshot) => (
                                    <div ref={provided.innerRef} {...provided.droppableProps}>
                                        <div ref={h.tasksRef}>
                                            {getColumn().tasks?.map((task, index) => (
                                                <Draggable key={task} draggableId={task} index={index}>
                                                    {(provided, snapshot) => (
                                                        <div
                                                            className={`${styles.taskwrapper} ${snapshot.isDragging ? styles.dragging : undefined}`}
                                                            {...provided.draggableProps}
                                                            {...provided.dragHandleProps}

                                                            ref={provided.innerRef}
                                                        >

                                                            <div>
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
                                                <div className={styles.ghost}
                                                     style={{
                                                         height: h.ghostElement.height,
                                                         top: h.ghostElement.offsetTop
                                                     }}/>
                                                : null
                                            }
                                            {provided.placeholder}
                                        </div>
                                    </div>
                                )}
                            </Droppable>
                            {h.isAdding ?
                                <>
                                    {/*TODO: close the create shit, when user clicks out (onBlur not optimal)*/}
                                    <Textarea onBlur={() => h.removeIsAdding()} autosize className={styles.add}
                                              textareaRef={h.nameRef}
                                              onKeyDown={(e) => {
                                                  if (e.key === "Enter" && !e.shiftKey) {
                                                      h.addTask()
                                                  }
                                              }}
                                              placeholder="Task name..."/>
                                </>

                                : null}
                            <div className={styles.footer}>
                                {!h.isAdding ?
                                    <button onClick={() => h.handleNewTask()}>
                                        <IconPlus/>
                                        <Text s="medium"> Add a Task </Text>
                                    </button> :

                                    <div>
                                        <Button variant="primary" gradient
                                                onClick={() => h.addTask()}> Create </Button>
                                        <CloseButton onClick={() => h.closeIsAdding()}/>
                                    </div>
                                }
                            </div>
                        </div>
                    </div>
                </div>
            )}
        </Draggable>
    )
}