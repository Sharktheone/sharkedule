import {Title} from "@mantine/core"
import {useStyles} from "./styles"
import Task from "./task/task"
import {kanbanColumnType} from "../types"
import {Draggable, Droppable} from "react-beautiful-dnd"

type ColumnProps = {
    column: kanbanColumnType
}

export default function Column({column}: ColumnProps) {
    const {classes, cx} = useStyles()
    return (
        <div className={cx(classes.column)}>
            <Title align="left" className={cx(classes.title)} order={3}>{column.name}</Title>
            <Droppable droppableId={column.uuid} direction="vertical">
                {(provided) => (
                    <div {...provided.droppableProps} ref={provided.innerRef}>
                        {column.tasks?.map((task, index) => (
                            <Draggable key={task.uuid} draggableId={task.uuid} index={index}>
                                {(provided, snapshot) => (
                                    <div
                                        // className={cx(classes.task, {[classes.taskDragging]: snapshot.isDragging})}
                                        {...provided.draggableProps}
                                        {...provided.dragHandleProps}

                                     ref={provided.innerRef}
                                    >
                                        <Task key={task.uuid} task={task}/>
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