import {Flex, Text} from "@mantine/core"
import {useStyles} from "./styles"
import {kanbanTaskType} from "../../types"
import {useRef, useState} from "react"

type TaskProps = {
    task: kanbanTaskType
    renameTask: (uuid: string, name: string) => void
}


export default function Task({task, renameTask}: TaskProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)

    function editText() {
        setEditable(true)
    }

    function handleBlur(e: any) {
        setEditable(false)
        renameTask(task.uuid, e.target.innerText)
    }

    return (
        <Flex className={cx(classes.task)}>
            <Text onClick={editText} onBlur={handleBlur} contentEditable={editable} w="100%" h="100%">
                {task.name}
            </Text>
        </Flex>
    )
}