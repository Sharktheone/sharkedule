import {Flex, Text} from "@mantine/core"
import {useStyles} from "./styles"
import {kanbanTaskType} from "../../types"
import {useRef, useState} from "react"

type TaskProps = {
    task: kanbanTaskType
}


export default function Task({task}: TaskProps) {
    const {classes, cx} = useStyles();
    const [name, setName] = useState(task.name)
    const ref = useRef<HTMLDivElement | null>(null)

    ref?.current?.addEventListener('change', function () {
        console.log(ref?.current?.innerText)
    })

    function editText() {
        ref!.current!.contentEditable = "true"
    }

    function handleChange(e: any) {
        setName(e.target.innerText)
    }

    function handleBlur(e: any) {
        ref!.current!.contentEditable = "false"
        console.log(e.target.innerText)
    }

    return (
        <Flex className={cx(classes.task)}>
            <Text onClick={editText} onInput={handleChange} onBlur={handleBlur} ref={ref} w="100%" h="100%">
                {task.name}
            </Text>
        </Flex>
    )
}