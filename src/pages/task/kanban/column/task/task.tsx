import {Flex} from "@mantine/core"
import {useStyles} from "./styles"
import {kanbanTaskType} from "../../types"

type TaskProps = {
    task: kanbanTaskType
}


export default function Task({task}: TaskProps) {
    const {classes, cx} = useStyles();

    return (
        <Flex className={cx(classes.task)}>
            {task.name}
        </Flex>
    )
}