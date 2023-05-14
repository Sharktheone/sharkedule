import {Title} from "@mantine/core"
import {useStyles} from "./styles"
import Task from "./task/task"
import {kanbanColumnType} from "../types"
import {Draggable, Droppable} from "react-beautiful-dnd"
import {useState} from "react"

type ColumnProps = {
    column: kanbanColumnType
    renameColumn: (uuid: string, name: string) => void
    renameTask: (uuid: string, name: string) => void
}

export default function Column({column, renameColumn,  renameTask}: ColumnProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)

    function editText() {
        setEditable(true)
    }

    function handleBlur(e: any) {
        setEditable(false)
        renameColumn(column.uuid, e.target.innerText)
    }


    return (
        <div className={cx(classes.column)}>
            <Title align="left" className={cx(classes.title)} order={3}>
                <div onClick={editText} contentEditable={editable} onBlur={handleBlur}>{column.name}</div>
            </Title>
            <Droppable droppableId={column.uuid} direction="vertical">
                {(provided) => (
                    <div {...provided.droppableProps} ref={provided.innerRef}>
                            {column.tasks?.map((task, index) => (
                                <Draggable key={task.uuid} draggableId={task.uuid} index={index}>
                                    {(provided) => (
                                        <div
                                            // className={cx(classes.task, {[classes.taskDragging]: snapshot.isDragging})}
                                            {...provided.draggableProps}
                                            {...provided.dragHandleProps}

                                            ref={provided.innerRef}
                                        >

                                            <div style={{paddingBottom: "0.625rem"}}>
                                                <Task key={task.uuid} task={task} renameTask={renameTask}/>
                                            </div>
                                        </div>
                                    )}
                                </Draggable>
                            ))}
                        {provided.placeholder}
                    </div>
                )}
            </Droppable>
        </div>
    )
}