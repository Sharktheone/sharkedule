import {Title} from "@mantine/core"
import {useStyles} from "./styles";
import Task from "./task/task"
import {kanbanColumnType} from "../types"

type ColumnProps = {
    column: kanbanColumnType
}

export default function Column({column}: ColumnProps) {
    const {classes, cx} = useStyles();
    return (
        <div className={cx(classes.column)}>
            <Title align="left" className={cx(classes.title)} order={3}>{column.name}</Title>
            <div>
                {column.tasks?.map((task) => (
                    <Task key={task.uuid} task={task}/>
                ))}
            </div>
        </div>
    )
}