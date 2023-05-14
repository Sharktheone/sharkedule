import {Flex, Text} from "@mantine/core"
import {useStyles} from "./styles"
import {kanbanTaskType} from "../../types"
import {useRef, useState} from "react"

type TaskProps = {
    task: kanbanTaskType
}


export default function Task({task}: TaskProps) {
    const {classes, cx} = useStyles()
    const [editable, setEditable] = useState(false)

    function editText() {
        setEditable(true)
    }

    function handleBlur(e: any) {
        setEditable(false)
        console.log(e.target.innerText)
    }

    return (
        <Flex className={cx(classes.task)}>
            <Text onClick={editText} onBlur={handleBlur} contentEditable={editable} w="100%" h="100%">
                {task.name}
            </Text>
        </Flex>
    )
}