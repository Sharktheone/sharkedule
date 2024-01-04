import {Text} from "@/components/ui"
import React, {useContext, useState} from "react"
import styles from "./styles.module.scss"
import {IconCircleCheck, IconTrash, IconX} from "@tabler/icons-react"
import {api} from "@/api/api"
import {useNavigate} from "react-router-dom"
import {SlotProvider} from "@kanban/column/task/slots/slotProvider"
import UpperSlot from "@kanban/column/task/slots/upper/upperSlot"
import LowerSlot from "@kanban/column/task/slots/lower/lowerSlot"
import {EnvironmentContext} from "@kanban/environment"
import TaskDetails from "@kanban/column/task/details/details"
import {useDoubleClick} from "@/hooks"
import {toast} from "react-toastify"

type TaskProps = {
    board: string
    column: string
    task: string
    renameTask: (uuid: string, name: string) => void
}


export default function Task({task, renameTask, board, column}: TaskProps) {
    const [editable, setEditable] = useState(false)
    const navigate = useNavigate()
    const {environment, setEnvironment} = useContext(EnvironmentContext)
    const [taskDetails, setTaskDetails] = useState(false)

    const clickHandler = useDoubleClick(openDetails, editText)


    const [t, setT] = useState(getTask(task))

    const ref = React.useRef<HTMLDivElement>(null)


    function getTask(uuid: string) {
        return environment?.tasks?.find((task) => task.uuid === uuid)
    }

    function editText() {
        ref?.current?.focus()
        // setEditable(true)
    }

    function handleBlur(e: any) {
        setEditable(false)
        renameTask(task, e.target.innerText)
    }

    function handleDelete() {
        api.delete(`/${environment.workspace}/kanban/board/${board}/column/${column}/task/${task}/delete`).then(
            (res) => {
                if (res.status > 300) {
                    toast("Error deleting task", {icon: <IconX/>, type: "error"})
                } else {
                    toast("Deleted Task", {type: "success"}) //TODO: undo button?
                    navigate("#")
                }
            }
        ).catch(
            (err) => {
                toast("Error deleting task", {icon: <IconX/>, type: "error"})
                console.log(err)
            }
        )
    }


    function openDetails() {
        if (editable) return
        setTaskDetails(true)
    }


    return (
        <SlotProvider task={task}>
            <div className={styles.task} onClick={clickHandler.onClick}
                 onDoubleClick={clickHandler.onDoubleClick}>
                <UpperSlot/>
                <div className={styles.taskname}>
                    <div className={styles.name}>
                        <IconCircleCheck/>
                        <Text a="left" onBlur={handleBlur} ref={ref}> {/*TODO: dont use contentEditable*/}
                            {t?.name as string}
                        </Text>
                    </div>
                </div>
                <LowerSlot/>
                <div className={styles.hover}>
                    <div>
                        <button onClick={handleDelete}>
                            <IconTrash/>
                        </button>
                    </div>
                </div>
            </div>
            <TaskDetails open={taskDetails} setOpen={setTaskDetails} uuid={task}/>
        </SlotProvider>
    )
}