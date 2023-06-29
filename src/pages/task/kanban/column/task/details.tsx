import Details from "@/components/details/details"
import {Dispatch, SetStateAction, useState} from "react"
import {Task} from "@kanban/types2"
import {getTask} from "@/pages/task/utils/task"

type Props = {
    open: boolean
    setOpen: Dispatch<SetStateAction<boolean>>
    uuid: string
}

export default function TaskDetails({open, setOpen, uuid}: Props) {

    const [task, setTask] = useState<Task>(getTask(uuid) ?? {} as Task)


    function onClose() {
        setOpen(false)

    }
    return (
        <Details open={open} onClose={onClose} title={task.name}>
            <div>
                hello
            </div>
        </Details>
    )
}